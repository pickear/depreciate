package repository

import(
	"depreciate/user"
	"fmt"
)
type UserRepositoryImpl struct{
}

func (UserRepositoryImpl) Save(user user.User) (user.User, error){
	fmt.Println("user save")
	return user,nil
}

func (UserRepositoryImpl) Delete(name string) error{
	fmt.Println("user delete")
	return nil
}

func (UserRepositoryImpl) Update(user user.User) (user.User,error){
	fmt.Println("user update")
	return user,nil
}

func NewUserRepository() *UserRepositoryImpl{

	return &UserRepositoryImpl{}
}