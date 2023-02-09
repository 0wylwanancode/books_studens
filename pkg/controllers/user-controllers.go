package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pkg/models"
	"pkg/utils"
)

//登录
func Logig(w http.ResponseWriter, r *http.Request) {
	//登录前端是用传递用户名跟密码
	if r.Method != "POST" {
		http.Error(w, "请求方式不对", http.StatusNotFound)
		return
	}
	//从body里面拿输出
	u := models.User{}
	utils.ParseBody(r, &u)
	fmt.Println(u)
	//调用 models中GetUserByUsername()查询用户信息
	user, _, err := models.GetUserByUsername(u.Username)
	fmt.Println(user)

	//查询失败返回
	if err != nil {
		http.Error(w, "GetUserByUsername()查询失败", http.StatusNotFound)
		return
	}
	//密码错误返回
	// 注册的时候用户密码加密了，那么用户输入的密码并非存在数据库的密码

	if !utils.VerifyPassword(u.Password, user.Password) {
		http.Error(w, "Wrong Username or Password", http.StatusNotFound)
		return
	}
	//签发token  类似创建cookie
	cookie := http.Cookie{
		Name:  user.Name,
		Value: user.Username,
	}
	w.Header().Set("accessToken", cookie.String())

	//返回请求结果
	/* data := make(map[interface{}]interface{})
	data["id"] = user.ID */
	data := model2Map(*user)
	res, _ := json.Marshal(data)
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK) //响应返回
	w.Write(res)                 //写入到w
}

//
func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("accessToken", "")
	w.WriteHeader(http.StatusOK)
}

//注册
func Register(w http.ResponseWriter, r *http.Request) {
	//登录前端是用传递用户名跟密码
	if r.Method != "POST" {
		http.Error(w, "请求方式不对", http.StatusNotFound)
		return
	}
	//从body里面拿输出
	u := models.User{}
	utils.ParseBody(r, &u)

	//将获得 用户信息创建成用户
	//判断用户是否存在
	//
	uu, _, _ := models.GetUserByUsername(u.Username)
	//当你通过从前端获取Username中去查询这个用户的信息，就比如 你输入的张三已经存在了 说明这个用户已经注册过了 无需再注册
	if uu.Username == u.Username {
		fmt.Fprintf(w, "用户%s已经注册了", u.Username)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//创建前完善用户信息,并给密码加密
	u.Password = utils.EncryptPasswords(u.Password)
	//创建完用户了
	user, _, _ := models.CreateUser(u)

	date := model2Map(*user)
	res, _ := json.Marshal(date)
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func model2Map(user models.User) map[string]interface{} {
	data := map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Eamil,
		"name":       user.Name,
		"profilePic": user.ProfilePic,
		"converPic":  user.CoverPic,
		"city":       user.City,
		"website":    user.WebSite,
	}
	return data
}
