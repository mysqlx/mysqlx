package mysqlx

type _table  struct {

}

func (this *_table)Where(where string) *_table {
	return  this
}

func (this *_table)OrderBy(...string) *_table {
	return  this
}

func (this *_table)bind(...string) *_table {
	return  this
}

func (this *_table)Execute()RowResult {
	return  RowResult{}
}

