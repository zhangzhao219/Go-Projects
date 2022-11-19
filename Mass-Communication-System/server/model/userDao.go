package model

import (
	"Go-Projects/Mass-Communication-System/common/message"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 使用工厂模式创建一个UserDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 在服务器启动后初始化一个userDao实例
var (
	MyUserDao *UserDao
)

// 定义一个userDao的结构体
type UserDao struct {
	pool *redis.Pool
}

// 根据用户id返回user实例
func (ud *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	// 把res反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 完成登录的校验
func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := ud.pool.Get()
	defer conn.Close()
	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 注册
func (ud *UserDao) Register(user *message.User) (err error) {
	conn := ud.pool.Get()
	defer conn.Close()
	_, err = ud.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	// 说明该用户还没有注册过，则可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSET", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误，err=", err)
		return
	}
	return
}
