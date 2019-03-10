package idStore

type IdStore interface {
	GetNewId() (string, error)
}
