package mysqlx

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

func (this *SqlResult) NextResult() bool {
	return true
}
