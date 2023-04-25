package main

import (
	"chatProject/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func login(userId int, userPwd string) (err error) {

	//下一个就要开始定义协议

	//fmt.Printf("userId = %d userPwd = %s\n", userId, userPwd)
	//return nil
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//准备通过coon发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将data赋给 mes.Data
	mes.Data = string(data)
	//将mes进行序列号
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write err=", err)
		return
	}
	//fmt.Println("客户端发送消息的长度=%d 内容=%s", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	//休眠
	time.Sleep(20 * time.Second)
	fmt.Println("休眠了")
	return
}
