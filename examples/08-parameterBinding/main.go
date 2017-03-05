package main

import "github.com/mysqlx/mysqlx"

func main() {

	// Collection.find() function with fixed values

	db := mysqlx.Database{}
	myColl := db.GetSchema().GetCollection("test")

	myDocRes := myColl.Find("name = 15").Execute()

	// Using the .bind() function to bind parameters

	myDocRes2 := myColl.Find("name = :param1 And age = :param2").
		Bind("param1", "kack").Bind("param2", 17).Execute()

	_ = myDocRes
	_ = myDocRes2

	params := map[string]interface{}{
		"name": "clare",
	}

	myColl.Modify("name =:param").
		Set("age", 37).
		Bind(params).
		Execute()

	myColl.Find("a = :=param").Fields(":param as b").Bind("param", "c")

}
