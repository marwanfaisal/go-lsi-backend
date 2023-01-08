package main

import (
	"go-lsi-backend/config"
	"go-lsi-backend/db"
)

func main() {
	db.Open()
	config.Router()
}
