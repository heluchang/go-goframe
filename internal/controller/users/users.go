package users

import (
	"context"
	"test/api/users"
	v1 "test/api/users/v1"

	"test/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type User struct{}

var _ users.IUser = (*User)(nil)

// GET /users
func (c *User) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res, err = service.User.GetList(ctx, req)
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln(res)
	return
}

// PUT /users/{id}
func (c *User) Put(ctx context.Context, req *v1.PutReq) (*v1.PutRes, error) {
	res, err := service.User.Put(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DELETE /users/{id}
func (c *User) Delete(ctx context.Context, req *v1.DeleteReq) (*v1.DeleteRes, error) {
	res, err := service.User.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// POST /users
func (c *User) Post(ctx context.Context, req *v1.PostReq) (*v1.PostRes, error) {
	res, err := service.User.Post(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GET /users/{id}
func (c *User) GetById(ctx context.Context, req *v1.GetUserByIdReq) (res *v1.GetUserByIdRes, err error) {
	res, err = service.User.GetById(ctx, req)
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln(res)
	return
}
