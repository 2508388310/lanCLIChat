package server
import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}


// 请输入ip和端口创建服务器，返回一个服务器对象
func NewServer(ip string, port int) *Server {
	
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (server *Server) handle(conn net.Conn){
	fmt.Println("连接建立成功")
}
//Start 启动服务器
func(server *Server) Start(){
	listener,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",server.Ip,server.Port))
	if err!=nil{
		fmt.Println("net.Listen err:",err)
		return
	}
	defer listener.Close()
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("listener.Accept err:",err)
			continue
		}
		go server.handle(conn)
	}
	

}
