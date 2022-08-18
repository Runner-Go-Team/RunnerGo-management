package handler

import (
	"kp-management/internal/pkg/biz/errno"
	"kp-management/internal/pkg/biz/response"
	"kp-management/internal/pkg/dal/rao"
	"kp-management/internal/pkg/logic/api"
	"kp-management/internal/pkg/logic/target"

	"github.com/gin-gonic/gin"
)

func SaveTarget(ctx *gin.Context) {
	var req rao.CreateTargetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ParamError, err.Error())
		return
	}

	if err := api.Save(ctx, &req); err != nil {
		response.ErrorWithMsg(ctx, errno.MysqlOperFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

func DeleteTarget(ctx *gin.Context) {
	var req rao.DeleteTargetReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ParamError, err.Error())
		return
	}

	if err := target.Delete(ctx, req.TargetID); err != nil {
		response.ErrorWithMsg(ctx, errno.MysqlOperFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

func ListTarget(ctx *gin.Context) {

}
