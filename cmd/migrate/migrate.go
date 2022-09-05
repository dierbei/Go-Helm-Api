package main

import (
	"log"

	"githup.com/dierbei/go-helm-api/internal/pkg/mysql"
	helmrepo "githup.com/dierbei/go-helm-api/internal/repository/helm"
)

func main() {
	if err := mysql.GetDb().AutoMigrate(
		helmrepo.NewRepository(),
	); err != nil {
		log.Fatal(err)
	}

	log.Println("AutoMigrate success...")
}
