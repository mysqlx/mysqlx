package mysqlx

type Database struct {

}

func (this *Database)GetSession() XSession  {
	return XSession{}
}

func (this *Database)GetSchema() Schema {
	return Schema{}
}
func (this *Database)GetName() string {
	return ""
}

func (this *Database)ExistsInDatabase() bool {
	return true
}

