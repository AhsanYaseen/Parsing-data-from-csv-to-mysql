package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	fname      string
	lname      string
	age        string
	bloodgroup string
}

func main() {

	dataparse()
}
func dataparse() {
	filepath := "src//csv//data.csv"

	openfile, err := os.Open(filepath) // passing file to Open function
	checkError("Error Reading fle", err)

	fileReader := csv.NewReader(openfile) //creating reader for file
	checkError("Error opening fle", err)

	for {
		value, err := fileReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		//for accessing values individually
		s := person{fname: value[0], lname: value[1], bloodgroup: value[2], age: value[3]}
		//Openining database , csvfile is the name of db
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/csvfile")

		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		fmt.Printf("sucessfully connected\n")
		//inserting data to database.
		insert, err := db.Query("INSERT into student (fname, lname, bgroup, age ) VALUES (" + s.fname + ", " + s.lname + "," + s.bloodgroup + ", " + s.age + ")")

		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()

	}
}
func checkError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
