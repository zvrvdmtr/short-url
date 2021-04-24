package api

import (
	"encoding/json"
	"github/zvrvdmtr/short-url/pkg/models"
	"io/ioutil"
	"net/http"
	"strings"
)

type Url struct {
	Url string
}

var url Url

func CreateLink(w http.ResponseWriter, r *http.Request) {
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

		link, err := models.CreateNewShortUrl(url.Url)
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

func RedirectTolink(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		short := strings.TrimLeft(r.URL.Path, "/")
		link, err := models.FindLinkByShortUrl(short)
		if err != nil {
			http.Error(w, "Bad link!", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusPermanentRedirect)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}