package main

import (
	"github.com/mysqlx/mysqlx"
	"log"
)

func main()  {
	//connect to server on localhost
	mySession,err := mysqlx.GetNodeSession(mysqlx.SessionSettings{
		Host:"127.0.0.1",
		Port:33060,
		DbUser:"wsc",
		DbPassword:"ooxx",
	})

	if err!=nil {
		log.Fatal(err)
	}

	defer mySession.Close()

	// Switch to use schema 'test'
	mySession.ExecuteSql("test")


	// In a NodeSession context the full SQL language can be used
	mySession.Sql("CREATE PROCEDURE my_add_one_procedure " +
		" (INOUT incr_param INT) " +
		"BEGIN " +
		"  SET incr_param = incr_param + 1;" +
		"END").Execute()

	mySession.Sql("SET @my_var = ?").
		Bind(1).
		Execute()

	mySession.Sql("CALL my_add_one_procedure(@my_var)").Execute()

	mySession.Sql("DROP PROCEDURE my_add_one_procedure").Execute()


	// return SqlResult
	myResult := mySession.Sql("SELECT @my_var").Execute()

	//return Row
	row := myResult.FetchOne()
	// Gets the row and prints the first column
	println(row.GetInt(0))

}
