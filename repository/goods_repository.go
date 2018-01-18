package repository

import(
	"depreciate/model"
	"depreciate/util"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)
type GoodsRepositoryImpl struct{
	Mgo *mgo.Session `inject:""`
}

func (repository GoodsRepositoryImpl) Save(goods model.Goods) (model.Goods, error){

	g,err := repository.Find(goods.Sku)
	if err != nil{
		goods.Id = bson.NewObjectId()
		err = repository.collection().Insert(goods)
		return goods,err
	}
	
	
	g.Prices = append(g.Prices,goods.Prices...)
	//去掉价格一样，时间是同一天的重复价格项
	repository.distinctPrice(&g)
	err = repository.collection().Update(bson.M{"sku":goods.Sku},g)
	if(err != nil){
		fmt.Println(err)
		return g,err
	}
	return g,nil
}

func (repository GoodsRepositoryImpl) Delete(sku string) error{
	err := repository.collection().Remove(bson.M{"sku":sku})
	if err != nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func (repository GoodsRepositoryImpl) Find(sku string) (model.Goods,error){
	var goods model.Goods
	err := repository.collection().Find(bson.M{"sku":sku}).One(&goods)
	return goods,err
}

func (repository GoodsRepositoryImpl) Search(name string) []model.Goods{
	var goodss []model.Goods
	err := repository.collection().Find(bson.M{"name":name}).All(&goodss)
	if(err != nil){
		fmt.Println(err)
	}
	return goodss
}

func (gr GoodsRepositoryImpl) collection() *mgo.Collection{
	collection := gr.Mgo.DB("depreciate").C("goods")
	return collection
}

//去重价格(如果价格一样，时间是同一天的价格条目就去掉)
func (gr GoodsRepositoryImpl) distinctPrice(googds *model.Goods){

	prices := make([]model.Price,0)
	//tmpMap := make(map[string]interface{})
	tmpMap := make(map[string]float64)

	for _,price := range googds.Prices{
		
		pTime,err := time.Parse(util.TimeLayout,price.Time)
		if err != nil{
			fmt.Printf("时间[%s]格式化失败\n",price.Time)
			return
		}
		key := pTime.Format(util.DayLayout)
		if p,ok := tmpMap[key];!ok && p != price.Price{
			prices = append(prices,price)
			//tmpMap[key] = struct{}{}
			tmpMap[key] = price.Price
		}
	}
	fmt.Println(prices)
	googds.Prices = prices
}

func NewGoodsRepository() *GoodsRepositoryImpl{

	return &GoodsRepositoryImpl{}
}