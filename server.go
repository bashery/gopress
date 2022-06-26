package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db = initDatabase()
	defer db.Close()

	loopping(db)

	http.HandleFunc("/auth", auth)
	http.HandleFunc("/time", expiration)
	http.HandleFunc("/new", newBoot)
	http.HandleFunc("/update", changeIpAddr)
	http.HandleFunc("/delete", deleteBoot)

	http.ListenAndServe(":8080", nil)
}

// deleteBoot
func deleteBoot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")

	// delete boot from db by serial
	delete(db, serial)

	fmt.Fprintf(w, "serial : %s\n", serial)
}

// delete boot from db by serial
func delete(db *sql.DB, serial string) {
	stmt, e := db.Prepare("delete from licenses.boots where serial=?")
	ErrorCheck(e)

	_, e = stmt.Exec(serial)
	ErrorCheck(e)
}

// changeIpAddr update ip addr
func changeIpAddr(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")
	ipaddress := url.Get("ip")

	stmt, e := db.Prepare("update licenses.boots set ipaddress=? where serial=?")
	ErrorCheck(e)

	// execute
	_, e = stmt.Exec(ipaddress, serial)
	ErrorCheck(e)

	fmt.Fprintf(w, "serial : %s\nipaddress %s", serial, ipaddress)
}

func newBoot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")
	ipaddress := url.Get("ip")

	// prepare
	stmt, e := db.Prepare("insert into licenses.boots(serial, ipaddress ) values (?, ?)")
	ErrorCheck(e)

	//execute
	_, e = stmt.Exec(serial, ipaddress) //,serial())
	ErrorCheck(e)

	fmt.Fprintf(w, "serial : %s\nipaddress %s", serial, ipaddress)
}

func expiration(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")

	time := ""
	row := db.QueryRow("select ts from licenses.boots where serial=?", serial)

	if err := row.Scan(&time); err != nil {
		ErrorCheck(err)
	}

	fmt.Fprintf(w, "%s", time)
}

// TODO check this auth func
// check by serial
func auth(w http.ResponseWriter, r *http.Request) {
	// check boot serial if not run on aother ip addres
	url := r.URL.Query()
	serial := url.Get("serial")
	fmt.Println(serial)

	ip := url.Get("ip")
	fmt.Println(ip)

	addr := r.RemoteAddr

	fmt.Fprintf(w, addr)
}

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

// loop ping for active db connextion
func loopping(db *sql.DB) {
	go func() {
		for {
			err := db.Ping()
			ErrorCheck(err)
			time.Sleep(time.Minute * 1)
		}
	}()
}

//fs := http.FileServer(http.Dir("static/"))
//http.Handle("/static/", http.StripPrefix("/static/", fs))
