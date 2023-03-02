package app

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Context struct {
	Port    int
	BaseURL string
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func NewContext() *Context {

	prt, err := strconv.Atoi(os.Getenv("APP_WEB_SERVER_PORT"))
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
	return &Context{
		Port:    prt,
		BaseURL: os.Getenv("APP_WEB_SERVER_URL"),
	}
}

func (c *Context) WebServer() {}
