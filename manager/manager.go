package manager

import (
	"fmt"
	"os"
	"student_manager/MyTinyLogger/mylogger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type student struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Age   int    `db:"age"`
	Score int    `db:"score"`
}

var log *mylogger.ConsoleLogger = mylogger.NewConsoleLogger("trace")

const dsn string = "root:ch981205@tcp(101.132.177.244)/test"

var db *sqlx.DB

func ShowMenu() {
	fmt.Println("**********学生信息管理系统**********")
	fmt.Println("********* 1. 展示学生列表 **********")
	fmt.Println("********* 2. 添加学生信息 **********")
	fmt.Println("********* 3. 编辑学生信息 **********")
	fmt.Println("********* 4. 删除学生信息 **********")
	fmt.Println("********* 5.   退出系统   **********")
}

func ShowStudent() {
	//query from mysql
	var stus []student
	sqlStr := `select id,name,age,score from students;`
	err := db.Select(&stus, sqlStr)
	if err != nil {
		log.Error("query failed,err:%s\n", err)
		return
	}
	fmt.Printf("学号\t学生姓名\t年龄\t分数\n")
	fmt.Println("=====================================")
	for _, v := range stus {
		fmt.Printf("%d\t%s\t\t%d\t%d\n", v.Id, v.Name, v.Age, v.Score)
	}
}

func AddStudent() {
	s := new(student)
	var num int
	fmt.Println("请输入学生序号：")
	fmt.Scan(&num)
	//查询该学生是否存在
	if QueryExist(num) {
		log.Info("该学号已存在！")
		// fmt.Printf("学号\t学生姓名\t年龄\t分数\n")
		// fmt.Println("=====================================")
		// fmt.Printf("%d\t%s\t\t%d\t%d\n", s.Id, s.Name, s.Age, s.Score)
		return
	}
	s.Id = num
	fmt.Println("请输入学生姓名：")
	fmt.Scan(&s.Name)
	fmt.Println("请输入学生年龄：")
	fmt.Scan(&s.Age)
	fmt.Println("请输入学生成绩：")
	fmt.Scan(&s.Score)
	sqlStr := `insert into students (id,name,age,score) values (?,?,?,?);`
	_, err := db.Exec(sqlStr, s.Id, s.Name, s.Age, s.Score)
	if err != nil {
		log.Error("insert failed,err:%s\n", err)
		return
	}
}

func QueryExist(n int) bool {
	sqlStr := `select * from students where id=?;`
	var s student
	err := db.Get(&s, sqlStr, n)
	return err == nil
}

func ModifyStudent() {
	var num int
	s := new(student)
	fmt.Println("请输入您要修改的学生序号：")
	fmt.Scan(&num)
	if !QueryExist(num) {
		log.Info("该学号不存在！")
		return
	}
	s.Id = num
	fmt.Println("请输入修改后的学生姓名：")
	fmt.Scan(&s.Name)
	fmt.Println("请输入修改后的学生年龄：")
	fmt.Scan(&s.Age)
	fmt.Println("请输入修改后的学生成绩：")
	fmt.Scan(&s.Score)
	//TODO 修改数据库
	sqlStr := `update students set name=?,age=?,score=? where id=?;`
	_, err := db.Exec(sqlStr, s.Name, s.Age, s.Score, s.Id)
	if err != nil {
		log.Info("修改信息失败")
		return
	}
	log.Info("修改成功！")
	fmt.Println("修改后的学生信息如下：")
	fmt.Printf("学号\t学生姓名\t年龄\t分数\n")
	fmt.Println("=====================================")
	fmt.Printf("%d\t%s\t\t%d\t%d\n", s.Id, s.Name, s.Age, s.Score)
}

func DeleteStudent() {
	var num int
	//s := new(student)
	fmt.Println("请输入您要删除的学生序号：")
	fmt.Scan(&num)
	if !QueryExist(num) {
		log.Info("该学号不存在！")
		return
	}
	sqlStr := `delete from students where id=?`
	_, err := db.Exec(sqlStr, num)
	if err != nil {
		log.Info("删除失败")
	}
	log.Info("删除成功")
}

func ExitSys() {
	fmt.Println("谢谢您的使用！")
	os.Exit(0)
}

func InitDB() (err error) {
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}
