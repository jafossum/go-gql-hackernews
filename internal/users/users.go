package users

import (
	"context"
	"database/sql"
	database "github.com/jafossum/go-gql-hackernews/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create(ctx context.Context) {
	stmt, err := database.Db.Prepare("INSERT INTO Users(Username,Password) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := HashPassword(user.Password)
	_, err = stmt.ExecContext(ctx, user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func (user *User) Authenticate(ctx context.Context) bool {
	stmt, err := database.Db.Prepare("SELECT Password FROM Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRowContext(ctx, user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

func GetUserIdByUsername(ctx context.Context, username string) (int, error) {
	stmt, err := database.Db.Prepare("SELECT ID FROM Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRowContext(ctx, username)

	var id int
	err = row.Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}
	return id, nil
}

func GetUserById(ctx context.Context, id string) (User, error) {
	stmt, err := database.Db.Prepare("SELECT ID, Username FROM Users WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRowContext(ctx, id)

	var user User
	err = row.Scan(&user.ID, &user.Username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return User{}, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
