package minio

import (
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

const (
	UserProfileFolder    = "user_profile"
	PropertyFolder       = "property"
	RoomStateFolder      = "room_state"
	FurnitureStateFolder = "furniture_state"
	DamageFolder         = "damage"
)

func UploadUserProfileImage(userId string, file *multipart.FileHeader) minio.UploadInfo {
	return uploadFile("image", UserProfileFolder+"/"+userId+"/"+file.Filename, file)
}

func UploadPropertyImage(propertyId string, file *multipart.FileHeader) minio.UploadInfo {
	return uploadFile("image", PropertyFolder+"/"+propertyId+"/"+file.Filename, file)
}

func UploadRoomStateImage(roomStateId string, file *multipart.FileHeader) minio.UploadInfo {
	return uploadFile("image", RoomStateFolder+"/"+roomStateId+"/"+file.Filename, file)
}

func UploadFurnitureStateImage(furnitureStateId string, file *multipart.FileHeader) minio.UploadInfo {
	return uploadFile("image", FurnitureStateFolder+"/"+furnitureStateId+"/"+file.Filename, file)
}

func UploadDamageImage(damageId string, file *multipart.FileHeader) minio.UploadInfo {
	return uploadFile("image", DamageFolder+"/"+damageId+"/"+file.Filename, file)
}

func ImageExists(filePath string) bool {
	return fileExists("image", filePath)
}

func GetImageURL(filePath string) string {
	return getFileURL("image", filePath)
}

func GetImageURLs(filePaths []string) []string {
	var images []string
	for _, path := range filePaths {
		img := GetImageURL(path)
		if img != "" {
			images = append(images, img)
		}
	}
	return images
}

func GetImageObj(filePath string) *minio.Object {
	return getFileObj("image", filePath)
}

func GetImageObjs(filePaths []string) []minio.Object {
	var images []minio.Object
	for _, path := range filePaths {
		img := GetImageObj(path)
		if img != nil {
			images = append(images, *img)
		}
	}
	return images
}
