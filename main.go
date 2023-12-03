package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// This is where all the APIs will be set up to run
	r := gin.Default()
	r.Run() // Runs on any port specified in environment variable, default is 8080
}
