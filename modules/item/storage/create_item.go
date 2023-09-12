package storage

import (
	"context"
	"social_todo/common"
	"social_todo/modules/item/model"
)

func (sql *sqlStorage) CreateItem(ctx context.Context, data *model.TodoItemCreate) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
