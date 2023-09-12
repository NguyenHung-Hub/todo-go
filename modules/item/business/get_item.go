package business

import (
	"context"
	"social_todo/common"
	"social_todo/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
}

type getItemBusiness struct {
	store GetItemStorage
}

func NewGetItemBusiness(store GetItemStorage) *getItemBusiness {
	return &getItemBusiness{store: store}
}

func (busi *getItemBusiness) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {

	data, err := busi.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}
