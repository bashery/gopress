// tool to manage lisences
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var ip, ser string

	flag.StringVar(&ser, "ser", "", "The 'ser' must be serial! Default is \"\"")
	flag.StringVar(&ip, "ip", "", "The 'ip' must be ip addres Default is \"\"")

	// parse flags from command line
	flag.Parse()

	// output
	fmt.Println("flugs is ", ser, ip)

	//
	nserial := newSerial()
	resp, err := http.Get("http://localhost:8080/new?serial=" + nserial + "&ip=" + ip)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("send: ", nserial)
	fmt.Println("rese: ", string(body))

}

func newSerial() (serial string) {
	chars := []string{
		"q", "w", "e", "r", "t", "y", "u", "i", "o",
		"a", "s", "d", "f", "g", "h", "l", "k", "j",
		"A", "B", "C", "D", "E", "F", "J", "H", "E"}

	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 10; i++ {
		serial += chars[rand.Intn(len(chars)-1)]
	}
	return serial
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	panic(err)
}

// initialaze database
func initDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/licenses")
	if err != nil {
		panic(err)
	}
	return db
}
