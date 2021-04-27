package services

import (
	"context"
	"fmt"
	// "fmt"
	"testing"

	"github.com/jackc/pgx/v4"
)

type mockRow struct {
	id int
	url string
}

func (mr mockRow) Scan(dest ...interface{}) error {
	x, ok := dest[0].(int)
	fmt.Println(x)
	fmt.Println(ok)
	id := dest[0].(*int)
	url := dest[1].(*string)

	*id = mr.id
	*url = mr.url
	return nil
}

type mockConnection struct {}

func(c mockConnection) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	row := mockRow{1, "https://google.com"}
	return row
}

var expectableLink = Link{1, "https://google.com", "b"}

func TestCreateNewShortUrl(t *testing.T) {
	mock := mockConnection{}
	link, _ := CreateNewShortUrl(mock, "https://google.com")
	if (link.Id != expectableLink.Id || link.Url != expectableLink.Url || link.ShortUrl != expectableLink.ShortUrl) {
		t.Errorf("got %v want %v", link, expectableLink)
	}
}

func TestFindLinkByShortUrl(t *testing.T) {
	mock := mockConnection{}
	link, _ := FindLinkByShortUrl(mock, "b")
	if (link.Id != expectableLink.Id || link.Url != expectableLink.Url || link.ShortUrl != expectableLink.ShortUrl) {
		t.Errorf("got %v want %v", link, expectableLink)
	}
}