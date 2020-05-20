package handler

import (
	"ant/meta"
	"ant/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bytes, e := ioutil.ReadFile("./static/view/index.html")
		if e != nil {
			io.WriteString(w, "server err!")
			return
		}
		io.WriteString(w, string(bytes))
	case http.MethodPost:
		//接收文件流存储到本地目录
		file, header, e := r.FormFile("file")
		if e != nil {
			fmt.Println("FormFile error : ", e)
			return
		}
		defer file.Close()
		fileMeta := meta.CreateFileMetaByHeader(header)

		newFile, e := os.Create(fileMeta.Localtion)
		if e != nil {
			fmt.Println("Create File error : ", e)
			return
		}
		defer newFile.Close()
		_, e = io.Copy(newFile, file)
		if e != nil {
			fmt.Println("Copy File error : ", e)
			return
		}

		newFile.Seek(0, 0)
		hex := util.GetFileSha256Hex(newFile)
		fileMeta.Id = hex
		meta.UpdaeFileMeta(fileMeta)
		fmt.Println(meta.GetFileMeta(hex))
		http.Redirect(w, r, "/file/upload/success", http.StatusOK)
	}

}

func UploadFileSuccess(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "success!")
}
