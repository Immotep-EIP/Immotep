package filesystem

import (
	"mime/multipart"
	"strings"

	"github.com/minio/minio-go/v7"
	"immotep/backend/models"
)

const (
	LeaseFolder = "lease"
)

func UploadLeasePDF(leaseId string, file File) minio.UploadInfo {
	return uploadPDF("document", LeaseFolder+"/"+leaseId+"/"+file.GetName(), file)
}

func UploadLeaseDocument(leaseId string, file *multipart.FileHeader) minio.UploadInfo {
	return uploadFile("document", LeaseFolder+"/"+leaseId+"/"+file.Filename, file)
}

func DocumentExists(filePath string) bool {
	return fileExists("document", filePath)
}

func GetDocument(filePath string) *models.DocumentResponse {
	url := getFileURL("document", filePath)
	if url == "" {
		return nil
	}
	return &models.DocumentResponse{
		Name: filePath[strings.LastIndex(filePath, "/")+1:], // Extract the file name from the path
		Link: url,
	}
}

func GetDocuments(filePaths []string) []models.DocumentResponse {
	var docs []models.DocumentResponse
	for _, path := range filePaths {
		doc := GetDocument(path)
		if doc != nil {
			docs = append(docs, *doc)
		}
	}
	return docs
}

func DeleteLeaseDocument(leaseId string, fileName string) bool {
	return deleteFile("document", LeaseFolder+"/"+leaseId+"/"+fileName)
}
