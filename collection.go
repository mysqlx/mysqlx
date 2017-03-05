package mysqlx

type Collection struct {
}

func (this *Collection) Add(documentOrJSON string) *_collectionAddStatement {
	return &_collectionAddStatement{}
}
func (this *Collection) Find(searchConditionStr string) *_collectionFindStatement {
	return &_collectionFindStatement{}
}

func (this *Collection) Modify(searchConditionStr string) *_collectionMidifyStatement {
	return &_collectionMidifyStatement{}
}

func (this *Collection) Remove(searchConditionStr string) *_collectionRemoveStatement {
	return &_collectionRemoveStatement{}
}

func (this *Collection) CreateIndex() {

}

func (this *Collection) DropIndex() {

}

func (this *Collection) GetIndexes() {

}

func (this *Collection) NewDoc() {

}

func (this *Collection) Count() int {
	return 0
}
