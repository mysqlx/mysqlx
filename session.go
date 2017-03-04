package mysqlx

import "sync"

type BaseSession struct {
}

//Browse Function
func (this *BaseSession) GetSchemas() []Schema {
	return []Schema{}
}

func (this *BaseSession) GetSchema(schemaName string ) Schema {
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

type XSession struct {
	BaseSession
}

type NodeSession struct {
	BaseSession
	schema *Schema
	rwMutex sync.RWMutex

}

// Switch to use schema By schema name
func (this *NodeSession)ExecuteSql(schemaName string)  {

}


//
func (this *NodeSession) Sql(sql string) *_sql {
	return &_sql{}
}

func (this *NodeSession) QuoteName() string {
	return ""
}

//setCurrentSchema
func (this *NodeSession) SetCurrentSchema(schemaName string)  (schema *Schema)  {

	this.rwMutex.Lock()
	//set
	this.schema= &Schema{}
	this.rwMutex.Unlock()

	//read
	this.rwMutex.RLock()
	schema = this.schema
	this.rwMutex.RUnlock()

	return
}

//get defaultSchema name
func (this *NodeSession) GetDefaultSchemaName() string {
	return  this.schema.String()
}

