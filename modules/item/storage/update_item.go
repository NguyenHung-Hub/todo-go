package storage

import (
	"context"
	"social_todo/common"
	"social_todo/modules/item/model"

	"gorm.io/gorm"
)

func (sql *sqlStorage) UpdateItem(ctx context.Context, conds map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {

	if err := sql.db.Where(conds).Updates(dataUpdate).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrDB(err)
	}

	return nil
}
