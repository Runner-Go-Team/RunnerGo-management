package rao

type ListReportsReq struct {
	TeamID int64 `form:"team_id"`
	Page   int   `form:"page"`
	Size   int   `form:"size"`
}

type ListReportsResp struct {
	Reports []*Report `json:"reports"`
	Total   int64     `json:"total"`
}

type Report struct {
	ReportID int64  `json:"report_id"`
	Name     string `json:"name"`
	Mode     int32  `json:"mode"`
	Status   int32  `json:"status"`
}