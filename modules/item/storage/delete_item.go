package storage

import (
	"context"
	"social_todo/common"
	"social_todo/modules/item/model"

	"gorm.io/gorm"
)

func (sql *sqlStorage) DeleteItem(ctx context.Context, conds map[string]interface{}) error {

	deletedStatus := model.ItemStatusDeleted
	if err := sql.db.Table(model.TodoItem{}.TableName()).
		Where(conds).
		Updates(map[string]interface{}{"status": deletedStatus.String()}).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrDB(err)
	}

	return nil
}
