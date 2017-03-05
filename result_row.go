package mysqlx

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

func (this *RowResult) GetColumns() []Column {
	return []Column{}
}

func (this *RowResult) GetColumnNames() []string {
	return []string{}
}

//
func (this *RowResult) GetFirstAutoIncrementValue() string {
	return ""
}

func (this *RowResult) GetAutoIncrementValues() string {
	return ""
}
