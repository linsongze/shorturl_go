package service

type StoreService interface {
	Save(shortCode, url string)
	IncAndGet() int64
	Get(shortCode string) string
}
