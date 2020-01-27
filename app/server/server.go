package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"time"
)

func StartServer(port int) {
	errLogger, err := newErrorLogger()
	if err != nil {
		log.Panicf("Error: failed to create error logger %v", err)
	}
	r := SetupRouter()
	fmt.Printf("Listening on port %v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
		// log error
		errLogger.Printf("%v", err)
		// print to stderr
		log.Printf("%v", err)
	}
}

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// timeout in one minute
	r.Use(middleware.Timeout(60 * time.Second))
	// TODO: initialize routes here
	return r
}

func newErrorLogger() (*log.Logger, error) {
	y, m, d := time.Now().Date()
	// create file for error logging
	errorLog, err := os.OpenFile(
		fmt.Sprintf("%v-%v-%v_err.log", y, m, d),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return nil, err
	}
	errLogger := log.New(errorLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return errLogger, nil
}
