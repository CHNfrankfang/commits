package logic

import (
	"context"

	"commits/internal/svc"
	"commits/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitAnalysisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitAnalysisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitAnalysisLogic {
	return &SubmitAnalysisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitAnalysisLogic) SubmitAnalysis(req *types.PathRepo) (resp *types.RepoAnalysisResp, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.RepoAnalysisResp)
	resp.Resutl = req.Repo + req.Name + req.Status + req.Foo
	return resp, nil
}
