package http

import (
	"file-share/transfer/internal/service"
	"file-share/transfer/pkg/logs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileAddPage(c *gin.Context) {
	l := logs.GetLogger()
	l.Info("File add page \n")

	userName := c.Request.FormValue("user")

	// Проверка если userName пустой то смысла дальше исполнять код нет
	if userName == "" {
		l.Errorf("Error userName is empty !")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user name empty",
		})
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		l.Errorf("Error: %v \v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v ", err),
		})
		return
	}

	defer file.Close()

	data := make([]byte, fileHeader.Size)
	file.Read(data)

	dirPath := service.CreatePath(userName)

	if !service.IsHaveDirectory(userName) {
		dirPath, err = service.CreateDirectoryByUserName(userName)
		if err != nil {
			l.Errorf("%v ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v ", err),
			})
			return
		}
	} else {
		l.Infof("user %s have", userName)
	}

	err = service.SaveImg(dirPath+"/"+fileHeader.Filename, data)
	if err != nil {
		l.Errorf("%v ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v ", err),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func GetAll(c *gin.Context) {

}
