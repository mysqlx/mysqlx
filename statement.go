package mysqlx

type Statement interface {
	Bind()
	ClearBindings()
	Execute()
	ExecuteAsync()
}

type AddStatement interface {
	Statement
	Add()
}
