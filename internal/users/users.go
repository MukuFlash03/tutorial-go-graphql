package users

import (
	"database/sql"
	"github.com/MukuFlash03/hackernews/internal/pkg/db/postgres"
	"golang.org/x/crypto/bcrypt"

	"log"
	"github.com/MukuFlash03/hackernews/pkg/utils"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) Create() {
	log.Print("Attempting to register new user...")

	stmt, err := database.Db.Prepare("INSERT INTO Users(Username, Password) VALUES($1, $2)")
	// print(stmt)
	utils.CheckError(err, "fatal")
	
	defer stmt.Close()

	hashedPassword, err := HashPassword(user.Password)
	_, err = stmt.Exec(user.Username, hashedPassword)
	utils.CheckError(err, "fatal")
}

//GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	log.Print("Attempting to fetch existing user...")

	stmt, err := database.Db.Prepare("SELECT ID FROM Users WHERE Username = $1")
	utils.CheckError(err, "fatal")
	
	row := stmt.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}

func (user *User) Authenticate() bool{
	log.Print("Attempting to authenticate user...")

	stmt, err := database.Db.Prepare("SELECT Password FROM Users WHERE Username = $1")
	utils.CheckError(err, "fatal")

	row := stmt.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err != sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}