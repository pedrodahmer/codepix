package main

import (
	"github.com/pedrodahmer/codepix/infrastructure/db"
	"github.com/pedrodahmer/codepix/application/grpc"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}