package mysqlx

type _collectionFindStatement struct {
}

func (this *_collectionFindStatement) Fields(projectedDocumentExprStr string) *_collectionFindStatement {
	return this

}

func (this *_collectionFindStatement) GroupBy(searchExprStrList ...string) *_collectionFindStatement {
	return this
}

func (this *_collectionFindStatement) Having(searchConditionStr string) *_collectionFindStatement {
	return this
}

func (this *_collectionFindStatement) Sort(sortExprStrList ...string) *_collectionFindStatement {
	return this

}

func (this *_collectionFindStatement) Limit(numberOfRows int) *_collectionFindStatement {
	return this
}

func (this *_collectionFindStatement) Offset(numberOfRows int) *_collectionFindStatement {
	return this

}

func (this *_collectionFindStatement) Bind(placeholderValues ...interface{}) *_collectionFindStatement {
	return this
}

func (this *_collectionFindStatement) Execute() {

}
