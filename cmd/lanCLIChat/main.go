package main

import (
	"flag"
	"fmt"
	"lanCLIChat/internal/server"
)

// 版本信息变量，将在构建时通过 ldflags 注入
var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

func main() {
	// 定义命令行参数
	var showVersion = flag.Bool("version", false, "显示版本信息")
	var showHelp = flag.Bool("help", false, "显示帮助信息")
	flag.Parse()

	// 处理版本信息
	if *showVersion {
		fmt.Printf("lanCLIChat 版本: %s\n", Version)
		fmt.Printf("构建时间: %s\n", BuildTime)
		fmt.Printf("Git提交: %s\n", GitCommit)
		return
	}

	// 处理帮助信息
	if *showHelp {
		fmt.Println("lanCLIChat - 局域网聊天工具")
		fmt.Printf("版本: %s\n\n", Version)
		fmt.Println("使用方法:")
		fmt.Println("  lanCLIChat [选项]")
		fmt.Println("\n选项:")
		fmt.Println("  -version    显示版本信息")
		fmt.Println("  -help       显示此帮助信息")
		return
	}

	// 显示启动信息
	fmt.Printf("lanCLIChat v%s 正在启动...\n", Version)
	
	// 启动服务器
	ser := server.NewServer("127.0.0.1", 8888)
	ser.Start()
	
	// ASCII 艺术字
	println("88                             ,ad8888ba,   88           88    ,ad8888ba,   88                                \n" +
		"88                            d8\"'    `\"8b  88           88   d8\"'    `\"8b  88                         ,d     \n" +
		"88                           d8'            88           88  d8'            88                         88     \n" +
		"88  ,adPPYYba,  8b,dPPYba,   88             88           88  88             88,dPPYba,   ,adPPYYba,  MM88MMM  \n" +
		"88  \"\"     `Y8  88P'   `\"8a  88             88           88  88             88P'    \"8a  \"\"     `Y8    88     \n" +
		"88  ,adPPPPP88  88       88  Y8,            88           88  Y8,            88       88  ,adPPPPP88    88     \n" +
		"88  88,    ,88  88       88   Y8a.    .a8P  88           88   Y8a.    .a8P  88       88  88,    ,88    88,    \n" +
		"88  `\"8bbdP\"Y8  88       88    `\"Y8888Y\"'   88888888888  88    `\"Y8888Y\"'   88       88  `\"8bbdP\"Y8    \"Y888  ")
}
