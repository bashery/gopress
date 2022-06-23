package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/auth", auth)

	http.ListenAndServe(":8080", nil)
}
func auth(w http.ResponseWriter, r *http.Request) {
	// check boot serial if not run on aother ip addres
	url := r.URL.Query()
	serial := url.Get("serial")
	fmt.Println(serial)

	ip := url.Get("ip")
	fmt.Println(ip)

	// get remote address
	addr := r.RemoteAddr

	fmt.Fprintf(w, addr)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	tn := time.Now().Unix()
	fmt.Fprintf(w, "time is :%d \n", tn)
}

// get time
func selectTime(db *sql.DB, serial string) string {
	time := ""
	row := db.QueryRow("select ts from boots where serial=?", serial)
	defer db.Close()

	if err := row.Scan(&time); err != nil {
		return err.Error()
	}
	return time
}

//fs := http.FileServer(http.Dir("static/"))
//http.Handle("/static/", http.StripPrefix("/static/", fs))

type Boot struct {
	Serial  string
	Address string
}

func initDatabase() *sql.DB {
	db, e := sql.Open("mysql", "root:123456@/lisence")
	ErrorCheck(e)

	// close database after all work is done
	defer db.Close()

	PingDB(db)
	return db
}

// INSERT INTO DB
func insert(db *sql.DB) {

	// prepare
	stmt, e := db.Prepare("insert into boots(address, serial) values (?, ?)")
	ErrorCheck(e)

	//execute
	_, e = stmt.Exec("123.123.123") //,serial())
	ErrorCheck(e)

}

//Update db
func Update(db *sql.DB) {

	stmt, e := db.Prepare("update boots set addres=? where bootid=?")
	ErrorCheck(e)

	// execute
	_, e = stmt.Exec("", "5")
	ErrorCheck(e)
}

// delete data
func delete(db *sql.DB) {
	stmt, e := db.Prepare("delete from boots where serial=?")
	ErrorCheck(e)

	// delete 5th boot
	_, e = stmt.Exec("5")
	ErrorCheck(e)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}
