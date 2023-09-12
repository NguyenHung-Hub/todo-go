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

func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		paging.Process()

		store := storage.NewSqlStorage(db)
		business := business.NewListItemBusiness(store)

		result, err := business.ListItem(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
