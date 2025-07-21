package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeId int
}

func (e Employee) PrintInfo() {
	fmt.Printf("员工姓名为%s,年龄为%d,工号为%d", e.Name, e.Age, e.EmployeeId)
}

/*
第二题
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/
func main() {
	var e = Employee{
		Person:     Person{Name: "张三", Age: 20},
		EmployeeId: 1,
	}
	e.PrintInfo()
}
