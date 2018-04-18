package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/takama/router"
)

var log = logrus.New()

// Run server: go build && step-by-step
// Try requests: curl http://127.0.0.1:8000/test
func main() {
	port := os.Getenv("SERVICE_PORT")

	r := router.New()
	r.Logger = logger
	r.GET("/", home)
	r.Listen(":" + port)
}
