package mysqlx

type Warning struct {
	Level   string
	Code    string
	Message string
}

//print function
func (this Warning) String() string {
	//TODO
	return ""
}
