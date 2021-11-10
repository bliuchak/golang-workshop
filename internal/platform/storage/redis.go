package storage

import "github.com/mediocregopher/radix/v3"

const (
	keyPrefix = "book"
)

type Redis struct {
	client radix.Client
}

func NewRedis(host, port string) (*Redis, error) {
	addr := host + ":" + port

	pool, err := radix.NewPool("tcp", addr, 10)
	if err != nil {
		return nil, err
	}

	return &Redis{client: pool}, nil
}

func (r *Redis) Status() (string, error) {
	var status string
	if err := r.client.Do(radix.Cmd(&status, "PING")); err != nil {
		return status, err
	}

	return status, nil
}

func (r *Redis) GetBook(id string) (Book, error) {
	var exists int

	if err := r.client.Do(radix.Cmd(&exists, "EXISTS", getKey(id))); err != nil {
		return Book{}, err
	}

	if exists == 0 {
		return Book{}, ErrNotFound
	}

	var name string
	if err := r.client.Do(radix.Cmd(&name, "GET", getKey(id))); err != nil {
		return Book{}, err
	}

	return Book{ID: id, Title: name}, nil
}

func (r *Redis) CreateBook(id, name string) error {
	return r.client.Do(radix.Cmd(&name, "SET", getKey(id), name))
}

func getKey(id string) string {
	return keyPrefix + "." + id
}
