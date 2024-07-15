package database

import (
	"gorm.io/gorm"
	"xorm.io/xorm"
)

type xtxFunc func(session *xorm.Session) (int64, error)
type gtxFunc func(session *gorm.DB) error

// TxFunc
// todo golang目前不支持范型方法。现在用起来很麻烦. 先把范型的约束留在这里
type TxFunc interface {
	xtxFunc | gtxFunc
}

type Transaction interface {
	Do(any) error
	Rollback() error
	Commit() error
}
