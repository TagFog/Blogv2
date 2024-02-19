package models

import (
	"server/models/ctype"
)

type AuthModel struct {
	// gorm.Model
	MODEL
	NickName       string           `gorm:"size:36" json:"nick_name"` //昵称
	UserName       string           `gorm:"size:36" json:"user_name"` //用户名
	Password       string           `gorm:"size:128" json:"password"` //密码
	Avatar         string           `grom:"size:128" json:"avatar"`   //头像
	Email          string           `gorm:"size:128" json:"email"`    //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`       //电话
	Addr           string           `gorm:"size:64" json:"addr"`      //地址
	Token          string           `gorm:"size:64" json:"token"`     //其他平台
	IP             string           `gorm:"size:20" json:"ip"`
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`                                                  //角色 1-管理员 2-普通用户 3-游客
	SignStatus     ctype.SignStatus `grom:"type=smallint(6)" json:"sign_status"`                                           //注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:AuthID" json:"-"`                                                    //发布的文章列表
	CollectsModels []ArticleModel   `gorm:"many2many:auth_collects;joinForeignKey:AuthID;JoinRefences:ArticleID" json:"-"` //收藏的文章列表
}
