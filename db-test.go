package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	//"os"
	//"strings"
	"flag"
)

var debug = flag.Bool("debug", true, "enable debugging")
var password = flag.String("password", "", "the database password")
var port *int = flag.Int("port", 1433, "the database port")

//var server = flag.String("server", "WDC1SQL12", "the database server")
var server = flag.String("server", "DP-CA53-SQL.dp.local", "the database server")
var user = flag.String("user", "", "the database user")
var database = flag.String("database", "IISLogs", "the database")

func main() {
	flag.Parse() // parse the command line args

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user=%s", *server, *port, *database, *user)

	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	/*
		stmt, err := conn.Prepare("select top 10 LogRow,LogFileName from IISLog")
		if err != nil {
			log.Fatal("Prepare failed:", err.Error())
		}
		defer stmt.Close()

		row := stmt.QueryRow()
		var somenumber int64
		var somechars string
		err = row.Scan(&somenumber, &somechars)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}

		fmt.Printf("somenumber:%d\n", somenumber)
		fmt.Printf("somechars:%s\n", somechars)*/

	rows, err := conn.Query("select top 10 LogRow,LogFileName from IISLog")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}

	for rows.Next() {
		var somenumber int64
		var somechars string
		err = rows.Scan(&somenumber, &somechars)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}
		fmt.Printf("somenumber:%d somechars:%s\n", somenumber, somechars)
	}

	fmt.Printf("bye\n")
}
