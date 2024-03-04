package logic

import (
	"Book/model"
	"Book/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpLoadUserPhoto() gin.HandlerFunc {
	return func(context *gin.Context) {
		//file, _ := context.FormFile("photofile")
		filess, _ := context.MultipartForm()
		files := filess.File["photofiles"]
		//a为从jwt中取出的1
		//a := int64(2)
		jid, _ := context.Get("userID")
		id := jid.(int64)
		pathlist := []string{}
		for i, file := range files {
			path := "/user/" + strconv.FormatInt(id, 10) + strconv.Itoa(i) + "pic.png"
			context.SaveUploadedFile(file, "./view"+path)
			pathlist = append(pathlist, "./view"+path)
		}
		up := model.UserPhoto{
			Uid:  id,
			Path: pathlist,
		}
		if !model.InsertUserPhoto(up) {
			context.JSON(http.StatusInternalServerError, tools.PhotoUpLoadErr)
		}
		context.JSON(http.StatusOK, tools.OK)
	}
}
func UpLoadBookPhoto() gin.HandlerFunc {
	return func(context *gin.Context) {
		//file, _ := context.FormFile("photofile")
		filess, _ := context.MultipartForm()
		files := filess.File["photofiles"]
		//a为从jwt中取出的1
		//a := int64(2)
		jid, _ := context.Get("userID")
		id := jid.(int64)
		pathlist := []string{}
		for i, file := range files {
			path := "/user/" + strconv.FormatInt(id, 10) + strconv.Itoa(i) + "pic.png"
			context.SaveUploadedFile(file, "./view"+path)
			pathlist = append(pathlist, "./view"+path)
		}
		up := model.BookPhoto{
			Bid:  id,
			Path: pathlist,
		}
		if !model.InsertBookPhoto(up) {
			context.JSON(http.StatusInternalServerError, tools.PhotoUpLoadErr)
		}
		context.JSON(http.StatusOK, tools.OK)
	}
}
