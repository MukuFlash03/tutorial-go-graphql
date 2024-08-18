package links

import (
	"fmt"
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

func (link Link) AddLink() int64 {
	fmt.Println("Attempting to save new link...")
	fmt.Printf("Link Title = %s\n", link.Title)
	fmt.Printf("Link Address = %s\n", link.Address)
    db := database.Db
    _, err := db.Exec("INSERT INTO links (title, address) VALUES ($1, $2)", link.Title, link.Address)
	fmt.Println("Insert complete / failed...")
	utils.CheckError(err, "fatal")

	fmt.Println("Insert complete")

    return 123
}


func (link Link) Save() int64 {
	log.Print("Attempting to save new link...")

	stmt, err := database.Db.Prepare("INSERT INTO Links(Title, Address) VALUES($1, $2) RETURNING id")
	utils.CheckError(err, "fatal")
	
	defer stmt.Close()

	var insertedID int64
	err = stmt.QueryRow(link.Title, link.Address).Scan(&insertedID)
	utils.CheckError(err, "fatal")
	
	log.Print("Row inserted successfully!")
	return insertedID
}

func GetAll() []Link {
	log.Print("Attempting to fetch links...")

	stmt, err := database.Db.Prepare("select id, title, address from Links")
	utils.CheckError(err, "fatal")
	defer stmt.Close()

	rows, err := stmt.Query()
	utils.CheckError(err, "fatal")
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		utils.CheckError(err, "fatal")
		links = append(links, link)
	}

	err = rows.Err();
	utils.CheckError(err, "fatal")

	log.Print("Rows fetched successfully!")
	return links
}