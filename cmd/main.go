package main

import (
	"fmt"
	"github/zvrvdmtr/short-url/pkg/api"
	"github/zvrvdmtr/short-url/pkg/models"
	"net/http"

	_ "github.com/jackc/pgx/v4"
)

func main() {
	conn, err := models.InitDB("postgres://postgres:postgres@0.0.0.0:5432/postgres")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer models.CloseDB()
	http.HandleFunc("/create", api.CreateLink(conn))
	http.HandleFunc("/", api.RedirectTolink(conn))
	fmt.Println("Start server on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
