// 公共复用的函数的包 models
package models

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 时间戳转换日期时间
func UniToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

//日期转换成时间戳
func DateToTime(date string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, date, time.Local)
	if err != nil {
		fmt.Println(err.Error())
	}
	return t.Unix()
}

// 获取当前时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取当前的时间
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取当前日期
func GetDay() string {
	template := "2006-01-02"
	return time.Now().Format(template)
}

// 判断手机号码是否存在
func CheckMobile(phone string) bool {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	regRuler := "^1[345789]{1}\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)

}

// 判断邮箱格式是否正确
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func RandomString(n int) string {
	if n < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string = ""
	for i := 1; i <= n; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar

}

func FromPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	return string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
}
