package interfaces

type Database interface {
	Count() (int32, error)
}
