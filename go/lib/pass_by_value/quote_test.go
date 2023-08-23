package pass_by_value

import "testing"

// 讨论 golang 中 两个概念
// 取引用  &&  解引用
// 使用 c 语言的时候，一般我们都值类型做值处理，指针做指针处理

type user struct {
	email string
}

func (u user) changeEmailByValue(email string) {
	u.email = email
}

func (u *user) changeEmailByPointer(email string) {
	u.email = email
}

func Test_Main(t *testing.T) {
	alice := user{"alice@163.com"}
	t.Log(alice.email)
	alice.changeEmailByValue("alice@gmail.com")
	t.Log("changeEmailByValue:", alice.email)

	bob := &user{"bob@163.com"}
	t.Log(bob.email)
	bob.changeEmailByValue("bob@gmai.com") // 用指针类型的实参调用形参为值类型的方法（会进行“自动解引用”）
	t.Log("changeEmailByValue:", bob.email)

	carol := user{"carol@163.com"}
	t.Log(carol.email)
	carol.changeEmailByPointer("carol@gmail.com") // 用值类型的实参调用形参为指针类型的方法（会进行“自动取引用”）
	t.Log("changeEmailByPointer:", carol.email)

	dave := &user{"dave@163.com"}
	t.Log(dave.email)
	dave.changeEmailByPointer("dave@gmail.com")
	t.Log("changeEmailByPointer:", dave.email)

}
