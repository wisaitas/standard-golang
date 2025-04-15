package scripts

import (
	"flag"
	"log"
	"os"
)

func HandleArgument() {
	sourcePath := flag.String("source", "", "Source directory containing SQL files")
	destPath := flag.String("dest", "", "Destination directory for generated model files")

	flag.Parse()

	if (*sourcePath != "" && *destPath == "") || (*sourcePath == "" && *destPath != "") {
		log.Fatal("Source and destination paths are required")
	}

	if *sourcePath != "" && *destPath != "" {
		converter := NewSQLToModelConverter(*sourcePath, *destPath)
		if err := converter.GenerateModels(); err != nil {
			log.Fatalf("error generating models: %v", err)
		}

		log.Println("models generated successfully")

		os.Exit(0)
	}
}
