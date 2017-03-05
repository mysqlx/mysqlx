package mysqlx

type Viewer interface {
	Count()
	Select(searchFields string)
}

type View struct {

}
