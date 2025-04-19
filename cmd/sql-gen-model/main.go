package main

import (
	"flag"
	"log"
	"os"

	"github.com/wisaitas/standard-golang/pkg"
)

// go run cmd/sql-gen-model/main.go -source deployment/docker-images/liquibase/changesets/standard-service/up/ -dest internal/models
func main() {
	sourcePath := flag.String("source", "", "Source directory containing SQL files")
	destPath := flag.String("dest", "", "Destination directory for generated model files")

	flag.Parse()

	if (*sourcePath != "" && *destPath == "") || (*sourcePath == "" && *destPath != "") {
		log.Fatal("Source and destination paths are required")
	}

	if *sourcePath != "" && *destPath != "" {
		converter := pkg.NewSQLToModelConverter(*sourcePath, *destPath)
		if err := converter.GenerateModels(); err != nil {
			log.Fatalf("error generating models: %v", err)
		}

		log.Println("models generated successfully")

		os.Exit(0)
	}

}
