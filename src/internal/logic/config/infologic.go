package config

import (
	"context"

	"github.com/zzopen/mysqldoc/src/common/response"
	"github.com/zzopen/mysqldoc/src/common/utils"
	"github.com/zzopen/mysqldoc/src/internal/config"
	"github.com/zzopen/mysqldoc/src/internal/svc"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type configResp struct {
	IsDev  bool   `json:"is_dev"`
	IsProd bool   `json:"is_prod"`
	Title  string `json:"title"`
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Info /api/config/info 获取服务端配置
func (l *InfoLogic) Info() (resp *response.ApiResponse, err error) {
	c := &configResp{
		IsDev:  utils.IsDev(),
		IsProd: utils.IsProd(),
		Title:  config.ProjectName,
	}

	return response.SuccessWithData(c), nil
}
