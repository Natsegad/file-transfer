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

func CreatePath(user string) string {
	return fmt.Sprintf("%s/%s/content", path, user)
}

// CreateDirectoryByUserName Создает директории по имени пользвателя: {username}/content/
func CreateDirectoryByUserName(user string) (string, error) {
	staticPath := CreatePath(user)
	return staticPath, os.MkdirAll(staticPath, 0777)
}

// IsHaveDirectory возвращает true если директория уже создана
func IsHaveDirectory(dir string) bool {
	logger := logs.GetLogger()

	files, err := os.ReadDir(path)
	if err != nil {
		logger.Errorf("%v ", err)
		return false
	}

	for _, v := range files {
		if v.IsDir() {
			if v.Name() == dir {
				return true
			}
		}
	}

	return false
}
