package accountenum

//go:generate enumer -type=AccountType -json -sql -transform=snake-upper -output=account_type_enumer.go
type AccountType int32

const (
	CheckingAccount AccountType = iota + 1
	SavingsAccount
	CDAccount
	IRAAccount
	BrokerageAccount
)
