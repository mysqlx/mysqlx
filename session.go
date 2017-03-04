package mysqlx

type BaseSession struct {
}

//Browse Function
func (this *BaseSession) GetSchemas() []Schema {
	return []Schema{}
}

func (this *BaseSession) GetSchema(schemaName string ) Schema {
	return Schema{}
}

func (this *BaseSession) GetDefaultSchema() Schema {
	return Schema{}
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

type XSession struct {
	BaseSession
}

type NodeSession struct {
	BaseSession
}

func (this *NodeSession) SQL(sql string) {

}

func (this *NodeSession) QuoteName() string {
	return ""
}

//setCurrentSchema
func (this *NodeSession) SetCurrentSchema(schemaName string)  Schema {
	return nil
}

