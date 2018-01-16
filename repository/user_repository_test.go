package repository

import (
	"testing"
	"depreciate/user"
	"github.com/facebookgo/inject"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T){
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
	
	_,err = userRepository.Save(user.User{Name:"Dylan",Password:"123",Email:"54646@qq.com"})

	assert.Equal(err,nil,"save user err")
}


func TestFind(t *testing.T){
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

	user := userRepository.Find("Dylan")
	assert.Equal(user.Name,"Dylan","find user dylan err")
}
