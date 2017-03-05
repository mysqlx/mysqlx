package mysqlx

type _sql struct {
}

func (this *_sql) Execute() SqlResult {
	return SqlResult{}
}

func (this *_sql) Bind(...interface{}) *_sql {
	return this
}
