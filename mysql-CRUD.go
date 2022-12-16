package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)
func main() {
    fmt.Println("Go MySQL Tutorial")

   
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

    if err != nil {
        log.Printf("Error %s when opening DB\n", err)
        panic(err.Error())
    }


    // fmt.Println("connected")
    name:="Abdul Shakoor"
    age:=21
    // insertUser(db,name,age)
    updateUser(db,name,age)
    // selectUsers(db)
    // deleteUser(db,age)
    defer db.Close()

    


}
// func insertUser(db *sql.DB, name string, age int) {
// 	result, err := db.Exec(`INSERT INTO test1 (name, age) VALUES (?, ?)`, name, age)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	id, err := result.LastInsertId()
// 	fmt.Println(id)
// }

// func selectUsers(db *sql.DB){
//     type user struct {
// 		name string
//         age int
// 	}
//     rows,err:=db.Query(`SELECT name,age FROM test1`)
//     if err!=nil{
//         log.Fatal(err)
//     }
//     defer rows.Close()
//     var users []user

//     for rows.Next() {
// 		var u user

// 		err := rows.Scan( &u.name, &u.age)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		users = append(users, u)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%#v", users)
// }


// func deleteUser(db *sql.DB,age int){
//     _,err:=db.Query(`DELETE FROM test1 WHERE age=?`,age)
//     if err!=nil{
//         log.Fatal(err)
//     }else{
//         fmt.Println("deleted successfully")
//     }
// }

func updateUser(db *sql.DB,name string,age int){
    _,err:=db.Query(`UPDATE test1  SET name=? WHERE age=?`,name,age)
    if err!=nil{
        log.Fatal(err,"error update")
    }
}