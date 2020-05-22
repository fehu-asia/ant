package main

import (
	"ant/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadFile)
	http.HandleFunc("/file/upload/success", handler.UploadFileSuccess)
	http.HandleFunc("/file/meta", handler.GetFileMetaDataHandler)
	http.HandleFunc("/file/allMeta", handler.GetAllFileMetaDataHandler)
	http.HandleFunc("/file/download", handler.DowloadFile)
	http.HandleFunc("/file/update", handler.UpdateFile)
	http.HandleFunc("/file/delete", handler.DeleteFile)
	// 用户登录
	http.HandleFunc("/user/signup", handler.UserRegister)
	http.HandleFunc("/user/signup", handler.UserRegister)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe error : ", err)
		return
	}
}
