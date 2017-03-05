package mysqlx

type BaseResult struct {
}

func (this *BaseResult) GetWarningsCount() int {
	return 0
}

func (this *BaseResult) GetWarnings() []Warning {
	return []Warning{}
}

type Result struct {
	BaseResult
}

func (this *Result) GetAffectedItemsCount() int {
	return 0
}

func (this *Result) GetLastInsertId() int {
	return 0
}

//return UUID String
func (this *Result) GetLastDocument() int {
	return 0
}
