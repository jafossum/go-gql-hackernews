package links

import (
	database "github.com/jafossum/go-gql-hackernews/internal/pkg/db/mysql"
	"github.com/jafossum/go-gql-hackernews/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address,UserId) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
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
	stmt, err := database.Db.Prepare("SELECT L.id, L.title, L.address, U.ID, U.username FROM Links L INNER JOIN Users U on L.ID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		}
		links = append(links, link)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
