package links

import (
	"context"
	database "github.com/jafossum/go-gql-hackernews/internal/pkg/db/mysql"
	"github.com/jafossum/go-gql-hackernews/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	UserID  *string
}

func (link *Link) User(ctx context.Context) *users.User {
	user, err := users.GetUserById(ctx, *link.UserID)
	if err != nil {
		log.Fatal(err)
	}
	return &user
}

func (link *Link) Save(ctx context.Context) int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address,UserId) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.ExecContext(ctx, link.Title, link.Address, link.UserID)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Row inserted")
	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("SELECT L.id, L.title, L.address, L.userID FROM Links L INNER JOIN Users U on L.ID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &link.UserID)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
