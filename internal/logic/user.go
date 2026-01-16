package logic

import (
	"context"
	v1 "test/api/users/v1"
	"test/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

// UserLogic 逻辑层结构体
type UserLogic struct{}

// 创建用户逻辑
func (l *UserLogic) GetById(ctx context.Context, req *v1.GetUserByIdReq) (*v1.GetUserByIdRes, error) {
	// 数据验证
	if req.Name == "" {
		return nil, gerror.New("用户名不能为空")
	}

	// 调用 Service 层处理数据库相关的逻辑
	return service.User.GetById(ctx, req)
}
