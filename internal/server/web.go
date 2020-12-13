package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/yvv4git/webhelpers"
)

const (
	timeoutRead              = 10
	timeoutWrite             = 10
	timeoutIdle              = 30
	timeoutGracefullShutdown = 10
)

// web - entity of web server.
type web struct {
	server *http.Server
}

// newWebServer - method for initilize server options.
func newWebServer(host string, port int, dirPath string) *web {
	return &web{
		server: &http.Server{
			Addr:         webhelpers.GetListenServerString(host, port),
			Handler:      getRouterWithShareDir(dirPath),
			ReadTimeout:  timeoutRead * time.Second,
			WriteTimeout: timeoutWrite * time.Second,
			IdleTimeout:  timeoutIdle * time.Second,
		},
	}
}

// Start - method for start web server.
func (s *web) Start() {
	log.Println("Run server")

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}

// Stop - method for stoping web server.
func (s *web) Stop() {
	log.Println("Stop web server")

	ctx, cancel := context.WithTimeout(context.Background(), timeoutGracefullShutdown*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
}

// getRouterWithShareDir function for create router
func getRouterWithShareDir(dirPath string) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", http.FileServer(http.Dir(dirPath)))
	return router
}
