package store

import (
	"errors"
	"github.com/Zelayan/dts/cmd/colletcor/config"
	"github.com/Zelayan/dts/pkg/store/span"
)

const (
	DefaultStoreType = "memory"
	MysqlStore       = "mysql"
	EsStore          = "es"
)

type ShareDaoFactory interface {
	SpanStore() span.Storage
}

type shareDaoFactory struct {
	storeType config.StoreType
	mem       span.Storage
}

func (f *shareDaoFactory) SpanStore() span.Storage {
	switch f.storeType {
	case DefaultStoreType:
		return f.mem
	case MysqlStore:
		// TODO mysql
		break
	case EsStore:
		return span.NewEsStore()
	default:

	}
	return nil
}

func NewDaoFactory(storeType config.StoreType) (ShareDaoFactory, error) {
	switch storeType {
	case DefaultStoreType:
	case MysqlStore:
		break
	case EsStore:
		break
	default:
		return nil, errors.New("does not support storage type")
	}

	return &shareDaoFactory{
		storeType: storeType,
		mem:       span.NewMemoryStorage(),
	}, nil
}
