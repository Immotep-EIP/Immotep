package imageservice

import (
	"immotep/backend/database"
	"immotep/backend/prisma/db"
)

func GetByID(id string) *db.ImageModel {
	pdb := database.DBclient
	image, err := pdb.Client.Image.FindUnique(db.Image.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return image
}

func Create(image db.ImageModel) db.ImageModel {
	pdb := database.DBclient
	newImage, err := pdb.Client.Image.CreateOne(
		db.Image.Data.Set(image.Data),
	).Exec(pdb.Context)
	if err != nil || newImage == nil {
		panic(err)
	}
	return *newImage
}
