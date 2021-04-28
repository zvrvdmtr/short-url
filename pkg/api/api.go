package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github/zvrvdmtr/short-url/pkg/services"
	"github/zvrvdmtr/short-url/pkg/models"
	
)

type Link struct {
	Url string
}

var link Link

func CreateLink(conn models.DBConnect) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, _ := ioutil.ReadAll(r.Body)
			unmarshalErr := json.Unmarshal(body, &link)
			if unmarshalErr != nil {
				http.Error(w, unmarshalErr.Error(), http.StatusBadRequest)
				return
			}
	
			if link.Url == "" {
				http.Error(w, "Url field must be not empty", http.StatusBadRequest)
				return
			}

			_, parseErr := url.ParseRequestURI(link.Url)
			if parseErr != nil {
				http.Error(w, "Invalid url", http.StatusBadRequest)
				return
			}
	
			link, createErr := services.CreateNewShortUrl(conn, link.Url)
			if createErr != nil {
				http.Error(w, createErr.Error(), http.StatusInternalServerError)
				return
			}
	
			jsData, marshalErr := json.Marshal(link)
			if marshalErr != nil {
				http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
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