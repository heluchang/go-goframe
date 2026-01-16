package service

import (
	"context"
	"fmt"
	v1 "test/api/users/v1"
	"test/internal/dao"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type UserService struct{}

var User = &UserService{}

// 获取用户列表
func (s *UserService) GetList(ctx context.Context, req *v1.GetListReq) (*v1.GetListRes, error) {
	record, err := dao.Users.Ctx(ctx).Fields("*").Limit(req.Page, req.Size).All()
	if err != nil {
		return nil, err // 数据库错误
	}

	// 2. 正确判断：gdb.Record 是 map，用 len(record) == 0 判断是否查到数据
	if len(record) == 0 {
		fmt.Println("未查询到用户记录")
		return &v1.GetListRes{
			Data: []v1.UserItem{},
			Code: 10001,
			Msg:  "message: user not found",
		}, nil
	}

	var userList []v1.UserItem
	for _, item := range record {
		userList = append(userList, v1.UserItem{
			ID:                gconv.Uint64(item["id"]),
			UserLogin:         gconv.String(item["user_login"]),
			UserNicename:      gconv.String(item["user_nicename"]),
			UserEmail:         gconv.String(item["user_email"]),
			UserUrl:           gconv.String(item["user_url"]),
			UserRegistered:    gconv.String(item["user_registered"]),
			UserActivationKey: gconv.String(item["user_activation_key"]),
			UserStatus:        gconv.Int(item["user_status"]),
			DisplayName:       gconv.String(item["display_name"]),
		})
	}
	// 3. 序列化验证：打印最终返回的 JSON
	res := &v1.GetListRes{
		Data: userList,
		Code: 200,
		Msg:  "message: user found",
	}
	jsonStr, _ := gjson.Encode(res)
	fmt.Printf("最终返回客户端的 JSON：%s\n", jsonStr)

	// 7. 封装返回
	return res, nil
}

// 获取指定用户
func (s *UserService) GetById(ctx context.Context, req *v1.GetUserByIdReq) (*v1.GetUserByIdRes, error) {
	// 1. 调用无参数 One()，返回 (gdb.Record, error)
	record, err := dao.Users.Ctx(ctx).Where("id = ?", req.ID).One()

	if err != nil {
		return nil, err // 数据库错误
	}

	// 2. 正确判断：gdb.Record 是 map，用 len(record) == 0 判断是否查到数据
	if len(record) == 0 {
		fmt.Println("未查询到 ID 为", req.ID, "的用户记录")
		return &v1.GetUserByIdRes{
			Data: v1.UserItem{},
			Code: 10001,
			Msg:  "message: user not found",
		}, nil
	}

	// 3. 调试：打印查到的原始数据（确认有值）
	fmt.Printf("数据库查询到的原始数据：%+v\n", record)

	// 4. 核心修复：手动映射字段（避开 gconv 版本问题，100%兼容）
	var userItem v1.UserItem
	// 逐个映射数据库下划线字段到结构体驼峰字段
	userItem.ID = gconv.Uint64(record["id"])                                 // 数据库 ID → 结构体 ID
	userItem.UserLogin = gconv.String(record["user_login"])                  // user_login → UserLogin
	userItem.UserNicename = gconv.String(record["user_nicename"])            // user_nicename → UserNicename
	userItem.UserEmail = gconv.String(record["user_email"])                  // user_email → UserEmail
	userItem.UserUrl = gconv.String(record["user_url"])                      // user_url → UserUrl
	userItem.UserRegistered = gconv.String(record["user_registered"])        // user_registered → UserRegistered
	userItem.UserActivationKey = gconv.String(record["user_activation_key"]) // user_activation_key → UserActivationKey
	userItem.UserStatus = gconv.Int(record["user_status"])                   // user_status → UserStatus
	userItem.DisplayName = gconv.String(record["display_name"])              // display_name → DisplayName

	// 5. 调试：打印转换后的结构体（确认字段有值）
	fmt.Printf("转换后的 UserItem：%+v\n", userItem)
	// 6. 序列化验证：打印最终返回的 JSON
	res := &v1.GetUserByIdRes{
		Data: userItem,
		Code: 200,
		Msg:  "message: user found",
	}
	jsonStr, _ := gjson.Encode(res)
	fmt.Printf("最终返回客户端的 JSON：%s\n", jsonStr)

	// 7. 封装返回
	return res, nil
}

// 更新用户
func (s *UserService) Put(ctx context.Context, req *v1.PutReq) (*v1.PutRes, error) {

	updateData := map[string]interface{}{
		"user_activation_key": 1,
		"user_nicename":       req.Name,
	}

	rows, err := dao.Users.Ctx(ctx).
		Where("id = ? AND user_activation_key = ?", req.ID, 1).
		Update(updateData)
	if err != nil {
		return &v1.PutRes{
			Data: "",
			Code: 10002,
			Msg:  "message: update error",
		}, nil
	}
	fmt.Printf("更新了 %d 条记录\n", rows)

	return &v1.PutRes{
		Data: "",
		Code: 200,
		Msg:  "message: update success",
	}, nil
}

// 删除用户
func (s *UserService) Delete(ctx context.Context, req *v1.DeleteReq) (*v1.DeleteRes, error) {
	_, err := dao.Users.Ctx(ctx).Where("id = ?", req.ID).Delete()
	if err != nil {
		return &v1.DeleteRes{
			Data: "",
			Code: 10003,
			Msg:  "message: deleted error",
		}, nil
	}
	return &v1.DeleteRes{
		Data: "",
		Code: 200,
		Msg:  "message: deleted success",
	}, nil
}

// 新增用户
func (s *UserService) Post(ctx context.Context, req *v1.PostReq) (*v1.PostRes, error) {
	_, err := dao.Users.Ctx(ctx).Insert([]g.Map{
		{"user_email": "alice", "user_nicename": "alice", "user_login": "alice", "user_pass": "$P$B$10002$10002$10002$10002$10002$10002"},
		{"user_email": "bob", "user_nicename": "bob", "user_login": "bob", "user_pass": "$P$B$10003$10003$10003$10003$10003$10003"},
	})

	if err != nil {
		return &v1.PostRes{
			Data: v1.UserItem{},
			Code: 10004,
			Msg:  "message: post error",
		}, nil
	}
	return &v1.PostRes{
		Data: v1.UserItem{
			ID:                gconv.Uint64(req.ID),
			UserNicename:      req.Name,
			UserActivationKey: req.UserActivationKey,
		},
		Code: 200,
		Msg:  "message: post success",
	}, nil
}
