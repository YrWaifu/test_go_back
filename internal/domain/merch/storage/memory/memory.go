package memory

import (
	"context"

	merchDomain "github.com/YrWaifu/test_go_back/internal/domain/merch"
)

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

var storage = []merchDomain.Merch{
	{Name: "t-shirt", Price: 80},
	{Name: "cup", Price: 20},
	{Name: "book", Price: 50},
	{Name: "pen", Price: 10},
	{Name: "powerbank", Price: 200},
	{Name: "hoody", Price: 300},
	{Name: "umbrella", Price: 200},
	{Name: "socks", Price: 10},
	{Name: "wallet", Price: 50},
	{Name: "pink-hoody", Price: 500},
}

func (s *Storage) GetByName(ctx context.Context, name string) (merchDomain.Merch, error) {
	for _, v := range storage {
		if v.Name == name {
			return v, nil
		}
	}

	return merchDomain.Merch{}, merchDomain.ErrMerchNotFound
}
