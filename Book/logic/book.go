package logic

import (
	"Book/model"
	"Book/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type BookDateAndPhoto struct {
	book   model.Book
	bppath []string
}

// @Summary 返回全部图书信息
// @description  返回全部图书接囗
// @Tags 图书查阅
// @Accept   json
// @Produce  json
// @Success 200 {object} tools.ECode
// @Router /BookIndex [post]
func FindAllBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		var page int64
		finallybid := context.PostForm("id")
		if finallybid != "" {
			page, _ = strconv.ParseInt(finallybid, 10, 64)
		} else {
			page = 0
		}
		var booklist []model.Book
		booklist = model.FindAllBook(page)
		context.JSON(http.StatusOK, tools.ECode{Data: booklist})
	}
}

// 返回指定图书信息
func FindBookByName() gin.HandlerFunc {
	return func(context *gin.Context) {
		var book model.Book
		var bppath []string
		name := context.Query("name")
		book = model.FindBookByName(name)
		bppath = model.FindBookPhotoByUserId(book.Id)
		context.JSON(http.StatusOK, tools.ECode{Data: BookDateAndPhoto{
			book:   book,
			bppath: bppath,
		},
		})
	}
}

// 修改图书信息
func UpdateBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		var book model.Book
		err := context.ShouldBind(&book)
		fmt.Println(book)
		book.UTime = time.Now().Format("2006-01-02 15:04:05")
		err1 := model.UpDateBook(book)
		fmt.Println(err1)
		if err != nil || !err1 {
			context.JSON(http.StatusInternalServerError, tools.BookErr)
		}
		context.JSON(http.StatusOK, tools.OK)
	}
}

// 添加图书信息
func AddBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		var book model.Book
		err := context.ShouldBind(&book)
		book.Ctime = time.Now().Format("2006-01-02 15:04:05")
		book.UTime = book.Ctime
		book.BUuid, _ = tools.Snow.GenerateID()
		model.AddBook(book)
		if err != nil {
			context.JSON(http.StatusInternalServerError, tools.BookErr)
		}
		context.JSON(http.StatusOK, tools.OK)
	}
}

// 删除选择的图书信息
func DeleteBook() gin.HandlerFunc {
	return func(context *gin.Context) {
		bookStr, _ := context.GetPostFormArray("book[]")
		booklist := make([]string, 0)
		for _, s := range bookStr {
			name := s
			booklist = append(booklist, name)
		}
		err := model.DeleteBook(booklist)
		if !err {
			context.JSON(http.StatusInternalServerError, tools.BookErr)
		}

		context.JSON(http.StatusOK, tools.OK)
	}
}
