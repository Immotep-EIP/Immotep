package database

import (
	"immotep/backend/prisma/db"
	"immotep/backend/services"
)

func GetImageByID(id string) *db.ImageModel {
	pdb := services.DBclient
	image, err := pdb.Client.Image.FindUnique(db.Image.ID.Equals(id)).Exec(pdb.Context)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil
		}
		panic(err)
	}
	return image
}

func MockGetImageByID(c *services.PrismaDB) db.ImageMockExpectParam {
	return c.Client.Image.FindUnique(
		db.Image.ID.Equals("1"),
	)
}

func CreateImage(image db.ImageModel) db.ImageModel {
	pdb := services.DBclient
	newImage, err := pdb.Client.Image.CreateOne(
		db.Image.Data.Set(image.Data),
	).Exec(pdb.Context)
	if err != nil || newImage == nil {
		panic(err)
	}
	return *newImage
}

func MockCreateImage(c *services.PrismaDB, image db.ImageModel) db.ImageMockExpectParam {
	return c.Client.Image.CreateOne(
		db.Image.Data.Set(image.Data),
	)
}
