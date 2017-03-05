package mysqlx

type Table struct {
	Database
}

func (this *Table) Insert(tableFields ...string) *_tableInsertStatement {
	return &_tableInsertStatement{}
}

func (this *Table) Select(projectSearchExprStrList ...string) *_tableSelectStatement {
	return &_tableSelectStatement{}
}

func (this *Table) Delete(searchConditionStr string) *_tableDeleteStatement{
	return &_tableDeleteStatement{}
}

func (this *Table) Update()*_tableUpdateStatement {
	return &_tableUpdateStatement{}
}

func (this *Table) DropIndex() {

}

func (this *Table) GetIndexes() {

}

func (this *Table) Count() int {
	return 0
}
