package mysqlx

type _tableUpdateStatement struct {
}

func (this *_tableUpdateStatement) Set(tableField string, exprOrLiteral string) *_tableUpdateStatement {
	return this
}

func (this *_tableUpdateStatement) Where(searchCondtionStr string) *_tableUpdateStatement {
	return this
}

func (this *_tableUpdateStatement) OrderBy(sortExprStrList ...string) *_tableUpdateStatement {
	return this
}

func (this *_tableUpdateStatement) Limit(numberOfRows int) *_tableUpdateStatement {
	return this
}

func (this *_tableUpdateStatement) Bind(placeholderValues ...interface{}) *_tableUpdateStatement {
	return this
}

func (this *_tableUpdateStatement) Execute() {

}
