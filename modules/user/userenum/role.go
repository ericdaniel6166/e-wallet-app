package userenum

//go:generate enumer -type=Role -json -sql -transform=snake-upper -output=role_enumer.go
type Role int32

const (
	RoleUser Role = 1 << iota
	RoleSeller
	RoleAdmin
	RoleAnonymous Role = 0
)
