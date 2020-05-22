package handler

import (
	"ant/db"
	"ant/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	passwordSalt = "#890"
)

// 用户注册
func UserRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bytes, e := ioutil.ReadFile("./static/view/signup.html")
		if e != nil {
			io.WriteString(w, "server err!")
			return
		}
		io.WriteString(w, string(bytes))
	case http.MethodPost:
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if len(username) < 3 || len(password) < 5 {
			w.Write([]byte("invalid parameter"))
			return
		}
		hash := util.GetStringSha256Hash(password + passwordSalt)
		if db.InserUser(username, hash) {
			w.Write([]byte("SUCCESS"))
		} else {
			w.Write([]byte("FAILED"))
		}

	}
}

// 用户登录
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	// 校验用户名和密码
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("invalid parameter"))
		return
	}
	hash := util.GetStringSha256Hash(password + passwordSalt)

	if !db.SelectUserByUserNameAndPassword(username, hash) {

	}
	// 生成访问凭证 token
	token := GenToken(username)
	// 登录成功后重定向到首页
	db.UpdateToken(username, token)
	w.Write([]byte("http://" + r.Host + "/static/view/home.html"))

}
func GenToken(username string) string {
	// md5 (username + timestamp + token_salt)
	ts := fmt.Sprintf("%x", time.Now().Unix())
	return util.GetStringSha256Hash(username + ts + passwordSalt)
}
