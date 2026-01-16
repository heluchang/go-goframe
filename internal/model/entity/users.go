// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id                uint64      `json:"iD"                orm:"ID"                  description:""` //
	UserLogin         string      `json:"userLogin"         orm:"user_login"          description:""` //
	UserPass          string      `json:"userPass"          orm:"user_pass"           description:""` //
	UserNicename      string      `json:"userNicename"      orm:"user_nicename"       description:""` //
	UserEmail         string      `json:"userEmail"         orm:"user_email"          description:""` //
	UserUrl           string      `json:"userUrl"           orm:"user_url"            description:""` //
	UserRegistered    *gtime.Time `json:"userRegistered"    orm:"user_registered"     description:""` //
	UserActivationKey string      `json:"userActivationKey" orm:"user_activation_key" description:""` //
	UserStatus        int         `json:"userStatus"        orm:"user_status"         description:""` //
	DisplayName       string      `json:"displayName"       orm:"display_name"        description:""` //
}
