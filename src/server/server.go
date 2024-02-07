package server

import (
	"MydroX/shadow-technical-test/src/server/router"
	"fmt"
	"net/http"
	"os"
	"time"
)

func Run() {
	r := router.New()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	s.ListenAndServe()
}
