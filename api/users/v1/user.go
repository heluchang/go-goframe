package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// type userColumns = dao.Users.Columns()
type UserItem struct {
	ID        uint64 `json:"id"`
	UserLogin string `json:"userLogin"`
	// UserPass          string `json:"userPass"`
	UserNicename      string `json:"userNicename"`
	UserEmail         string `json:"userEmail"`
	UserUrl           string `json:"userUrl"`
	UserRegistered    string `json:"userRegistered"`
	UserActivationKey string `json:"userActivationKey"`
	UserStatus        int    `json:"userStatus"`
	DisplayName       string `json:"displayName"`
}

// ------------------ Get Users ------------------
type GetListReq struct {
	g.Meta `path:"/list" method:"get" summary:"获取用户列表"`
	Page   int `json:"page" in:"query" default:"0" description:"页码"`
	Size   int `json:"size" in:"query" default:"10" description:"每页条数"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	Data   []UserItem `json:"data"`
	Code   int        `json:"code" example:"200"`
	Msg    string     `json:"msg" example:"user list found"`
}

// ------------------ Update Users ------------------
type PutReq struct {
	g.Meta `path:"/users/{id}" tags:"Users" method:"put" summary:"Update a user"`
	ID     int    `json:"id" v:"required"`
	Name   string `json:"name"`
}

type PutRes struct {
	g.Meta `mime:"application/json"`
	Data   string `json:"data"`
	Code   int    `json:"code" example:"200"`
	Msg    string `json:"msg" example:"user updated"`
}

// ------------------ Delete Users ------------------
type DeleteReq struct {
	g.Meta `path:"/users/{id}" tags:"Users" method:"delete" summary:"Delete a user"`
	ID     int `json:"id" v:"required"`
}

type DeleteRes struct {
	g.Meta `mime:"application/json"`
	Data   string `json:"data"`
	Code   int    `json:"code" example:"200"`
	Msg    string `json:"msg" example:"user deleted"`
}

// ------------------ Post Users ------------------
type PostReq struct {
	g.Meta            `path:"/users" tags:"Users" method:"post" summary:"Create a user"`
	ID                string `json:"id"`
	Name              string `json:"name"`
	UserActivationKey string `json:"userActivationKey"`
}

type PostRes struct {
	g.Meta `mime:"application/json"`
	Data   UserItem `json:"data"`
	Code   int      `json:"code" example:"200"`
	Msg    string   `json:"msg" example:"user created"`
}

// GetUserListReq 用户列表请求（查询参数）
type GetUserByIdReq struct {
	g.Meta `path:"/users/{id}" tags:"Users" method:"get" summary:"获取指定用户"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
}

type GetUserByIdRes struct {
	g.Meta `mime:"application/json"`
	Data   UserItem `json:"data"`
	Code   int      `json:"code" example:"200"`
	Msg    string   `json:"msg" example:"user found"`
}
