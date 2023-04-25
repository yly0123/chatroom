package main

import (
	"chatProject/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
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
func writePkg(conn net.Conn) (data []byte, err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	//根据pgkLen 读取消息内容
	n, err := conn.Write(data)
	if err != nil || n != int(pgkLen) {
		fmt.Println("onn.Write err=", err)
		return
	}

	return
}
