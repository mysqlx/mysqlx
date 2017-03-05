package mysqlx

//metadata related to the origin and types of results from relational queries

type Columner interface {

}


type Column struct {
	DatabaseName     string
	TableName        string
	TableLabel       string
	ColumnName       string
	ColumnLabel      string
	Type             string
	Length           int
	FractionalDigits int
	NumberSigned     bool
	CollationName    string
	CharacterSetName string
	Padded           bool
}

func (this Column) GetSchemaName() string {
	return this.DatabaseName
}

func (this Column) GetTableName() string {
	return this.TableName
}

func (this Column) GetTableLabel() string {
	return this.GetTableLabel()
}

func (this Column) GetColumnName() string {
	return this.ColumnName
}

func (this Column) GetColumnLabel() string {
	return this.ColumnLabel
}

func (this Column) GetType() string {
	return this.Type
}

func (this Column) GetLength() int {
	return this.Length
}
func (this Column) GetFractionalDigits() int {
	return this.FractionalDigits
}

func (this Column) IsNumberSigned() bool {
	return this.NumberSigned
}

func (this Column) GetCollationName() string {
	return this.CollationName
}

func (this Column) GetCharacterSetName() string {
	return this.CharacterSetName
}
func (this Column) IsPadded() bool {
	return this.Padded
}
