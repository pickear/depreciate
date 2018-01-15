package user

type User struct {
	Id string
	Name string    //用户名
	Password string   //密码
	Email string    //邮箱
	CareSku []string    //关注价格变动的sku
}


type UserRepository interface {
	Save(user User) (u User, err error)
	Delete(name string) error
	Update(user User) (u User,err error)
}