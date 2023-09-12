package ginitem

import (
	"net/http"
	"social_todo/common"
	"social_todo/modules/item/business"
	"social_todo/modules/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSqlStorage(db)
		business := business.NewGetItemBusiness(store)

		data, err := business.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
