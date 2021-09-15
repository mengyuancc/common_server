
/**
 * 用户模型
 */
package models

type UserInfo struct {
	Id           int    `xorm:"not null pk autoincr comment('主键ID') INT(10)" form:"id"`
	Name         string `xorm:"not null comment('中文名') VARCHAR(50)" form:"name"`
	SysCreated   int    `xorm:"not null default 0 comment('创建时间') INT(10)" form:"sys_created"`
	SysUpdated   int    `xorm:"not null default 0 comment('最后修改时间') INT(10)" form:"sys_updated"`
}
