package business

import (
	"context"
	"social_todo/common"
	"social_todo/modules/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, condition map[string]interface{}) error
}

type deleteItemBusiness struct {
	store DeleteItemStorage
}

func NewDeleteItemBusiness(store DeleteItemStorage) *deleteItemBusiness {
	return &deleteItemBusiness{store: store}
}

func (busi *deleteItemBusiness) DeleteItemById(ctx context.Context, id int) error {

	data, err := busi.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	if err := busi.store.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
