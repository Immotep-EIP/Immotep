package services

import (
	"context"
	"testing"

	"keyz/backend/prisma/db"
)

type PrismaDB struct {
	Client  *db.PrismaClient
	Context context.Context
}

var DBclient = &PrismaDB{}

func ConnectDB() (*PrismaDB, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	DBclient.Client = client
	DBclient.Context = context.Background()
	return DBclient, nil
}

func ConnectDBTest() (*PrismaDB, *db.Mock, func(t *testing.T)) {
	client, mock, ensure := db.NewMock()

	DBclient.Client = client
	DBclient.Context = context.Background()

	return DBclient, mock, ensure
}
