package handler

import (
	"ant/meta"
	"ant/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// 上传文件
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
		hex := util.GetFileSha256Hash(newFile)
		fileMeta.Id = hex
		meta.UpdaeFileMeta(fileMeta)
		fmt.Println(meta.GetFileMeta(hex))
		http.Redirect(w, r, "/file/upload/success", http.StatusOK)
	}

}

// 上传文件成功
func UploadFileSuccess(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "success!")
}

// 获取文件的元数据
func GetFileMetaDataHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form["filehash"][0]
	fileMeta := meta.GetFileMeta(fileHash)

	bytes, e := json.Marshal(fileMeta)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

// 获取所有文件的元数据
func GetAllFileMetaDataHandler(w http.ResponseWriter, r *http.Request) {
	fileMeta := meta.GetAllFileMeta()
	bytes, e := json.Marshal(fileMeta)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

// 下载文件，
func DowloadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form["filehash"][0]
	fileMeta := meta.GetFileMeta(fileHash)
	file, e := os.Open(fileMeta.Localtion)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/octect-stream")
	//w.Header().Set("Content-Description", "attachment;filename="+fileMeta.Name)
	w.Header().Set("Content-Type", "application/octect-stream")
	// attachment表示文件将会提示下载到本地，而不是直接在浏览器中打开
	w.Header().Set("content-disposition", "attachment; filename=\""+fileMeta.Name+"\"")
	io.Copy(w, file)
}

// 更新数据
func UpdateFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	op := r.Form.Get("op")
	fileHash := r.Form.Get("sha256")
	newFileName := r.Form.Get("filename")
	if op != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fileMeta := meta.GetFileMeta(fileHash)
	fileMeta.Name = newFileName
	meta.UpdaeFileMeta(fileMeta)
	bytes, e := json.Marshal(fileMeta)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// 更新数据
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form.Get("sha256")
	fileMeta := meta.DeleteFileMeta(fileHash)
	os.Remove(fileMeta.Localtion)
	w.WriteHeader(http.StatusOK)
}
