package main

import (
	"fmt"
	"log"
	"net/http"
	"webLab/action"
)

func main() {
	fmt.Println("hello world")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 启动静态文件服务

	http.HandleFunc("/input", action.InputHandler)

	http.HandleFunc("/uploadIndex", action.UploadIndex) //http://localhost:8000/uploadIndex
	http.HandleFunc("/upload", action.UploadHandle)

	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
