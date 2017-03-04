package mysqlx

type Schema struct {
	Database
	
}
//Return schema name
func (this *Schema)String() string {
	return  "schema name"
}

func (this *Schema)GetCollection(collectionName string)Collection  {
	return Collection{}
}

func (this *Schema)GetCollections()[]Collection  {
	return  []Collection{}
}

func (this *Schema)GetTable()Table  {
	return  Table{}
}

func (this *Schema)GetCollectionAsTable()Table  {
	return  Table{}
}

func (this *Schema) CreateCollection() {

}
