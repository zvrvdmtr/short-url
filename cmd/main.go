package main

import (
	"fmt"
	"github/zvrvdmtr/short-url/pkg/api"
	"github/zvrvdmtr/short-url/pkg/models"
	"net/http"

	_ "github.com/jackc/pgx/v4"
)

func main() {
	models.InitDB("postgres://postgres:postgres@0.0.0.0:5432/postgres")
	defer models.CloseDB()
	http.HandleFunc("/create", api.CreateLink)
	http.HandleFunc("/", api.RedirectTolink)
	fmt.Println("Start server on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
