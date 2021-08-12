package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int //这是一个类型别名

// 为 Bitcoin 类型实现 方法，String 这个方法接口在 fmt 里定义了
// 这能实现 Bitcoin 的 %s 打印字符串
func (b Bitcoin) String() string {
	// 生成字符串 Sprintf()
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin //余额，小写开头，私有
}

// 存
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Println("定金里 余额 的地址是", &w.balance)
	w.balance += amount
}

// 余额
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// 错误，是值
var InsufficientFundsError = errors.New("不能提取，资金不足")

// 提取
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		// errors.New 使用给定的错误信息构造一个基本的 error 值
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
