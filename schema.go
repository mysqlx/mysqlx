package mysqlx

type Schema struct {
	Database
}

//Return schema name
func (this *Schema) String() string {
	return "schema name"
}

//
func (this *Schema) GetName() string {
	return this.String()
}

func (this *Schema) GetCollection(collectionName string) *Collection {
	return &Collection{}
}

func (this *Schema) GetCollections() []Collection {
	return []Collection{}
}

func (this *Schema) GetTable(tableName string) Table {
	return Table{}
}

func (this *Schema) GetCollectionAsTable() Table {
	return Table{}
}

func (this *Schema) CreateCollection(collectionName string) *Collection {
	return &Collection{}
}

//ReuseExistingObject
//this function is golang don't
func (this *Schema) CreateCollectionReuseExistingObject(collectionName string) *Collection {
	return &Collection{}
}
