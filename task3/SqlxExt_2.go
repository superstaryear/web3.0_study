package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Employee struct {
	ID         uint
	Name       string
	Department string
	Salary     decimal.Decimal `gorm:"type:decimal(10,2)"`
}

func initEmployeeData(db *gorm.DB) {
	emplyees := []Employee{
		{Name: "张三", Department: "市场部", Salary: decimal.NewFromFloat(1001.1)},
		{Name: "李四", Department: "技术部", Salary: decimal.NewFromFloat(2000)},
		{Name: "王五", Department: "产品部", Salary: decimal.NewFromFloat(800)},
		{Name: "赵六", Department: "技术部", Salary: decimal.NewFromFloat(3000)},
	}
	db.Create(emplyees)
}

/*
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。,
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中
*/
func main() {
	//DB.AutoMigrate(&Employee{})
	//initEmployeeData(DB)
	var employees []Employee
	//表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。,
	//DB.Model(&Employee{}).Where("department = ?", "技术部").Scan(&employees)
	err := DB.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		panic(err)
	}
	fmt.Println(employees)
	fmt.Println("-----------分割线----------------")
	//使用Sqlx查询 employees 表中工资最高的员工信息
	var employee Employee
	//DB.Model(&Employee{}).Order("salary desc").Limit(1).Find(&employee)
	DB.Get(&employee, "select * from employees order by salary desc limit 1")
	fmt.Println(employee)

}
