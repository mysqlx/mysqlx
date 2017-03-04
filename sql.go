package mysqlx

type sql struct {
	
}

func (this *sql)Execute() Result {
	return Result{}
}


func (this *sql)Bind(...interface{}) *sql {
	return this
}

