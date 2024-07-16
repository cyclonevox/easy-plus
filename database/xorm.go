package database

import (
	"errors"
	"xorm.io/xorm"
)

type xtx struct {
	session *xorm.Session
}

// NewXtx 创建一个数据库事务接口(xorm 实现)
func NewXtx(engine *xorm.Engine) (Transaction, error) {
	session := engine.NewSession()
	if err := session.Begin(); err != nil {
		return nil, err
	}

	return &xtx{session: session}, nil
}

// Do 执行事务操作，出现错误会进行Rollback
func (t *xtx) Do(f any) error {
	var err error

	fn, ok := f.(xtxFunc)
	if !ok {
		return errors.New("plz using *xorm.Session in tx func")
	}
	if _, err = fn(t.session); err != nil {
		_ = t.Rollback()

		return err
	}

	return nil
}

// Rollback 回滚事务
func (t *xtx) Rollback() error {
	return t.session.Rollback()
}

// Commit 提交事务
func (t *xtx) Commit() error {
	return t.session.Commit()
}
