package store

import (
	"errors"
	"fmt"
	"github.com/Zelayan/dts/cmd/collector/config"
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
	mysql     span.Storage
	es        span.Storage
}

func (f *shareDaoFactory) SpanStore() span.Storage {
	switch f.storeType {
	case DefaultStoreType:
		return f.mem
	case MysqlStore:
		return f.mem
	case EsStore:
		return f.es
	default:

	}
	return nil
}

func NewDaoFactory(config config.Config) (ShareDaoFactory, error) {
	var (
		mem   = span.NewMemoryStorage()
		es    span.Storage
		mysql span.Storage
		err   error
	)

	switch config.StoreType {
	case DefaultStoreType:
	case MysqlStore:
		break
	case EsStore:
		es, err = span.NewEsStore()
		if err != nil {
			return nil, fmt.Errorf("create es client failed:%v", err)
		}
		break
	default:
		return nil, errors.New("does not support storage type")
	}

	return &shareDaoFactory{
		storeType: config.StoreType,
		mem:       mem,
		mysql:     mysql,
		es:        es,
	}, nil
}
