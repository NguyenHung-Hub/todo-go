package business

import (
	"context"
	"social_todo/common"
	"social_todo/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type updateItemBusiness struct {
	store UpdateItemStorage
}

func NewUpdateItemBusiness(store UpdateItemStorage) *updateItemBusiness {
	return &updateItemBusiness{store: store}
}

func (busi *updateItemBusiness) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {

	data, err := busi.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return common.ErrEntityDeleted(model.EntityName, model.ErrItemDeleted)
	}

	if err := busi.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
