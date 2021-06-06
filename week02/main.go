package main

import (
	"fmt"
	"geek/model"
	"log"
)

//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
//是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

// 答:
// 我认为应该根据业务的不同需要来决定是否应该wrap该error，如果该数据库操作只有单一的应用场景
// 则可以直接按业务需求直接wrap该错误。
// 如果该数据操作wrap sql.ErrNoRows错误后，会给不同业务场景调用处理错误时带来不好的操作，我认为可以直接
// 交给上一层进行处理。
// 我平时工作习惯是丢给上一层进行处理，而不是在dao层中直接wrap该error.


func main() {
	user,err :=model.GetUserByName("12345678910")
	if err != nil{
		log.Println(err)
		return
	}
	fmt.Println(user)
}


