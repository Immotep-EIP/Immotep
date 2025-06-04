package models

import (
	"encoding/base64"
	"strings"

	"immotep/backend/prisma/db"
)

type ImageRequest struct {
	Data string `binding:"required,datauri" json:"data"`
}

func (i *ImageRequest) ToDbImage() *db.ImageModel {
	var imgType db.ImageType

	switch {
	case strings.HasPrefix(i.Data, "data:image/png;base64,"):
		i.Data = strings.TrimPrefix(i.Data, "data:image/png;base64,")
		imgType = db.ImageTypePng
	case strings.HasPrefix(i.Data, "data:image/jpeg;base64,"):
		i.Data = strings.TrimPrefix(i.Data, "data:image/jpeg;base64,")
		imgType = db.ImageTypeJpeg
	case strings.HasPrefix(i.Data, "data:image/jpg;base64,"):
		i.Data = strings.TrimPrefix(i.Data, "data:image/jpg;base64,")
		imgType = db.ImageTypeJpeg
	default:
		return nil
	}

	decoded, err := base64.StdEncoding.DecodeString(i.Data)
	if err != nil {
		return nil
	}

	return &db.ImageModel{
		InnerImage: db.InnerImage{
			Data: decoded,
			Type: imgType,
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
	switch model.Type {
	case db.ImageTypePng:
		i.Data = "data:image/png;base64,"
	case db.ImageTypeJpeg:
		i.Data = "data:image/jpeg;base64,"
	default:
		panic("unknown image type")
	}
	i.Data += base64.StdEncoding.EncodeToString(model.Data)
	i.CreatedAt = model.CreatedAt
}

func DbImageToResponse(pc db.ImageModel) ImageResponse {
	var resp ImageResponse
	resp.FromDbImage(pc)
	return resp
}
