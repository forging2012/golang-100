package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	http.HandleFunc("/", insert)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func insert(resp http.ResponseWriter, req *http.Request) {
	db, _ := sql.Open("mysql", "root:root@/test")
	defer db.Close()

	_, err := db.Exec(
		"insert into pressure(val) values(456)",
	)

	if err != nil {
		fmt.Println(err)
	}

}