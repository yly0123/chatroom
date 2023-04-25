package main

import (
	"chatProject/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("conn.Read err=", err)
		return
	}

	//buf[:4]转成uint32
	pgkLen := binary.BigEndian.Uint32(buf[:4])
	//根据pgkLen 读取消息内容
	n, err := conn.Read(buf[:pgkLen])
	if err != nil || n != int(pgkLen) {
		//fmt.Println("conn.Read err=", err)
		return
	}
	//把pkgLen反序列化
	err = json.Unmarshal(buf[:pgkLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//编写一个serverPeocesslogin处理登录请求

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200

	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户或密码错误"
	}
	//将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//发送data,将其封装到writePkg函数
	err = writePkg(conn, data)
	return

}

// 编写一个serverProcessMes 函数
// 功能:根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		err = serverProcessLogin(conn, mes)

	case message.LoginResMesType:
		fmt.Println("注册")
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return

}

// 处理和客户端的通讯
func process(conn net.Conn) {
	//延时关闭
	defer conn.Close()

	//循环的客户端发送消息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err != io.EOF {
				fmt.Println("客户端退出，服务器也退出")
				return
			} else {
				fmt.Println("readPkg err=", err)
				return
			}

		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}

}
func main() {

	//提示信息
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", ":8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}
	//监听成功，等待客户端
	for {
		fmt.Println("等待客户端链接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err", err)
		}
		//链接成功
		go process(conn)

	}
}
