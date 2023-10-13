package table

import (
	"github.com/zzopen/mysqldoc/src/common/response"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zzopen/mysqldoc/src/internal/logic/table"
	"github.com/zzopen/mysqldoc/src/internal/svc"
)

func InfoHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := table.NewInfoLogic(c.Request.Context(), svcCtx)
		resp, err := l.Info()
		if err != nil {
			c.JSON(http.StatusOK, response.FailWithMsg(err.Error()))
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
