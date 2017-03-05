package main

import (
	"github.com/mysqlx/mysqlx"
	"log"
)

//import "github.com/mysqlx/mysqlx"

func main() {
	mySession, err := mysqlx.GetNodeSession(mysqlx.SessionSettings{
		Host:       "127.0.0.1",
		Port:       33060,
		DbUser:     "wsc",
		DbPassword: "ooxx",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer mySession.Close()

	// New method chaining used for executing an SQL SELECT statement
	// Recommended way for executing queries

	db := mysqlx.Database{}

	employees := db.GetSchema().GetTable("test")

	employees.Select("name", "age").
		Where("name like :param").
		OrderBy("name").Bind(map[string]string{"name": "m%"}).
		Execute()

	// Traditional SQL execution by passing an SQL string
	// This is only available when using a NodeSession
	// It should only be used when absolutely necessary

	mySession.Sql("SELECT name, age " +
		"FROM employee " +
		"WHERE name like ? " +
		"ORDER BY name").Bind("m%").Execute()

}
