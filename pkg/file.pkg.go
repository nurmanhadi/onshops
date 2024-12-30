package pkg

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func AddFileImage(file *multipart.FileHeader) (string, error) {
	fileType := [3]string{".jpg", ".jpeg", ".png"}
	fileSize := file.Size
	if fileSize > 10*1024*1024 {
		return "", fmt.Errorf("max size 10MB")
	}
	fileName := time.Now().Format("20060102150405") + "-" + file.Filename
	fileName = strings.ToLower(fileName)
	ext := filepath.Ext(fileName)
	if ext != fileType[0] && ext != fileType[1] && ext != fileType[2] {
		return "", fmt.Errorf("only jpg, jpeg, png")
	}

	path := "./core/presentation/resources/img/products/" + fileName

	if err := SaveFile(file, path); err != nil {
		return "", fmt.Errorf("cannot save file because %s", err)
	}
	return fileName, nil
}

func SaveFile(fileHeader *multipart.FileHeader, filePath string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("cannot open image")
	}
	defer src.Close()
	dst, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("cannot create image")
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("cannot copy image")
	}
	return nil
}
func DeleteFile(fileName string) error {
	os.Remove("./core/presentation/resources/img/products/" + fileName)
	return nil
}
