package mysqlx

type table  struct {

}

func (this *table)Where(where string) *table {
	return  this
}

func (this *table)OrderBy(...string) *table {
	return  this
}

func (this *table)bind(...string) *table {
	return  this
}

func (this *table)Execute()RowResult {
	return  RowResult{}
}

