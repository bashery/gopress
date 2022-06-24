package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db = initDatabase()
	defer db.Close()
	//PingDB(db)

	http.HandleFunc("/auth", auth)
	http.HandleFunc("/time", expiration)
	http.HandleFunc("/new", newBoot)

	http.ListenAndServe(":8080", nil)
}

func newBoot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")

	fmt.Fprintf(w, "%s", serial)
}

func expiration(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")
	time := selectTime(db, serial)
	fmt.Println("time exp is : ", time)

	fmt.Fprintf(w, "time is :%s\n", time)
}

// get time
func selectTime(db *sql.DB, serial string) string {
	time := ""
	row := db.QueryRow("select ts from licenses.boots where serial=?", serial)

	if err := row.Scan(&time); err != nil {
		fmt.Println("error here", err)
		return err.Error()
	}
	return time
}

// check by serial
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

// INSERT INTO DB
func insert(db *sql.DB) {
	// prepare
	stmt, e := db.Prepare("insert into licenses.boots(address, serial) values (?, ?)")
	ErrorCheck(e)

	//execute
	_, e = stmt.Exec("123.123.123") //,serial())
	ErrorCheck(e)

}

//Update db
func Update(db *sql.DB) {

	stmt, e := db.Prepare("update licenses.boots set addres=? where bootid=?")
	ErrorCheck(e)

	// execute
	_, e = stmt.Exec("", "5")
	ErrorCheck(e)
}

// delete data
func delete(db *sql.DB) {
	stmt, e := db.Prepare("delete from licenses.boots where serial=?")
	ErrorCheck(e)

	// delete 5th boot
	_, e = stmt.Exec("5")
	ErrorCheck(e)
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

//fs := http.FileServer(http.Dir("static/"))
//http.Handle("/static/", http.StripPrefix("/static/", fs))

// initialaze database
func initDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/licenses")
	ErrorCheck(err)

	return db
}

func ErrorCheck(err error) {
	if err != nil {
		println(err.Error())
	}
}
