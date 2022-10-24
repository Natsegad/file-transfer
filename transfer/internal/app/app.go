package app

import (
	shttp "file-share/transfer/internal/delivery/http"
	"file-share/transfer/pkg/logs"

	"github.com/gin-gonic/gin"
)

func Start() {
	l := logs.GetLogger()

	l.Info("Initialize Router !")
	router := gin.Default()

	l.Info("Start initialize routers !")
	router.Handle("POST", "/file-add", shttp.FileAddPage)
	router.Handle("GET", "/file-add", shttp.GetAll)

	l.Info("Starting server port: 8080 !")
	router.Run(":8080")
}
