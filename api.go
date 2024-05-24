package main

import (
	"log"
	"net/http"
	"time"
)

type apiServer struct {
	addr string
}

func newApiServer(addr string) *apiServer {
	return &apiServer{addr: addr}
}

type router struct {
	*http.ServeMux
}

func (s *apiServer) run() error {
	router := &router{http.NewServeMux()}
	router.setupRoutes()

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	middlewareChain := middlewareChain(requestLoggerMiddleware, requireAuthMiddleware)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(v1),
	}

	log.Printf("Starting server on %s", s.addr)

	return server.ListenAndServe()
}

type middleware func(http.Handler) http.HandlerFunc

func middlewareChain(middlewares ...middleware) middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next.ServeHTTP
	}
}

func requestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] Request: %s %s", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func requireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		providedToken := r.Header.Get("Authorization")

		if providedToken != getEnvVariable("API_TOKEN") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
