package repository

import(
	"depreciate/modle"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type UserRepositoryImpl struct{
	Mgo *mgo.Session `inject:""`
}

func (repository UserRepositoryImpl) Save(user modle.User) (modle.User, error){

	u,err := repository.Find(user.Name)
	if err != nil{
		user.Id = bson.NewObjectId()
		repository.collection().Insert(user)
		return user,err
	}

	for k,v := range user.CareSkus{
		if u.CareSkus == nil{
			u.CareSkus = make(map[string]float64)
		}
		u.CareSkus[k] = v
	}
	_,err = repository.Update(u)
	return user,err
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

func (repository UserRepositoryImpl) Update(user modle.User) (modle.User,error){
	fmt.Println("user update")

	err := repository.collection().Update(bson.M{"name":user.Name},user)
	if err != nil{
		fmt.Println(err)
		return user,err
	}
	return user,nil
}

func (repository UserRepositoryImpl) Find(name string) (modle.User,error){
	var user modle.User
	err := repository.collection().Find(bson.M{"name":name}).One(&user)
	if(err != nil){
		fmt.Println(err)
	}
	return user,err
}

func (ur UserRepositoryImpl) collection() *mgo.Collection{
	collection := ur.Mgo.DB("depreciate").C("user")
	return collection
}

func NewUserRepository() *UserRepositoryImpl{

	return &UserRepositoryImpl{}
}