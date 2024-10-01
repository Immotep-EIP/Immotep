package database

import (
	"context"
	"immotep/backend/prisma/db"
	"log"
	"testing"
)

type PrismaDB struct {
	Client  *db.PrismaClient
	Context context.Context
}

var DBclient = &PrismaDB{}

func ConnectDB() *PrismaDB {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	DBclient.Client = client
	DBclient.Context = context.Background()
	return DBclient
}

func ConnectDBTest() (*PrismaDB, *db.Mock, func(t *testing.T)) {
	client, mock, ensure := db.NewMock()

	DBclient.Client = client
	DBclient.Context = context.Background()

	return DBclient, mock, ensure
}
