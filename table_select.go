package mysqlx

type _tableSelectStatement struct {

}

func (this *_tableSelectStatement)Where(searchConditionStr string)*_tableSelectStatement  {
	return  this
}


func (this *_tableSelectStatement)GroupBy(searchExprStrList string)*_tableSelectStatement  {
	return  this
}



func (this *_tableSelectStatement)Having(searchConditionStr string)*_tableSelectStatement  {
	return  this
}

func (this *_tableSelectStatement)OrderBy(orderExprStrList ...string)*_tableSelectStatement  {
	return  this
}


func (this *_tableSelectStatement)Limit(numberOfRows int)*_tableSelectStatement  {
	return  this
}



func (this *_tableSelectStatement)Offset(numberOfRows int)*_tableSelectStatement  {
	return  this
}

func (this *_tableSelectStatement)Bind(placeholderValues ...interface{})*_tableSelectStatement  {
	return  this
}


func (this *_tableSelectStatement)Execute()  {

}




