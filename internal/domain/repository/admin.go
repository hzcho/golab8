package repository

type Admin interface {
	Get(id uint64) bool
}
