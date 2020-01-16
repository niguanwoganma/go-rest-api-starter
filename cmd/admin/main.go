package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ifaceless/go-starter/pkg/admin"
	"github.com/ifaceless/portal"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	portal.SetDebug(false)
	defer portal.CleanUp()

	err := godotenv.Load()
	if err != nil {
		logrus.Warn(err)
	}
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := admin.NewRouter()

	srvPort := os.Getenv("SERVER_PORT")
	if srvPort == "" {
		srvPort = "8000"
	}

	fmt.Printf("Admin Server is listening at: ':%s'\n", srvPort)
	err = r.Run(fmt.Sprintf(":%s", srvPort))
	if err != nil {
		panic(err)
	}
}
