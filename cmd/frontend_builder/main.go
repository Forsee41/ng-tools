package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	components "github.com/Forsee41/ng-tools/components"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cant load doentv: ", err)
		return
	}
	frontendDir, ok := os.LookupEnv("FRONTEND_SRC_DIR")
	if !ok {
		fmt.Println("Can't load FRONTEND_SRC_DIR env var")
		return
	}
	filename := "index.html"
	fullFilename := filepath.Join(frontendDir, filename)

	f, err := os.OpenFile(fullFilename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cant open file, ", err)
		return
	}
	defer f.Close()

	ctx := context.Background()
	err = f.Truncate(0)
	if err != nil {
		fmt.Println("Unable to truncate file, ", err)
		return
	}

	err = components.Index().Render(ctx, f)
	if err != nil {
		fmt.Println(err)
	}
}
