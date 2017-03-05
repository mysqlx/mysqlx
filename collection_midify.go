package mysqlx

type _collectionMidifyStatement struct {
}

type _collectionMidifyStatementPrePare struct {
}

//this is modify
func (this *_collectionMidifyStatement) Set(collectionField string, exprOrLiteral string) *_collectionMidifyStatementPrePare {
	return &_collectionMidifyStatementPrePare{}
}

func (this *_collectionMidifyStatement) Unset(collectionFields ...string) *_collectionMidifyStatementPrePare {
	return &_collectionMidifyStatementPrePare{}
}
func (this *_collectionMidifyStatement) ArrayInsert(colelctionField string, exprOrLiteral string) *_collectionMidifyStatementPrePare {
	return &_collectionMidifyStatementPrePare{}
}
func (this *_collectionMidifyStatement) ArrayAppend(collecttionField string, exprOrListeral string) *_collectionMidifyStatementPrePare {
	return &_collectionMidifyStatementPrePare{}
}
func (this *_collectionMidifyStatement) ArrayDelete(collectionField string) *_collectionMidifyStatementPrePare {
	return &_collectionMidifyStatementPrePare{}
}

func (this *_collectionMidifyStatementPrePare) Sort(sortExprStrList ...string) *_collectionMidifyStatementPrePare {
	return this
}

func (this *_collectionMidifyStatementPrePare) Limit(numberOfRows int) *_collectionMidifyStatementPrePare {
	return this
}
func (this *_collectionMidifyStatementPrePare) Bind(placeholderValues ...interface{}) *_collectionMidifyStatementPrePare {
	return this
}
func (this *_collectionMidifyStatementPrePare) Execute() {
	return
}
