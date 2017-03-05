package mysqlx

type _table struct {
}

func (this *_table) Where(where string) *_table {
	return this
}

func (this *_table) OrderBy(...string) *_table {
	return this
}

func (this *_table) Bind(map[string]string) *_table {
	return this
}

func (this *_table) Execute() RowResult {
	return RowResult{}
}

//get is a chan type is RowResult
func (this *_table) ExecuteAsync(rowResult chan<- *RowResult) {
	go func(rowResult chan<- *RowResult) {
		//TODO

		rowResult <- &RowResult{}
	}(rowResult)
}
