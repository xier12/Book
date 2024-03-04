package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type BookPhoto struct {
	Bid  int64
	Path []string
}

// 上传个人头像
func InsertBookPhoto(up BookPhoto) bool {
	//指定操作的数据库名以及要操作的数据集合名
	c := ConnectMongoDB.Database("vote").Collection("BookPhoto")
	_, err := c.InsertOne(context.TODO(), up)
	if err != nil {

		log.Fatal(err)
		return false
	}
	return true
}

// 取到个人头像图片地址
func FindBookPhotoByUserId(id int64) []string {
	//指定操作的数据库名以及要操作的数据集合名
	c := ConnectMongoDB.Database("vote").Collection("BookPhoto")
	//获取指定字段的,指定数值的,所有符合要求的文档
	c2, _ := c.Find(context.TODO(), bson.D{{Key: "uid", Value: id}})
	//c2, _ := c.Find(context.Background(), bson.D{{"age", 29}})
	defer c2.Close(context.Background())
	var str []string
	for c2.Next(context.Background()) {
		var result bson.D
		err := c2.Decode(&result)
		if err != nil {
			return []string{}
		}
		//获取整个数据
		fmt.Printf("result: %v\n", result)
		//获取数据的数据部分
		fmt.Printf("result.Map(): %v\n", result.Map())
		//获取数据的数据部分的指定字段
		fmt.Printf("result.Map()[\"name\"]: %v\n", result.Map()["path"])
		str = result.Map()["path"].([]string)
	}
	if err := c2.Err(); err != nil {
		return []string{}
	}
	return str
}

// 更新数据
func UpdateBookPhoto() {
	c := ConnectMongoDB.Database("vote").Collection("BookPhoto")
	update := bson.D{{"$set", bson.D{{Key: "uid", Value: 3}, {Key: "path", Value: "abcdefg"}}}}
	//可以将多条符合要求的数据都更新
	ur, err := c.UpdateOne(context.TODO(), bson.D{{"uid", 2}}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", ur.ModifiedCount)

}

// 删除头像
func DeleteBookPhoto() {
	c := ConnectMongoDB.Database("vote").Collection("BookPhoto")
	//删除所有符合条件的数据
	dr, err := c.DeleteMany(context.TODO(), bson.D{{Key: "uid", Value: 3}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount)
}
