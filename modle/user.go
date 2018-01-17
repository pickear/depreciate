package modle

import (
	"gopkg.in/mgo.v2/bson"
)
type User struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Name string    //用户名
	Password string   //密码
	Email string    //邮箱
	Mobile string  //电话号码
	CareSkus map[string]float64    //关注价格变动的sku
}

type UserRepository interface {
	//保存用户(新增，更新)
	Save(user User) (User, error)

	//根据用户名删除用户
	Delete(name string) error

	//根据用户名更新用户
	Update(user User) (User,error)

	//根据用户名查找用户
	Find(name string) (User,error)
}