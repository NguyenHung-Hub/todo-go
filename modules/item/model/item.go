package model

import (
	"errors"
	"social_todo/common"
)

const (
	EntityName = "Item"
)

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrItemDeleted  = errors.New("item is deleted")
)

type TodoItem struct {
	common.SQLModelBase
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreate struct {
	Id          int         `json:"-" gorm:"column:id;"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemCreate) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
