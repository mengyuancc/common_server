/**

*/
package models

type DailyWord struct {
	Id           int    `xorm:"not null pk autoincr comment('主键ID') INT(10)" form:"id"`
	Author       string `xorm:"not null comment('作者') VARCHAR(50)" form:"author"`
	Content      string `xorm:"not null comment('内容') VARCHAR(500)" form:"content"`
	SDate        string `xorm:"not null comment('日期') VARCHAR(255)" form:"s_date"`
	IsShow       int    `xorm:"not null default 0 comment('是否展示') INT(10)" form:"is_show"`
	CreatedAt    int    `xorm:"not null default '' comment('创建时间') VARCHAR(10)" form:"-"`
	UpdatedAt    int    `xorm:"not null default '' comment('最后修改时间') VARCHAR(10)" form:"-"`
}
