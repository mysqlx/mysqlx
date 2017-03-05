package main

import "github.com/mysqlx/mysqlx"

func main() {
	db := mysqlx.Database{}
	myColl := db.GetSchema().GetCollection("test")

	myColl.Remove("name =:param AND age = :param2 ").Execute()

}
