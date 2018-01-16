package repository

import(
	"depreciate/user"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type UserRepositoryImpl struct{
	Mgo *mgo.Session `inject:""`
}

func (repository UserRepositoryImpl) Save(user user.User) (user.User, error){
	fmt.Println("user save")
	_,err := repository.collection().Upsert(bson.M{"name":user.Name},user)
	if(err != nil){
		fmt.Println(err)
		return user,err
	}

	err = repository.collection().Find(bson.M{"name":user.Name}).One(&user)
	if(err != nil){
		fmt.Println(err)
		return user,err
	}
	return user,nil
}

func (repository UserRepositoryImpl) Delete(name string) error{
	fmt.Println("user delete")
	err := repository.collection().Remove(bson.M{"name":name})
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func (repository UserRepositoryImpl) Update(user user.User) (user.User,error){
	fmt.Println("user update")

	err := repository.collection().Update(bson.M{"name":user.Name},bson.M{"$set":bson.M{"name":user.Name,"password":user.Password,"email":user.Email}})
	if err != nil{
		fmt.Println(err)
		return user,err
	}
	return user,nil
}

func (repository UserRepositoryImpl) Find(name string) user.User{
	var user user.User
	err := repository.collection().Find(bson.M{"name":name}).One(&user)
	if(err != nil){
		fmt.Println(err)
	}
	return user
}

func (ur UserRepositoryImpl) collection() *mgo.Collection{
	collection := ur.Mgo.DB("depreciate").C("user")
	return collection
}

func NewUserRepository() *UserRepositoryImpl{

	return &UserRepositoryImpl{}
}