package logic

import (
	"context"

	"fmt"
	"strconv"

	"commits/internal/svc"
	"commits/internal/types"

	"github.com/imroc/req/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

type CommitsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommitsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommitsListLogic {
	return &CommitsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Rslt struct {
	Repo       string `json:"repo"`
	Branch     string `json:"branch"`
	Total      string `json:"total"`
	DeptNameT0 string `json:"dept_name_t0"`
	DeptNameT1 string `json:"dept_name_t1"`
	DeptNameT2 string `json:"dept_name_t2"`
	DeptNameT3 string `json:"dept_name_t3"`
}
type IDataResp struct {
	Code int
	Msg  string
	Data struct {
		Count    string
		CountAll string
		Rslt     []Rslt
	}
}

func (l *CommitsListLogic) CommitsList(reqq *types.ListRepoReq) (resp *[]types.RepoInfo, err error) {
	// todo: add your logic here and delete this line
	client := req.C()
	url := fmt.Sprintf("http://10.85.128.137/idata/api/v1/stat_meta_data/active_repo_by_dept_total30_52780?commit_time=%s&commit_time%%2B1=%s", reqq.StartDate, reqq.EndDate)
	var ir IDataResp
	_, err = client.R().SetHeader("apikey", "af766ad795504d5ea4a1c90acba54f27").SetSuccessResult(&ir).
		Get(url)
	resp = new([]types.RepoInfo)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	// resp.Code = 200
	// resp.Msg = trace.TraceIDFromContext(l.ctx)
	resp = covert(ir)
	return resp, nil
}

func covert(result IDataResp) *[]types.RepoInfo {
	list := make([]types.RepoInfo, 0)
	if result.Data.Count > "0" {
		for _, v := range result.Data.Rslt {
			t, _ := strconv.Atoi(v.Total)
			tmp := types.RepoInfo{
				Repo:       v.Repo,
				Branch:     v.Branch,
				Total:      t,
				DeptNameT0: v.DeptNameT0,
				DeptNameT1: v.DeptNameT1,
				DeptNameT2: v.DeptNameT2,
				DeptNameT3: v.DeptNameT3,
			}
			list = append(list, tmp)
		}
	}
	return &list
}
