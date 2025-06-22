package user

import "net"

func NewUser(conn net.Conn) string{
	userName:=conn.RemoteAddr().String()
	return userName
}
