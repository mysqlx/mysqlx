package mysqlx

type TransactionReply struct {
	WarningCount int
	Warnings     []Warning
}

//Transaction Handling
// Start a transaction
func (this *NodeSession) StartTransaction() {

}

// submit  transaction
func (this *NodeSession) Commit() TransactionReply {
	return TransactionReply{}
}

// Rollback  transaction
func (this *NodeSession) Rollback() TransactionReply {
	return TransactionReply{}
}
