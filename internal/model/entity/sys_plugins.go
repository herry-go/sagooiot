// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPlugins is the golang structure for table sys_plugins.
type SysPlugins struct {
	Id        int         `json:"id"        description:"ID"`
	Name      string      `json:"name"      description:"名称"`
	Title     string      `json:"title"     description:"标题"`
	Intro     string      `json:"intro"     description:"介绍"`
	Version   string      `json:"version"   description:"版本"`
	Status    int         `json:"status"    description:"状态"`
	Types     string      `json:"types"     description:"插件类型"`
	Author    string      `json:"author"    description:""`
	StartTime *gtime.Time `json:"startTime" description:""`
}
