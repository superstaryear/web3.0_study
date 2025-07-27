package main

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
type Student struct {
	Id    uint
	Name  string
	Age   uint
	Grade string
}

func main() {
	//1、向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	student := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	DB.Create(&student)
	//2、查询 students 表中所有年龄大于 18 岁的学生信息
	var studentQuery Student
	DB.Where("age > ?", 18).Find(&studentQuery)

	//3、students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	DB.Model(&Student{}).Where("name = ?", "张三").Updates(Student{
		Grade: "四年级",
	})

	//4、删除 students 表中年龄小于 15 岁的学生记录。
	DB.Model(&Student{}).Where("age < ?", 15).Delete(&Student{})
}
