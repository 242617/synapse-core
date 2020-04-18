package types

type Lifecycle interface {
	Start() error
	Stop() error
}
