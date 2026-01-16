package users

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

import (
	"context"

	v1 "test/api/users/v1"
)

type IUser interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetById(ctx context.Context, req *v1.GetUserByIdReq) (res *v1.GetUserByIdRes, err error)
	Put(ctx context.Context, req *v1.PutReq) (res *v1.PutRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Post(ctx context.Context, req *v1.PostReq) (*v1.PostRes, error)
}
