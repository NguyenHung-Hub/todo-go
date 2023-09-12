package business

import (
	"context"
	"social_todo/common"
	"social_todo/modules/item/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBusiness struct {
	store ListItemStorage
}

func NewListItemBusiness(store ListItemStorage) *listItemBusiness {
	return &listItemBusiness{store: store}
}

func (busi *listItemBusiness) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.TodoItem, error) {

	data, err := busi.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
