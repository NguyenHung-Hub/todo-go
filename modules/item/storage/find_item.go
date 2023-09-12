package storage

import (
	"context"
	"social_todo/common"
	"social_todo/modules/item/model"

	"gorm.io/gorm"
)

func (sql *sqlStorage) GetItem(ctx context.Context, conds map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := sql.db.Where(conds).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
