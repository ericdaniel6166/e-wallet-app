package enum

//go:generate enumer -type=Currency -json -sql -transform=snake-upper
type Currency int32

const (
	USD Currency = iota + 1
	EUR
	JPY
	CHF
	CNY
	GBP
	AUD
	CAD
	NZD
)
