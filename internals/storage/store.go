package storage

type store interface {
	Put(key, value string)
	Get(key string) (string, bool)
	Delete(key string)
}
