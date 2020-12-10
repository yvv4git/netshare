package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	timeoutRead              = 10
	timeoutWrite             = 10
	timeoutIdle              = 30
	timeoutGracefullShutdown = 10
)

// WebSrv - entity of web server.
type WebSrv struct {
	server *http.Server
}

// newWebServer - method for initilize server options.
func newWebServer(host string, port int, dirPath string) *WebSrv {
	return &WebSrv{
		server: &http.Server{
			Addr:         getListenServerString(host, port),
			Handler:      getRouterWithShareDir(dirPath),
			ReadTimeout:  timeoutRead * time.Second,
			WriteTimeout: timeoutWrite * time.Second,
			IdleTimeout:  timeoutIdle * time.Second,
		},
	}
}

// Start - method for start web server.
func (s *WebSrv) Start() {
	log.Println("Run server")

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}

// Stop - method for stoping web server.
func (s *WebSrv) Stop() {
	log.Println("Stop web server")

	ctx, cancel := context.WithTimeout(context.Background(), timeoutGracefullShutdown*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
}

// getListenServerString function for generate server host:port string.
func getListenServerString(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// getRouterWithShareDir function for create router
func getRouterWithShareDir(dirPath string) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir(dirPath)))
	return router
}
