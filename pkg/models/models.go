package models

import (
	"context"
	"github/zvrvdmtr/short-url/pkg/generator"
)

type Link struct {
	Id int
	Url string 
	ShortUrl string
}

func CreateNewShortUrl(url string) (Link, error) {
	conn := GetDB()
	var link Link
	row := conn.QueryRow(context.Background(), "INSERT INTO link (url, short_url) values ($1, $2) RETURNING id, url", url, "shortUrl")
	err := row.Scan(&link.Id, &link.Url)
	shortPath := generator.ShortUrlGenerator(link.Id)
	row = conn.QueryRow(context.Background(), "UPDATE link set short_url = $1 where id = $2 RETURNING short_url", shortPath, link.Id)
	err = row.Scan(&link.ShortUrl)
	return link, err
}

func FindLinkByShortUrl(shortUrl string) (Link, error) {
	var link Link
	row := conn.QueryRow(context.Background(), "select * from link where short_url = $1", shortUrl)
	err := row.Scan(&link.Id, &link.Url, &link.ShortUrl)
	return link, err
}