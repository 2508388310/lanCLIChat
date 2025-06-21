package main

import (
	"lanCLIChat/internal/server"
)

func main() {
	srv := server.NewServer("127.0.0.1", 8888)
	println(srv.Ip)
	println(srv.Port)
	println("88                             ,ad8888ba,   88           88    ,ad8888ba,   88                                \n" +
		"88                            d8\"'    `\"8b  88           88   d8\"'    `\"8b  88                         ,d     \n" +
		"88                           d8'            88           88  d8'            88                         88     \n" +
		"88  ,adPPYYba,  8b,dPPYba,   88             88           88  88             88,dPPYba,   ,adPPYYba,  MM88MMM  \n" +
		"88  \"\"     `Y8  88P'   `\"8a  88             88           88  88             88P'    \"8a  \"\"     `Y8    88     \n" +
		"88  ,adPPPPP88  88       88  Y8,            88           88  Y8,            88       88  ,adPPPPP88    88     \n" +
		"88  88,    ,88  88       88   Y8a.    .a8P  88           88   Y8a.    .a8P  88       88  88,    ,88    88,    \n" +
		"88  `\"8bbdP\"Y8  88       88    `\"Y8888Y\"'   88888888888  88    `\"Y8888Y\"'   88       88  `\"8bbdP\"Y8    \"Y888  ")
}
