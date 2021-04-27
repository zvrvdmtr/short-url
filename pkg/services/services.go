package services

import (
	"context"
	"github/zvrvdmtr/short-url/pkg/generator"
	"github/zvrvdmtr/short-url/pkg/models"
)

type Link struct {
	Id int
	Url string 
	ShortUrl string
}

func CreateNewShortUrl(conn models.DBConnect, url string) (Link, error) {
	var link Link
	row := conn.QueryRow(context.Background(), "insert into link (url) values ($1) returning id, url", url)
	err := row.Scan(&link.Id, &link.Url)
	shortPath := generator.ShortUrlGenerator(link.Id)
	link.ShortUrl = shortPath
	return link, err
}

func FindLinkByShortUrl(conn models.DBConnect, shortUrl string) (Link, error) {
	var link Link
	id := generator.BijectiveDecode(shortUrl)
	row := conn.QueryRow(context.Background(), "select * from link where id = $1", id)
	err := row.Scan(&link.Id, &link.Url)
	link.ShortUrl = shortUrl
	return link, err
}