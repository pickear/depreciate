package repository

import(
	"depreciate/user"
)
type UserRepositoryImpl struct{
}

func (UserRepositoryImpl) Save(user user.User) (user.User, error){

	return user,nil
}

func (UserRepositoryImpl) Delete(name string) error{
	
	return nil
}

func (UserRepositoryImpl) Update(user user.User) (user.User,error){

	return user,nil
}