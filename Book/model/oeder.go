package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func FindAllOrderWithUser(uid int64) []Order {
	OrderList := make([]Order, 0)
	a := GetRedis("OrderList:" + strconv.FormatInt(uid, 10))
	err := json.Unmarshal([]byte(a), &OrderList)
	if err != nil {
		fmt.Println(err)
	}
	if OrderList != nil {
		fmt.Println(11)
		return OrderList
	}
	ConnectMysql.Raw("select * from order where uid = ?  ", uid).Scan(&OrderList)
	JsonBook, err1 := json.Marshal(OrderList)
	err2 := SetRedis("OrderList:"+strconv.FormatInt(uid, 10), JsonBook, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return OrderList
}
func FindOrder(orderid int64) Order {
	var order Order
	a := GetRedis("Order:" + strconv.FormatInt(orderid, 10))
	err := json.Unmarshal([]byte(a), &order)
	if err != nil {
		fmt.Println(err)
	}
	if order.ID > 0 {
		fmt.Println(11)
		return order
	}
	ConnectMysql.Raw("select * from order where orderid = ?  ", orderid).Scan(&order)
	JsonBook, err1 := json.Marshal(order)
	err2 := SetRedis("OrderList:"+strconv.FormatInt(orderid, 10), JsonBook, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return order
}
func AddOrder(order Order) error {
	err := ConnectMysql.Table("order").Create(&order).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	JsonBook, err1 := json.Marshal(order)
	err2 := SetRedis("order:orderid:"+strconv.FormatInt(order.orderid, 10), JsonBook, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		return err1
	}
	return nil
}
func DelOrder(id int64) {
	ConnectMysql.Raw("delete from order where id =?", id)
}
