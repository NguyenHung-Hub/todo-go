package ginitem

import (
	"net/http"
	"social_todo/common"
	"social_todo/modules/item/business"
	"social_todo/modules/item/model"
	"social_todo/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSqlStorage(db)
		business := business.NewCreateItemBusiness(store)

		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))

	}
}
