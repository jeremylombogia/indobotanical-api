package transaction

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

func UploadFile(file *multipart.FileHeader, transactionId string) (string, error) {
	extension := filepath.Ext(file.Filename)
	randomNumber := strconv.Itoa(rand.Intn(1000))

	fileName := fmt.Sprintf("cdn/payment-proof/%s%s%s", transactionId, randomNumber, extension)

	src, err := file.Open()
	if err != nil {
		return fileName, err
	}
	defer src.Close()

	dst, err := os.Create(fileName)
	if err != nil {
		return fileName, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return fileName, err
	}

	return fileName, err
}

func CheckFileExtension(file *multipart.FileHeader) bool {
	extension := filepath.Ext(file.Filename)

	if extension == ".JPG" || extension == ".jpg" || extension == ".jpeg" || extension == ".png" {
		return true
	}

	return false
}
