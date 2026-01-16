// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta            `orm:"table:users, do:true"`
	Id                any         //
	UserLogin         any         //
	UserPass          any         //
	UserNicename      any         //
	UserEmail         any         //
	UserUrl           any         //
	UserRegistered    *gtime.Time //
	UserActivationKey any         //
	UserStatus        any         //
	DisplayName       any         //
}
