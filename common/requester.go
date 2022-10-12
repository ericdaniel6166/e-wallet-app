package common

type Requester interface {
	GetUsername() string
	GetEmail() string
	//GetRole() string
}
