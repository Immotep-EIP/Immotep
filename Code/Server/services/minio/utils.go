package minio

import (
	"io"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"immotep/backend/services"
)

type File interface {
	io.Reader
	GetName() string
	GetSize() int64
}

func uploadPDF(bucket string, filePath string, file File) minio.UploadInfo {
	fc := services.FileClient

	info, err := fc.Client.PutObject(fc.Context, bucket, filePath, file, file.GetSize(), minio.PutObjectOptions{ContentType: "application/pdf"})
	if err != nil {
		panic(err)
	}
	return info
}

func uploadFile(bucket string, filePath string, file *multipart.FileHeader) minio.UploadInfo {
	fc := services.FileClient

	fileContent, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer fileContent.Close()

	info, err := services.FileClient.Client.PutObject(fc.Context, bucket, filePath, fileContent, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		panic(err)
	}
	return info
}

func fileExists(bucket string, filePath string) bool {
	fc := services.FileClient
	_, err := fc.Client.StatObject(fc.Context, bucket, filePath, minio.StatObjectOptions{})
	return err == nil
}

func getFileObj(bucket string, filePath string) *minio.Object {
	fc := services.FileClient

	if !fileExists(bucket, filePath) {
		return nil
	}
	object, err := fc.Client.GetObject(fc.Context, bucket, filePath, minio.GetObjectOptions{})
	if err != nil {
		panic(err)
	}
	return object
}

func getFileURL(bucket string, filePath string) string {
	fc := services.FileClient

	if !fileExists(bucket, filePath) {
		return ""
	}
	u, err := fc.Client.PresignedGetObject(fc.Context, bucket, filePath, time.Hour*24, url.Values{})
	if err != nil {
		panic(err)
	}
	return u.String()
}

func deleteFile(bucket string, filePath string) bool {
	fc := services.FileClient

	if !fileExists(bucket, filePath) {
		return false
	}
	err := fc.Client.RemoveObject(fc.Context, bucket, filePath, minio.RemoveObjectOptions{})
	if err != nil {
		panic(err)
	}
	return true
}
