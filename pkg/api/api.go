package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github/zvrvdmtr/short-url/pkg/services"
	"github/zvrvdmtr/short-url/pkg/models"
	
)

type Url struct {
	Url string
}

var url Url

func CreateLink(conn models.DBConnect) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, _ := ioutil.ReadAll(r.Body)
			err := json.Unmarshal(body, &url)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
	
			if url.Url == "" {
				http.Error(w, "Url field must be not empty", http.StatusBadRequest)
				return
			}
	
			link, err := services.CreateNewShortUrl(conn, url.Url)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	
			jsData, err := json.Marshal(link)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsData)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func RedirectTolink(conn models.DBConnect) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			short := strings.TrimLeft(r.URL.Path, "/")
			link, err := services.FindLinkByShortUrl(conn, short)
			if err != nil {
				http.Error(w, "Bad link!", http.StatusNotFound)
				return
			}
	
			http.Redirect(w, r, link.Url, http.StatusPermanentRedirect)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}