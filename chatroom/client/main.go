package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {
	//接受用户选择
	var key int
	//判断是否继续显示菜单
	var loop = true

	for loop {
		fmt.Println("---------------欢迎登录多人聊天系统------------")
		fmt.Println("\t\t\t 1.登录聊天室")
		fmt.Println("\t\t\t 2.注册系统")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册系统")
			loop = false
		case 3:
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	if key == 1 {
		fmt.Println("请输入用户的id")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&userPwd)
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}

	} else if key == 2 {
		fmt.Println("进行用户注册")
	}
}
