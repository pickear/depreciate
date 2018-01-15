package user

type User struct {
	Id string
	Name string    //用户名
	Password string   //密码
	Email string    //邮箱
	CareSku []string    //关注价格变动的sku
	repository UserRepository `inject:""`
}


type UserRepository interface {
	Save(user User) (u User, err error)
	Delete(name string) error
	Update(user User) (u User,err error)
}


func (u User) Save() (User,error){

	return u.repository.Save(u)
}

func (u User) Update() (User,error){

	return u.repository.Update(u)
}

func (u User) Delete() error{

	return u.repository.Delete(u.Name)
}