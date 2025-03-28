package models

import (
	"encoding/base64"

	"immotep/backend/prisma/db"
)

type ImageRequest struct {
	Data string `binding:"required,base64" json:"data"`
}

func (i *ImageRequest) ToDbImage() *db.ImageModel {
	decoded, err := base64.StdEncoding.DecodeString(i.Data)
	if err != nil {
		return nil
	}

	return &db.ImageModel{
		InnerImage: db.InnerImage{
			Data: decoded,
		},
	}
}

func StringToDbImage(data string) *db.ImageModel {
	return (&ImageRequest{Data: data}).ToDbImage()
}

type ImageResponse struct {
	ID        string      `json:"id"`
	Data      string      `json:"data"`
	CreatedAt db.DateTime `json:"created_at"`
}

func (i *ImageResponse) FromDbImage(model db.ImageModel) {
	i.ID = model.ID
	i.Data = base64.StdEncoding.EncodeToString(model.Data)
	i.CreatedAt = model.CreatedAt
}

func DbImageToResponse(pc db.ImageModel) ImageResponse {
	var resp ImageResponse
	resp.FromDbImage(pc)
	return resp
}
