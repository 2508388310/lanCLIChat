package server
import (
	"fmt"
	"net"
	"sync"

	"lanCLIChat/internal/user"
)

type Server struct {
	Ip   string
	Port int
	OnlineUserMap map[string]string
	mapLock sync.RWMutex
}


// 请输入ip和端口创建服务器，返回一个服务器对象
func NewServer(ip string, port int) *Server {
	
	server := &Server{
		Ip:   ip,
		Port: port,
		OnlineUserMap: make(map[string]string),
	}
	return server
}

func (server *Server) handle(conn net.Conn){
	fmt.Println("连接建立成功")
	newUserName := user.NewUser(conn)
	server.mapLock.Lock()
	server.OnlineUserMap[newUserName] = newUserName
	server.mapLock.Unlock()
}
func (server *Server)userPrint(){
	for username, userInfo := range server.OnlineUserMap {
    fmt.Printf("Username: %s, UserInfo: %+v\n", username, userInfo)
}
}
//Start 启动服务器
func(server *Server) Start(){
	listener,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",server.Ip,server.Port))
	if err!=nil{
		fmt.Println("net.Listen err:",err)
		return
	}
	defer listener.Close()
	defer server.userPrint()

	jishu:=0
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("listener.Accept err:",err)
			continue
		}
		go server.handle(conn)
		jishu++
		if jishu>2{
			break 
		}
	}
	

}
