package mysqlx

type BaseSession struct {
}

//Browse Function
func (this *BaseSession) GetSchemas() []Schema {
	return []Schema{}
}

func (this *BaseSession) GetSchema(schemaName string) Schema {
	return Schema{}
}

func (this *BaseSession) GetDefaultSchema() *Schema {
	return &Schema{}
}

func (this *BaseSession) CreateSchema() Schema {
	return Schema{}
}

func (this *BaseSession) DropSchema(schema string) {

}

func (this *BaseSession) DropCollection(schema string, collection string) {

}

func (this *BaseSession) DropTable(schema string, table string) {

}

func (this *BaseSession) GetUri() {

}

func (this *BaseSession) Close() {

}
