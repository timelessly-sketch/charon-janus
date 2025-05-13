package input

type PageReq struct {
	Page int `json:"page" example:"1" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码"`
	Size int `json:"size" example:"10" d:"10" v:"min:1|max:200#每页数量最小值不能低于1|最大值不能大于200" dc:"每页数量"`
}
