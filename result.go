package mysqlx

type BaseResult struct {
}

func (this *BaseResult) GetWarningsCount() int {
	return 0
}

func (this *BaseResult) GetWarnings() {

}

type Result struct {
	BaseResult
}

func (this *Result) GetAffectedItemsCount() int {
	return 0
}

func (this *Result) GetLastInsertId() int {
	return 0
}

//return UUID String
func (this *Result) GetLastDocument() int {
	return 0
}

type DocResult struct {
	BaseResult
}

func (this *DocResult) FetchOne() {

}

func (this *DocResult) FetchAll() {

}

type RowResult struct {
	BaseResult
}

func (this *RowResult) FetchOne() Row {
	return Row{}
}

func (this *RowResult) FetchAll() {

}

func (this *RowResult) GetColumnCount() int {
	return 0
}

func (this *RowResult) GetColumns() {

}

func (this *RowResult) GetColumnNames() []string {
	return []string{}
}

type SqlResult struct {
	RowResult
}

func (this *SqlResult) GetAffectedRowsCount() int {
	return 0
}

func (this *SqlResult) GetLastInsertId() int {
	return 0
}

func (this *SqlResult) HasData() bool {
	return true
}

func (this *SqlResult) NextResultt() bool {
	return true
}
