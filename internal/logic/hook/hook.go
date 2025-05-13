package hook

import (
	"charon-janus/internal/library/contexts"
	"charon-janus/internal/model/entity"
	"charon-janus/internal/service"
	"charon-janus/utility/location"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sHook struct{}

func NewHook() *sHook {
	return &sHook{}
}

func init() {
	service.RegisterHook(NewHook())
}

func (s *sHook) BeforeServe(r *ghttp.Request) {}

func (s *sHook) AfterOutput(r *ghttp.Request) {

	s.accessLog(r)
}

func (s *sHook) accessLog(r *ghttp.Request) {
	var (
		ctx        = r.Context()
		mCtx       = contexts.Get(ctx)
		request    = ghttp.RequestFromCtx(ctx)
		mReq       = mCtx.Request
		response   = mCtx.Response
		user       = mCtx.User
		takeUpTime = gtime.Now().Sub(gtime.New(r.EnterTime)).Milliseconds() // 耗时
		client     = location.GetClientIp(r)
		getData    = gjson.New(request.URL.Query())
		postData   = gjson.New("{}")
		headerData = gjson.New("{}")
		errorData  = gjson.New("{}")
		errorCode  int
		traceID    string
	)
	if user.Id == 0 {
		return
	}
	// 响应数据
	if response != nil {
		errorCode = response.Code
		errorData = gjson.New(response.Message)
		traceID = gctx.CtxId(ctx)
		if response.Code > 0 {
			errorData = gjson.New(response.Message)
		}
	}

	if reqHeadersBytes, _ := gjson.New(request.Header).MarshalJSON(); len(reqHeadersBytes) > 0 {
		headerData = gjson.New(reqHeadersBytes)
	}

	// post参数
	postData, _ = gjson.DecodeToJson(mReq.Body)

	// post表单
	postForm := gjson.New(gconv.String(request.PostForm)).Map()
	if len(postForm) > 0 {
		for k, v := range postForm {
			postData.MustSet(k, v)
		}
	}
	data := entity.SysLog{
		Id:         0,
		ReqId:      traceID,
		Username:   user.Username,
		Url:        mReq.Path,
		Method:     mReq.Method,
		GetData:    getData,
		PostData:   postData,
		HeaderData: headerData,
		ErrorData:  errorData,
		UserAgent:  r.UserAgent(),
		TakeUpTime: gconv.String(takeUpTime),
		ClientIp:   client,
		Timestamp:  gtime.Now(),
		Code:       errorCode,
		CreatedAt:  nil,
		UpdatedAt:  nil,
	}
	Queue().Push(ctx, data)
}
