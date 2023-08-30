package store

import (
	"errors"
	"github.com/Zelayan/dts/pkg/store/span"
)

type StoreType = string

const (
	DefaultStoreType = ("memory")
	MysqlStore       = StoreType("mysql")
)

type ShareDaoFactory interface {
	SpanStore() span.Storage
}

type shareDaoFactory struct {
	storeType StoreType
}

func (f *shareDaoFactory) SpanStore() span.Storage {
	switch f.storeType {
	case DefaultStoreType:
		return span.NewMemoryStorage()
	case MysqlStore:
		// TODO mysql
		break
	default:

	}
	return nil
}

func NewDaoFactory(storeType StoreType) (ShareDaoFactory, error) {
	switch storeType {
	case DefaultStoreType:
	case MysqlStore:
		break
	default:
		return nil, errors.New("does not support storage type")
	}

	return &shareDaoFactory{storeType: storeType}, nil
}
