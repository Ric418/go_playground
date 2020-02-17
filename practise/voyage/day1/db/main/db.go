package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func insert(com string, db *sql.DB) {
	d := time.Now().Format("2006/01/02")
	//fmt.Print(com)
	//fmt.Print(d)
	ins, err := db.Prepare("INSERT INTO reports(yd, de) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(d, com)
}

func show(db *sql.DB){
	
}

//type ResponseWriter{
//
//}
//
//type Report struct {
//
//}
//
//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}

func main() {
	flag.Parse()
	com := flag.Arg(0)
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/tutorial")
	if err != nil {
		panic(err)
	}
	//fmt.Println(reflect.TypeOf(db))
	//fmt.Printf("db = %+v\n", db.Stats())
	insert(com, db)
	show(db)
}