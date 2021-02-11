package main

import (
	"os"

	"github.com/cheldontk/codepix/application/grpc"
	"github.com/cheldontk/codepix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
