package entity

// Entity should be implemented by entities
type Entity interface {
	GetID() string
	String() string
	SetID(string)
}
