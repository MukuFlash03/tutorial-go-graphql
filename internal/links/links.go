package links

import (
	database "github.com/MukuFlash03/hackernews/internal/pkg/db/postgres"
	"github.com/MukuFlash03/hackernews/internal/users"
	"log"

	"github.com/MukuFlash03/hackernews/pkg/utils"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	log.Print("Attempting to save new link...")

	stmt, err := database.Db.Prepare("INSERT INTO Links(Title, Address, UserID) VALUES($1, $2, $3) RETURNING id")
	utils.CheckError(err, "fatal")
	
	defer stmt.Close()

	var insertedID int64
	err = stmt.QueryRow(link.Title, link.Address, link.User.ID).Scan(&insertedID)
	utils.CheckError(err, "fatal")
	
	log.Print("Row inserted successfully!")
	return insertedID
}

func GetAll() []Link {
	log.Print("Attempting to fetch links...")

	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID") // changed
	utils.CheckError(err, "fatal")
	defer stmt.Close()

	rows, err := stmt.Query()
	utils.CheckError(err, "fatal")
	defer rows.Close()

	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		utils.CheckError(err, "fatal")
		link.User = &users.User{
			ID: id,
			Username: username,
		}
		links = append(links, link)
	}

	err = rows.Err();
	utils.CheckError(err, "fatal")

	log.Print("Rows fetched successfully!")
	return links
}