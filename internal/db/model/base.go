package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"time"
)

type BaseModel struct {
	ID        string `json:"id" gorm:"primaryKey" size:"10"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDel     bool
}

type Base interface {
	BeforeCreate() error
	BeforeUpdate() error
	BeforeDelete() error
	ToString() string
}

func (bm *BaseModel) beforeCreate() error {
	s, err := gonanoid.New(10)
	if err != nil {
		return err
	} else {
		bm.ID = s
	}
	bm.CreatedAt = time.Now()
	return nil
}

func (bm *BaseModel) beforeUpdate() {
	bm.UpdatedAt = time.Now()
}

func (bm *BaseModel) beforeDelete() {
	bm.IsDel = true
	bm.DeletedAt = time.Now()
}
