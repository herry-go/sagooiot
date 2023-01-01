// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta        `orm:"table:sys_user, do:true"`
	Id            interface{} //
	UserName      interface{} // 用户名
	UserTypes     interface{} // 系统 system 企业 company
	Mobile        interface{} // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname  interface{} // 用户昵称
	Birthday      interface{} // 生日
	UserPassword  interface{} // 登录密码;cmf_password加密
	UserSalt      interface{} // 加密盐
	UserEmail     interface{} // 用户登录邮箱
	Sex           interface{} // 性别;0:保密,1:男,2:女
	Avatar        interface{} // 用户头像
	DeptId        interface{} // 部门id
	Remark        interface{} // 备注
	IsAdmin       interface{} // 是否后台管理员 1 是  0   否
	Address       interface{} // 联系地址
	Describe      interface{} // 描述信息
	LastLoginIp   interface{} // 最后登录ip
	LastLoginTime *gtime.Time // 最后登录时间
	Status        interface{} // 用户状态;0:禁用,1:正常,2:未验证
	IsDeleted     interface{} // 是否删除 0未删除 1已删除
	CreateBy      interface{} // 创建者
	CreatedAt     *gtime.Time // 创建日期
	UpdateBy      interface{} // 更新者
	UpdatedAt     *gtime.Time // 修改日期
	DeletedBy     interface{} // 删除人
	DeletedAt     *gtime.Time // 删除时间
}
