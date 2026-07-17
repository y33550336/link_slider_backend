package handler

import (
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/y33550336/link_slider/repository"
	"google.golang.org/grpc"
)

func newHTTPHandler(grpcServer *grpc.Server) http.Handler {
	wrappedGRPCWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if wrappedGRPCWebServer.IsGrpcWebRequest(r) || wrappedGRPCWebServer.IsAcceptableGrpcCorsRequest(r) || wrappedGRPCWebServer.IsGrpcWebSocketRequest(r) {
			wrappedGRPCWebServer.ServeHTTP(w, r)
			return
		}

		contentType := r.Header.Get("Content-Type")
		if contentType == "application/grpc" || contentType == "application/grpc+proto" || contentType == "application/grpc+json" {
			wrappedGRPCWebServer.ServeHTTP(w, r)
			return
		}

		grpcServer.ServeHTTP(w, r)
	})
}

// NewServer creates an HTTP server that serves gRPC and gRPC-Web requests.
func NewServer(addr string, repo *repository.Repository) *http.Server {
	grpcServer := newGRPCServer(repo)

	var protocols http.Protocols
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)

	return &http.Server{
		Addr:      addr,
		Handler:   newHTTPHandler(grpcServer),
		Protocols: &protocols,
	}
}
