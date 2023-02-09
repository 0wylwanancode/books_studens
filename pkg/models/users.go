package models

import (
	"database/sql"
	"fmt"
	"pkg/config"
	"strings"
)

var db *sql.DB

type User struct {
	ID         int    `json:"id" form:"id"`                 //用户编号
	Username   string `json:"username" form:"username"`     //用户名
	Password   string `json:"password" form:"password"`     //密码
	Eamil      string `json:"email" form:"email"`           //邮箱
	Name       string `json:"name" form:"name"`             //昵称
	CoverPic   string `json:"coverPic" form:"coverPic"`     //背景图
	ProfilePic string `json:"profilePic" form:"profilePic"` //头像
	City       string `json:"city" form:"city"`             //城市
	WebSite    string `json:"webSite" form:"webSite"`       //个人网址
}

func init() {
	db = config.GetDB()
}

//获得所有用户信息
func GetAllUser() ([]User, error) {
	//1.准备sql语句
	sqlStr := "select * from users"
	//预处理sql
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare 预处理失败", err)
		return nil, err
	}
	//延迟关闭链接
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("Query 查询失败", err)
		return nil, err
	}
	//返回结果
	return rowsUser(rows), nil
}

//通过用户id查询某个用户信息
func GetUserByID(userID int) (*User, error) {
	//准备sql语句
	sqlStr := "select * from users where id=?"
	//预处理
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare 预处理失败", err)
		return nil, err
	}
	//延迟关闭stmt
	defer stmt.Close()
	//执行sql语句
	rows := stmt.QueryRow(userID)

	//填充内容
	return rowUser(rows), nil
}

func GetUserByUsername(username string) (*User, *sql.DB, error) {
	//准备sql语句
	sqlStr := "select * from users where username=?"
	//预处理
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare 预处理失败", err)
		return nil, nil, err
	}
	//延迟关闭stmt
	defer stmt.Close()
	//执行sql语句
	fmt.Println("username内容", username)
	rows := stmt.QueryRow(username)
	user := User{}
	err = rows.Scan(&user.ID, &user.Username, &user.Eamil, &user.Password, &user.Name, &user.ProfilePic, &user.CoverPic, &user.City, &user.WebSite)
	if err != nil {
		fmt.Println(err)
	}

	//填充内容
	return &user, db, nil
}

//创建用户信息
func CreateUser(user User) (*User, *sql.DB, error) {
	//准备sql语句
	sqlStr := "insert into users(username,email,password,name,coverPic,profilePic,city,website)value(?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("CreateUser 预处理失败", err)
		return nil, nil, err
	}
	defer stmt.Close() //延迟关闭链接
	result, err := stmt.Exec(user.Username, user.Eamil, user.Password, user.Name, user.CoverPic, user.ProfilePic, user.City, user.WebSite)
	if err != nil {
		fmt.Println("CreateUser 执行创建用户信息失败", err)
		return nil, nil, err
	}
	//获得执行后主键id
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return &user, db, nil
}

//修改某个用户信息
func UpdateUser(user User) (*User, *sql.DB, error) {
	tempSql := "update users set"
	//在修改时要注意一个问题 用户未填的信息代表没有修改，只有不为空才修改了
	var params []interface{}
	if user.Name != "" {
		tempSql = tempSql + ",name=?"
		params = append(params, user.Name)
	}
	if user.City != "" {
		tempSql = tempSql + ",city=?"
		params = append(params, user.City)
	}
	if user.WebSite != "" {
		tempSql = tempSql + ",webSite=?"
		params = append(params, user.WebSite)
	}
	if user.ProfilePic != "" {
		tempSql = tempSql + ",profilePic=?"
		params = append(params, user.ProfilePic)
	}
	if user.CoverPic != "" {
		tempSql = tempSql + ",coverPic=?"
		params = append(params, user.CoverPic)
	}
	sqlStr := tempSql + " where id=?"
	params = append(params, user.Username)
	//这样会导致Update <表名> set,name=? 或者,Eamil=?
	//所以我们将字符串中第一个,替换成""
	sqlStr = strings.Replace(sqlStr, ",", " ", 1)
	fmt.Println(tempSql)
	fmt.Println(params...)
	//预执行sql
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("UpdateUser中Prepare错误:", err)
		return nil, nil, err
	}
	//延迟关闭
	defer stmt.Close()
	//执行sql语句
	result, err := stmt.Exec(params...)
	if err != nil {
		fmt.Println("UpdateUser中Exec错误", err)
		return nil, nil, err
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)

	return &user, db, nil
}

//读取所有用户信息的内部函数
func rowsUser(rows *sql.Rows) []User {
	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Username, &user.Eamil, &user.Password, &user.Name, &user.CoverPic, &user.City, &user.WebSite)
		users = append(users, user)
	}
	return users
}

//读取某个用户信息的内部函数
func rowUser(rows *sql.Row) *User {
	user := User{}
	rows.Scan(&user.ID, &user.Username, &user.Eamil, &user.Password, &user.Name, &user.CoverPic, &user.City, &user.WebSite)
	return &user
}
