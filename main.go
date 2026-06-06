package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	testgen "github.com/y33550336/link_slider/pb"

	httpserver "github.com/y33550336/link_slider/http"
)

// gameServer は GameService の最小実装です。
type gameServer struct {
	testgen.UnimplementedGameServiceServer
}

func (s *gameServer) Ping(ctx context.Context, req *testgen.PingRequest) (*testgen.PingResponse, error) {
	log.Printf("Ping from client: %s", req.GetMessage())
	return &testgen.PingResponse{Message: "pong: " + req.GetMessage()}, nil
}

func (s *gameServer) MovePlayer(ctx context.Context, req *testgen.MovePlayerRequest) (*testgen.MovePlayerResponse, error) {
	log.Printf("MovePlayer player=%s x=%d y=%d", req.GetPlayerId(), req.GetX(), req.GetY())
	return &testgen.MovePlayerResponse{Ok: true, Message: "moved"}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Register our generated GameService implementation
	testgen.RegisterGameServiceServer(grpcServer, &gameServer{})

	httpHandler := httpserver.NewHandler(grpcServer)

	wrappedGRPCWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			// ここでCORSの許可を行う。必要に応じてoriginをチェックして許可する。
			return true // すべてのオリジンを許可
		}),
	)

	roothandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// gRPC-Webのリクエストかどうかを判定して適切なハンドラーに渡す
		if wrappedGRPCWebServer.IsGrpcWebRequest(r) || wrappedGRPCWebServer.IsAcceptableGrpcCorsRequest(r) || wrappedGRPCWebServer.IsGrpcWebSocketRequest(r) {
			wrappedGRPCWebServer.ServeHTTP(w, r)
			return
		}

		// HTTP/2でgRPCリクエストかどうかを判定
		contentType := r.Header.Get("Content-Type")
		if r.ProtoMajor == 2 && strings.HasPrefix(contentType, "application/grpc") {
			wrappedGRPCWebServer.ServeHTTP(w, r)
			return
		}

		// それ以外のリクエストは通常のHTTPリクエストとして処理
		httpHandler.ServeHTTP(w, r)
	})

	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(roothandler, h2s),
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
