package http

import (
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func NewHandler(grpcServer *grpc.Server) http.Handler {
	wrappedGRPCWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			// ここでCORSの許可を行う。必要に応じてoriginをチェックして許可する。
			return true // すべてのオリジンを許可
		}),
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// gRPC-Webのリクエストかどうかを判定して適切なハンドラーに渡す
		if wrappedGRPCWebServer.IsGrpcWebRequest(r) || wrappedGRPCWebServer.IsAcceptableGrpcCorsRequest(r) || wrappedGRPCWebServer.IsGrpcWebSocketRequest(r) {
			wrappedGRPCWebServer.ServeHTTP(w, r)
			return
		}

		// HTTP/2でgRPCリクエストかどうかを判定
		contentType := r.Header.Get("Content-Type")
		if contentType == "application/grpc" || contentType == "application/grpc+proto" || contentType == "application/grpc+json" {
			wrappedGRPCWebServer.ServeHTTP(w, r)
			return
		}

		grpcServer.ServeHTTP(w, r)
	})
}