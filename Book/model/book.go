package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func FindAllBook(id int64) []Book {
	BookList := make([]Book, 0)
	limit := 5
	a := GetRedis("BookList:" + strconv.FormatInt(id, 10))
	err := json.Unmarshal([]byte(a), &BookList)
	if err != nil {
		fmt.Println(err)
	}
	if BookList != nil {
		fmt.Println(11)
		return BookList
	}
	ConnectMysql.Raw("select * from book where id > ? limit ?  ", id, limit).Scan(&BookList)
	JsonBook, err1 := json.Marshal(BookList)
	err2 := SetRedis("BookList:"+strconv.FormatInt(id, 10), JsonBook, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return BookList
}
func FindBookByName(name string) Book {
	var book Book
	a := GetRedis("Book:Name:" + name)
	err := json.Unmarshal([]byte(a), &book)
	if err != nil {
	}
	if book.Id > 0 {
		return book
	}
	ConnectMysql.Raw("select * from book where name = ?", name).Scan(&book)
	JsonBook, err1 := json.Marshal(book)
	err2 := SetRedis("Book:Name:"+book.Name, JsonBook, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return book
}
func FindBookById(id int64) Book {
	var book Book
	a := GetRedis("Book:Id:" + strconv.FormatInt(id, 10))
	err := json.Unmarshal([]byte(a), &book)
	if err != nil {
		fmt.Println(err)
	}
	if book.Id > 0 {
		return book
	}
	ConnectMysql.Raw("select * from book where id = ?", id).Scan(&book)
	JsonBook, err1 := json.Marshal(book)
	err2 := SetRedis("Book:Id:"+strconv.FormatInt(book.Id, 10), JsonBook, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return book
}
func FindBookByIdForUpdate(id int64) Book {
	var book Book
	a := GetRedis("Book:Id:" + strconv.FormatInt(id, 10))
	err := json.Unmarshal([]byte(a), &book)
	if err != nil {
		fmt.Println(err)
	}
	if book.Id > 0 {
		return book
	}
	ConnectMysql.Raw("select * from book where id = ? FOR UPDATE ", id).Scan(&book)
	JsonBook, err1 := json.Marshal(book)
	err2 := SetRedis("Book:Id:"+strconv.FormatInt(book.Id, 10), JsonBook, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return book
}
func UpDateBook(book Book) bool {
	err := ConnectMysql.Exec("update book set num = ? , d_num = ? , u_time = ? where name = ? ", book.Num, book.DNum, book.UTime, book.Name).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return false
	}
	JsonBook, err1 := json.Marshal(FindBookByName(book.Name))
	err2 := SetRedis("Book:Name:"+book.Name, JsonBook, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return true
}
func DeleteBook(namelist []string) bool {
	if err := ConnectMysql.Transaction(func(tx *gorm.DB) error {
		for _, name := range namelist {
			err := tx.Exec("delete from book where name =?", name).Error
			if err != nil {
				return err
			}
		}
		for _, name := range namelist {
			err := DelRedis("Book:Name:" + name)
			if err != nil {
				fmt.Println(err)
			}
		}
		return nil
	}); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func AddBook(book Book) bool {
	err := ConnectMysql.Table("book").Create(&book).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	JsonBook, err1 := json.Marshal(book)
	err2 := SetRedis("Book:Name:"+book.Name, JsonBook, 3600*time.Second)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
	}
	return true
}
