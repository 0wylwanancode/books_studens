package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//获取前端body里面传过来的数据
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		//将body反序列化
		err := json.Unmarshal([]byte(body), x) //[]byte()是转化格式 格式相同不会影响
		if err != nil {
			return
		}
	}
}

//加密函数
func EncryptPasswords(password string) string {
	md := md5.New()
	md.Write([]byte(password))
	return hex.EncodeToString(md.Sum(nil))
}

//验证
func VerifyPassword(password, encrypPasswrod string) bool {
	//password是明文的密码,encrypPasswrod是加密后的密码

	return EncryptPasswords(password) == encrypPasswrod
}
