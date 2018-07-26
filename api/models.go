package api

import (
	"github.com/wangjun861205/nborm"
)

type User struct {
	ID         *nborm.NBInt      `column:"id" auto_increment:"true" primary_key:"true"`
	Username   *nborm.NBString   `column:"username" unique:"true"`
	Password   *nborm.NBString   `column:"password"`
	Phone      *nborm.NBString   `column:"phone" unique:"true"`
	Status     *nborm.NBInt      `column:"status"`
	SessionID  *nborm.NBString   `column:"sessionid" unique:"true"`
	ExpireTime *nborm.NBDatetime `column:"expire_time"`
	Email      *nborm.NBString   `column:"email" unique:"true"`
}

func (u *User) DatabaseName() string {
	return "bk_dalian"
}

func (u *User) TableName() string {
	return "auth"
}

func (u *User) Rels() []nborm.Rel {
	return nil
}
