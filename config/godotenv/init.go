package config_godotenv

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	getConfig()
}

func Setup() {
	fmt.Println("Initial configuration")
}

var projectDirName = "base-gin-go"

func loadEnv() {
	projectName := regexp.MustCompile("^(.*" + projectDirName + ")")
	currentWork, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWork))
	err := godotenv.Load(string(rootPath) + "/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
