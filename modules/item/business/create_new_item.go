package business

import (
	"context"
	"social_todo/common"
	"social_todo/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreate) error
}

type createItemBusiness struct {
	store CreateItemStorage
}

func NewCreateItemBusiness(store CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{store: store}
}

func (busi *createItemBusiness) CreateNewItem(ctx context.Context, data *model.TodoItemCreate) error {

	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := busi.store.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}
