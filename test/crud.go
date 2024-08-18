package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
    db, err := sql.Open("postgres", "user=postgres password=mukuflash dbname=hackernews host=localhost port=5432 sslmode=disable")
    if err != nil {
        return nil, err
    }
    return db, nil
}

func createUser(username, email string) error {
    db, err := connectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", username, email)
    if err != nil {
        return err
    }

    return nil
}

func readUser(userID int) (string, string, error) {
    db, err := connectDB()
    if err != nil {
        return "", "", err
    }
    defer db.Close()

    var username, email string
    err = db.QueryRow("SELECT username, email FROM users WHERE id=$1", userID).Scan(&username, &email)
    if err != nil {
        return "", "", err
    }

    return username, email, nil
}

func updateUser(userID int, newEmail string) error {
    db, err := connectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("UPDATE users SET email=$1 WHERE id=$2", newEmail, userID)
    if err != nil {
        return err
    }

    return nil
}

func deleteUser(userID int) error {
    db, err := connectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM users WHERE id=$1", userID)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    // Create a new user
    err := createUser("JohnDoe", "johndoe@example.com")
    if err != nil {
        fmt.Println("Error creating user:", err)
    } else {
        fmt.Println("User created successfully")
    }

    // Read user data
    username, email, err := readUser(1)
    if err != nil {
        fmt.Println("Error reading user data:", err)
    } else {
        fmt.Printf("Username: %s, Email: %s\n", username, email)
    }

    // Update user data
    err = updateUser(1, "newemail@example.com")
    if err != nil {
        fmt.Println("Error updating user data:", err)
    } else {
        fmt.Println("User data updated successfully")
    }

    // Delete a user
    err = deleteUser(1)
    if err != nil {
        fmt.Println("Error deleting user:", err)
    } else {
        fmt.Println("User deleted successfully")
    }
}
