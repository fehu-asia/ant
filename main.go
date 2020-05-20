package main

import (
	"ant/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadFile)
	http.HandleFunc("/file/upload/success", handler.UploadFileSuccess)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe error : ", err)
		return
	}
}
