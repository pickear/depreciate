package repository

import (
	"testing"
	"depreciate/modle"
	"github.com/facebookgo/inject"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestUserSave(t *testing.T){
	assert := assert.New(t)
	var graph inject.Graph

	userRepository := NewUserRepository()
	session := MongoConnect()
	defer session.Close()
	err := graph.Provide(
		&inject.Object{Value: userRepository},
		&inject.Object{Value: session},
	)

	if(err != nil){
		fmt.Println(err)
	}

	if err = graph.Populate();err != nil{
		fmt.Println(err)
	}
	
	u := modle.User{Name:"Dylan",Password:"123",Email:"54646@qq.com",CareSkus:map[string]float64{"J_1997796332":3398.0}}
	_,err = userRepository.Save(u)

	assert.Equal(err,nil,"save user err")
}


func TestUserFind(t *testing.T){
	assert := assert.New(t)
	var graph inject.Graph

	userRepository := NewUserRepository()
	session := MongoConnect()
	defer session.Close()
	err := graph.Provide(
		&inject.Object{Value: userRepository},
		&inject.Object{Value: session},
	)

	if(err != nil){
		fmt.Println(err)
	}

	if err = graph.Populate();err != nil{
		fmt.Println(err)
    }

	user,_ := userRepository.Find("Dylan")
	assert.Equal(user.Name,"Dylan","find user dylan err")
}

func TestUserUpdate(t *testing.T){
	assert := assert.New(t)
	var graph inject.Graph

	userRepository := NewUserRepository()
	session := MongoConnect()
	defer session.Close()
	err := graph.Provide(
		&inject.Object{Value: userRepository},
		&inject.Object{Value: session},
	)

	if(err != nil){
		fmt.Println(err)
	}

	if err = graph.Populate();err != nil{
		fmt.Println(err)
    }

	u,_ := userRepository.Find("Dylan")
	u.Password = "456"
	u,err = userRepository.Update(u)
	assert.Equal(err,nil,"update user err")
}
