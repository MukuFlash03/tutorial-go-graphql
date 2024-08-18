package links

import (
	"fmt"
	database "github.com/MukuFlash03/hackernews/internal/pkg/db/postgres"
	"github.com/MukuFlash03/hackernews/internal/users"
	"log"

	"github.com/MukuFlash03/hackernews/internal/utils"
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
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("Insert complete")

    return 123
}


func (link Link) Save() int64 {
	fmt.Println("Attempting to save new link...")
	fmt.Printf("Link Title = %s\n", link.Title)
	fmt.Printf("Link Address = %s\n", link.Address)

	stmt, err := database.Db.Prepare("INSERT INTO Links(Title, Address) VALUES($1, $2) RETURNING id")
	fmt.Println("Insert complete / failed...")
	utils.CheckError(err, "fatal")
	
	defer stmt.Close()

	fmt.Println("Insert complete")

	var insertedID int64
	err = stmt.QueryRow(link.Title, link.Address).Scan(&insertedID)
	
	utils.CheckError(err, "fatal")
	
	log.Print("Row inserted!")
	return insertedID
}