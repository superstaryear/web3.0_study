package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

/*
题目2：事务语句,
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

type Account struct {
	ID      uint
	Balance decimal.Decimal `gorm:"type:decimal(10,2)"`
}

type Transaction struct {
	ID            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        decimal.Decimal `gorm:"type:decimal(10,2)"`
}

/*
初始化数据
*/
func initData(db *gorm.DB) {
	var account []Account = []Account{{
		ID:      1,
		Balance: decimal.NewFromFloat(100.32),
	}, {
		ID:      2,
		Balance: decimal.NewFromInt(0),
	}}
	db.Create(&account)
}

/*
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func transfer(db *gorm.DB, fromAccountId uint, toAccountId uint, money decimal.Decimal) error {
	error := db.Transaction(func(tx *gorm.DB) error {
		var fromAccount Account
		if err := db.Take(&fromAccount, fromAccountId).Error; err != nil {
			return fmt.Errorf("账户%d查询失败,%v", fromAccountId, err.Error())
		}
		var toAccount Account
		if err := db.Take(&toAccount, toAccountId).Error; err != nil {
			return fmt.Errorf("账户%d查询失败,%v", toAccountId, err.Error())
		}
		if fromAccount.Balance.Compare(decimal.NewFromInt(100)) < 0 {
			return fmt.Errorf("账户A余额不足")
		}
		//账户 A 扣除 100 元
		if err := db.Model(&fromAccount).Update("balance", fromAccount.Balance.Sub(money)).Error; err != nil {
			return fmt.Errorf("账户%d扣除金额失败,%v", toAccountId, err.Error())
		}
		//向账户 B 增加 100 元
		if err := db.Model(&toAccount).Update("balance", toAccount.Balance.Add(money)).Error; err != nil {
			return fmt.Errorf("账户%d增加金额失败,%v", toAccountId, err.Error())
		}
		//并在 transactions 表中记录该笔转账信息
		var trasaction Transaction = Transaction{
			ToAccountId:   toAccountId,
			FromAccountId: fromAccountId,
			Amount:        money,
		}
		db.Create(&trasaction)
		return nil
	})
	return error
}
func main() {
	//DB.AutoMigrate(&Account{})
	//DB.AutoMigrate(&Transaction{})
	initData(DB)
	result := transfer(DB, 1, 2, decimal.NewFromInt(100))
	fmt.Println(result)
}
