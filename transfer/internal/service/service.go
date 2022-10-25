package service

import (
	"file-share/transfer/pkg/logs"
	"fmt"
	"os"
)

const path = "C:\\Users\\ssss\\Desktop\\static"

func SaveImg(fileName string, data []byte) error {
	logger := logs.GetLogger()

	crFile, err := os.Create(fileName)
	if err != nil {
		logger.Fatalf("%v ", err)
		return err
	}

	defer crFile.Close()

	_, err = crFile.Write(data)

	return err
}

// CreateDirectoryByUserName Создает директории по имени пользвателя: {username}/content/
func CreateDirectoryByUserName(user string) (string, error) {
	staticPath := fmt.Sprintf("%s/%s/content", path, user)
	return staticPath, os.MkdirAll(staticPath, 0777)
}
