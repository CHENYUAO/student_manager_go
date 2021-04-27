package main

import (
	"fmt"
	"student_manager/manager"
)

func main() {
	err := manager.InitDB()
	if err != nil {
		fmt.Printf("connect to mysql failed,err:%s\n", err)
		return
	}
	manager.ShowMenu()
	//fmt.Println("connect to mysql success!")
	for {
		var choice int
		fmt.Println("请输入你的选择：")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			manager.ShowStudent()
		case 2:
			manager.AddStudent()
		case 3:
			manager.ModifyStudent()
		case 4:
			manager.DeleteStudent()
		case 5:
			manager.ExitSys()
		default:
			fmt.Println("无效的输入！")
		}
	}
}
