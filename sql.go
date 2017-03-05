package mysqlx

type SqlExecute struct {
}

func (this *SqlExecute) Execute() SqlResult {
	return SqlResult{}
}

func (this *SqlExecute) Bind(...interface{}) *SqlExecute {
	return this
}

type SQLPlaceholderValues struct {
}

type SQLPlaceholderName struct {
}

type SQLLiteral struct {
}
