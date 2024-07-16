package database

import (
	"errors"
	"gorm.io/gorm"
)

type gtx struct {
	session *gorm.DB
}

// NewGtx 创建一个数据库事务接口(gorm 实现)
func NewGtx(engine *gorm.DB) Transaction {
	tx := engine.Begin()

	return &gtx{session: tx}
}

// Do 执行事务操作，出现错误会进行Rollback
func (t *gtx) Do(f any) error {
	var err error

	fn, ok := f.(gtxFunc)
	if !ok {
		return errors.New("plz using *gorm.db in tx func")
	}

	if err = fn(t.session); err != nil {
		_ = t.Rollback()

		return err
	}

	return nil
}

// Rollback 回滚事务
func (t *gtx) Rollback() error {
	return t.session.Rollback().Error
}

// Commit 提交事务
func (t *gtx) Commit() error {
	return t.session.Commit().Error
}
