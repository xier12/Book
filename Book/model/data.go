package model

type Book struct {
	Id    int64   `json:"id,omitempty" form:"id" gorm:"id"`
	BUuid int64   `json:"uuid,omitempty" form:"uuid" gorm:"uuid"`
	Name  string  `json:"name,omitempty" form:"name" gorm:"name"`
	Num   int     `json:"num,omitempty" form:"num" gorm:"num"`       // 图书总数
	DNum  int     `json:"d_num,omitempty" form:"d_num" gorm:"d_num"` // 已借出的图书数量
	Ctime string  `json:"c_time,omitempty" form:"c_time" gorm:"c_time"`
	UTime string  `json:"u_time,omitempty" form:"u_time" gorm:"u_time"`
	Price float32 `json:"price" gorm:"price" form:"price"`
}
type BookWithUser struct {
	Id    int64  `json:"id,omitempty" gorm:"id"`
	Uid   int64  `json:"uid,omitempty" gorm:"uid"`
	Bid   int64  `json:"bid,omitempty" gorm:"bid"`
	Num   int    `json:"num,omitempty" gorm:"num"` // 借了多少本书
	CTime string `json:"c_time,omitempty" gorm:"c_time"`
	UTime string `json:"u_time,omitempty" gorm:"u_time"`
}

type User struct {
	Id       int64  `json:"id,omitempty" form:"id" gorm:"id"`
	Uuid     int64  `json:"Uuid,omitempty" form:"Uuid" gorm:"Uuid"`
	Name     string `json:"name,omitempty" form:"name" gorm:"name"`
	Password string `json:"password,omitempty" form:"password" gorm:"password"`
	Ctime    string `json:"c_time,omitempty" form:"c_time" gorm:"c_time"`
	UTime    string `json:"u_time,omitempty" form:"u_time" gorm:"u_time"`
	Power    string `json:"power,omitempty" form:"power" gorm:"power"`
	Email    string `json:"email" gorm:"email" form:"email"`
	Tel      string `json:"tel" gorm:"tel" form:"tel"`
}
type Order struct {
	ID         int64   `json:"id" gorm:"id"`
	Bid        int64   `json:"bid" gorm:"bid"`
	Uid        int64   `json:"uid" gorm:"uid"`
	Num        int64   `json:"num" gorm:"num"`
	Price      float32 `json:"price" gorm:"price"`
	Createtime string  `json:"createtime" gorm:"createtime"`
	Mybool     int64   `json:"mybool" gorm:"mybool"`
	orderid    int64   `json:"orderid" gorm:"orderid"`
}
