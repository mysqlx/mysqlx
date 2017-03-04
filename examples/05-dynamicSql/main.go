package main

import (
	"github.com/mysqlx/mysqlx"
	"log"
)

func main()  {
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

	createTestTable(mySession,"test1")
	createTestTable(mySession,"test2")

	
}

func createTestTable(session *mysqlx.NodeSession,name string) mysqlx.Table  {

	quoted_name := "`"+
		session.GetDefaultSchemaName()+
		"`.`" + name + "`"
	session.Sql("DROP TABLE IF EXISTS"+quoted_name).Execute()


	create := "CREATE TABLE "
	create += quoted_name
	create += "(id INT NOT NULL PRIMARY KEY AUTO_INCREMENT)"

	session.Sql(create).Execute()

	return session.GetDefaultSchema().GetTable()
}
