package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	// wallet := Wallet{}

	// // wallet.Deposit(10)
	// wallet.Deposit(Bitcoin(10))

	// got := wallet.Balance()

	// // fmt.Println("测试里 余额 的地址是", &wallet.balance)

	// // want := 10
	// want := Bitcoin(10) //生成 bitcoin

	// if got != want {
	// 	// t.Errorf("获取 %d 预期 %d", got, want)
	// 	t.Errorf("获取 %s 预期 %s", got, want)
	// }
	// =========================================================

	t.Run("定金", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("提取", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		// wallet.Withdraw(10)
		// 这里通过 errcheck . 检查出未接受检查返回的错误
		err := wallet.Withdraw(10)

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("提取 资金不足", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, Bitcoin(20))

		assertError(t, err, InsufficientFundsError)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("获得 %s 预期 %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		// t.Fatal 被调用，将停止测试
		t.Fatal("想要得到一个错误，但未得到")
	}

	// 这种设计是：测试返回的错误信息
	if got != want {
		t.Errorf("获得 '%s', 预期 '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("想要得到一个错误，但未得到")
	}
}

// 安装 linters（代码检测工具）
// go get -u github.com/kisielk/errcheck
// 使用命令 errcheck .
