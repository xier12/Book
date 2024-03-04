package model

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

var lock sync.Mutex

// 借书
func BorrowBooks(bu BookWithUser) bool {
	if err := ConnectMysql.Transaction(func(tx *gorm.DB) error {
		err := tx.Table("bookwithuser").Create(&bu).Error
		if err != nil {
			return err
		}
		err1 := tx.Exec("update book set d_num = d_num+? where id =?", bu.Num, bu.Bid).Error
		if err1 != nil {
			return err1
		}
		return nil
	}); err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

// 乐观锁
func BorrowBooksOptimistic(bu BookWithUser, num int) bool {
	if err := ConnectMysql.Transaction(func(tx *gorm.DB) error {
		if num != FindBookById(bu.Bid).DNum {
			return fmt.Errorf("数据异常重新请求")
		}
		err := tx.Table("bookwithuser").Create(&bu).Error
		if err != nil {
			return err
		}
		err1 := tx.Exec("update book set d_num = d_num+? where id =?", bu.Num, bu.Bid).Error
		if err1 != nil {
			return err1
		}
		return nil
	}); err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

// 悲观锁
func BorrowBooksPessimistic(bu BookWithUser) bool {
	if err := ConnectMysql.Transaction(func(tx *gorm.DB) error {
		if bu.Num > FindBookByIdForUpdate(bu.Bid).DNum {
			return fmt.Errorf("数据异常重新请求")
		}
		err := tx.Table("bookwithuser").Create(&bu).Error
		if err != nil {
			return err
		}
		err1 := tx.Exec("update book set d_num = d_num+? where id =?", bu.Num, bu.Bid).Error
		if err1 != nil {
			return err1
		}
		return nil
	}); err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

// 还书
func ReturnBook(uid, bid int64, num int) bool {
	if err := ConnectMysql.Transaction(func(tx *gorm.DB) error {
		mytime := time.Now().Format("2006-01-02 15:04:05")
		err := tx.Exec("update bookwithuser set u_time=? where uid =? and bid =?", mytime, uid, bid).Error
		if err != nil {
			return err
		}
		err1 := tx.Exec("update book set d_num = d_num-? where id =?", num, bid).Error
		if err1 != nil {
			return err1
		}
		return nil
	}); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func FindAllRecord() []BookWithUser {
	var record []BookWithUser
	err := ConnectMysql.Raw("select * from bookwithuser").Scan(&record).Error
	if err != nil {
		fmt.Println(err)
	}
	return record
}

func FindRecordByUserId(uid int64) []BookWithUser {
	var record []BookWithUser
	err := ConnectMysql.Raw("select * from bookwithuser where uid = ?", uid).Scan(&record).Error
	if err != nil {
		fmt.Println(err)
	}
	return record
}
