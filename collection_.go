package mysqlx

//private collection
type _collection struct {

}

func (this *_collection)Limit(limit int) *_collection {
	return this
}

func (this *_collection)Bind(...interface{})*_collection  {
	return  this
}

func (this *_collection)Execute() Document {
	return Document{}
}
