package handler

import (
	"github.com/gin-gonic/gin"

	"kp-management/internal/pkg/biz/errno"
	"kp-management/internal/pkg/biz/jwt"
	"kp-management/internal/pkg/biz/response"
	"kp-management/internal/pkg/dal/rao"
	"kp-management/internal/pkg/logic/variable"
)

// SaveVariable 保存变量
func SaveVariable(ctx *gin.Context) {
	var req rao.SaveVariableReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	if err := variable.SaveVariable(ctx, &req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

// DeleteVariable 删除变量
func DeleteVariable(ctx *gin.Context) {
	var req rao.DeleteVariableReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	if err := variable.DeleteVariable(ctx, req.TeamID, req.VarID); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

// ListGlobalVariables 变量列表
func ListGlobalVariables(ctx *gin.Context) {
	var req rao.ListVariablesReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	v, cnt, err := variable.ListGlobalVariables(ctx, req.TeamID, req.Size, (req.Page-1)*req.Size)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.ListVariablesResp{Variables: v, Total: cnt})
	return
}

// SyncGlobalVariables 同步变量
func SyncGlobalVariables(ctx *gin.Context) {
	var req rao.SyncVariablesReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	if err := variable.SyncGlobalVariables(ctx, req.TeamID, req.Variables); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

func ListSceneVariables(ctx *gin.Context) {
	var req rao.ListSceneVariablesReq
	if err := ctx.ShouldBind(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	v, cnt, err := variable.ListSceneVariables(ctx, req.TeamID, req.SceneID, req.Size, (req.Page-1)*req.Size)
	if err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.SuccessWithData(ctx, rao.ListVariablesResp{Variables: v, Total: cnt})
	return
}

func SyncSceneVariables(ctx *gin.Context) {
	var req rao.SyncSceneVariablesReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	if err := variable.SyncSceneVariables(ctx, req.TeamID, req.SceneID, req.Variables); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}

func ImportSceneVariables(ctx *gin.Context) {
	var req rao.ImportVariablesReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrParam, err.Error())
		return
	}

	if err := variable.ImportSceneVariables(ctx, &req, jwt.GetUserIDByCtx(ctx)); err != nil {
		response.ErrorWithMsg(ctx, errno.ErrMysqlFailed, err.Error())
		return
	}

	response.Success(ctx)
	return
}
