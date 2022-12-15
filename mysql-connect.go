package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root@localhost:mysql@tcp(127.0.0.1:3306)/test")

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Printf("Error %s when opening DB\n", err)
        panic(err.Error())
    }

    err=db.ping()
    if err != nil {
        log.Printf("Error %s when pimg db\n", err)
        panic(err.Error())
    }
    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    // perform a db.Query insert
    // insert, err := db.Query("INSERT INTO test1 VALUES ( "susu", 21)"

    // if there is an error inserting, handle it
    // if err != nil {
    //     panic(err.Error())
    // }
    // // be careful deferring Queries if you are using transactions
    // defer insert.Close()


}
