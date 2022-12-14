package rao

type SaveTeamReq struct {
	TeamID int64  `json:"team_id"`
	Name   string `json:"name"`
}

type SaveTeamResp struct {
	TeamID int64 `json:"team_id"`
}

type ListTeamReq struct {
}

type ListTeamResp struct {
	Teams []*Team `json:"teams"`
}

type Team struct {
	Name            string `json:"name"`
	Type            int32  `json:"type"`
	Sort            int32  `json:"sort"`
	TeamID          int64  `json:"team_id"`
	RoleID          int64  `json:"role_id"`
	CreatedUserID   int64  `json:"created_user_id"`
	CreatedUserName string `json:"created_user_name"`
	CreatedTimeSec  int64  `json:"created_time_sec"`
	Cnt             int64  `json:"cnt"`
}

type ListMembersReq struct {
	TeamID int64 `form:"team_id" binding:"required,gt=0"`
}

type ListMembersResp struct {
	Members []*Member `json:"members"`
}

type Member struct {
	UserID         int64  `json:"user_id"`
	Avatar         string `json:"avatar"`
	Email          string `json:"email"`
	Nickname       string `json:"nickname"`
	RoleID         int64  `json:"role_id"`
	InviteUserID   int64  `json:"invite_user_id"`
	InviteUserName string `json:"invite_user_name"`
	JoinTimeSec    int64  `json:"join_time_sec,omitempty"`
}

type GetInviteMemberURLReq struct {
	TeamID int64 `form:"team_id" binding:"required,gt=0"`
	RoleID int64 `form:"role_id" binding:"required,gt=0"`
}

type GetInviteMemberURLResp struct {
	URL     string `json:"url"`
	Expired int64  `json:"expired"`
}

type CheckInviteMemberURLReq struct {
	TeamID int64 `json:"team_id" binding:"required,gt=0"`
	RoleID int64 `json:"role_id" binding:"required,gt=0"`
	//Email  string `json:"email" binding:"required,gt=0"`
}

type CheckInviteMemberURLResp struct {
}

type InviteMemberReq struct {
	TeamID  int64           `json:"team_id" binding:"required,gt=0"`
	Members []*InviteMember `json:"members"`
	//MemberEmail []string `json:"member_email"`
}

type InviteMember struct {
	Email  string `json:"email"`
	RoleID int64  `json:"role_id"`
}

type InviteMemberResp struct {
	RegisterNum      int      `json:"register_num"`
	UnRegisterNum    int      `json:"un_register_num"`
	UnRegisterEmails []string `json:"un_register_emails"`
}

type RoleUserReq struct {
	TeamID int64 `json:"team_id" binding:"required,gt=0"`
	RoleID int64 `json:"role_id" binding:"required,oneof=2 3"`
	UserID int64 `json:"user_id" binding:"required,gt=0"`
}

type RoleUserResp struct {
}

type RemoveMemberReq struct {
	TeamID   int64 `json:"team_id" binding:"required,gt=0"`
	MemberID int64 `json:"member_id" binding:"required,gt=0"`
}

type RemoveMemberResp struct {
}

type QuitTeamReq struct {
	TeamID int64 `json:"team_id" binding:"required,gt=0"`
}

type QuitTeamResp struct {
}

type GetTeamRoleReq struct {
	TeamID int64 `form:"team_id"`
}

type GetTeamRoleResp struct {
	RoleID int64 `json:"role_id"`
}

type DisbandTeamReq struct {
	TeamID int64 `json:"team_id"`
}

type DisbandTeamResp struct {
}

type TransferTeamReq struct {
	TeamID   int64 `json:"team_id"`
	ToUserID int64 `json:"to_user_id"`
}

type TransferTeamResp struct {
}
