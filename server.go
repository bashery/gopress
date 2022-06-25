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

// deleteBoot update ip addr
func deleteBoot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")

	delete(db, serial)

	fmt.Fprintf(w, "serial : %s\n", serial)
}

// delete data
func delete(db *sql.DB, serial string) {
	stmt, e := db.Prepare("delete from licenses.boots where serial=?")
	ErrorCheck(e)

	// delete 5th boot
	_, e = stmt.Exec(serial)
	ErrorCheck(e)
}

// changeIpAddr update ip addr
func changeIpAddr(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")
	ipaddress := url.Get("ip")

	updateIpAddr(db, ipaddress, serial)

	fmt.Fprintf(w, "serial : %s\nipaddress %s", serial, ipaddress)
}

//updateIpAddr in db
func updateIpAddr(db *sql.DB, ipaddress, serial string) {

	stmt, e := db.Prepare("update licenses.boots set ipaddress=? where serial=?")
	ErrorCheck(e)

	// execute
	_, e = stmt.Exec(ipaddress, serial)
	ErrorCheck(e)
}

func newBoot(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")
	ipaddress := url.Get("ip")

	insert(db, serial, ipaddress)

	fmt.Fprintf(w, "serial : %s\nipaddress %s", serial, ipaddress)
}

// INSERT INTO DB
func insert(db *sql.DB, serial, ipaddress string) {
	// prepare
	stmt, e := db.Prepare("insert into licenses.boots(serial, ipaddress ) values (?, ?)")
	ErrorCheck(e)

	//execute
	_, e = stmt.Exec(serial, ipaddress) //,serial())
	ErrorCheck(e)
}

func expiration(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()
	serial := url.Get("serial")
	time := selectTime(db, serial)
	fmt.Println("time exp is : ", time)

	fmt.Fprintf(w, "%s", time)
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
