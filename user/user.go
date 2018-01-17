package user

import (
	"gopkg.in/mgo.v2/bson"
)
type User struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string    //用户名
	Password string   //密码
	Email string    //邮箱
	Mobile string  //电话号码
	CareSku []string    //关注价格变动的sku
}


type UserRepository interface {
	Save(user User) (User, error)
	Delete(name string) error
	Update(user User) (User,error)
	Find(name string) User
}