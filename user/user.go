package user

import (
	"gopkg.in/mgo.v2/bson"
)
type User struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string    //用户名
	Password string   //密码
	Email string    //邮箱
	CareSku []string    //关注价格变动的sku
	Repository UserRepository `inject:""`
}


type UserRepository interface {
	Save(user User) (u User, err error)
	Delete(name string) error
	Update(user User) (u User,err error)
	Find(name string) User
}


func (u User) Save() (User,error){

	return u.Repository.Save(u)
}

func (u User) Update() (User,error){

	return u.Repository.Update(u)
}

func (u User) Delete() error{

	return u.Repository.Delete(u.Name)
}

func (u User) Find() User{
	return u.Repository.Find(u.Name)
}