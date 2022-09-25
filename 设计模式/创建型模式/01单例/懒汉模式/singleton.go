package 懒汉模式

import "sync"

/*
非并发安全
*/
//type singleton struct {
//
//}
//
//var ins *singleton
//
//func GetInsOr()*singleton  {
//	if ins==nil{
//		return &singleton{}
//	}
//	return ins
//}
//-----------------------------------------------//
//-----------------------------------------------//
//-----------------------------------------------//
//并发安全
//import "sync"
//type singleton struct {
//}
//var ins *singleton = &singleton{}
//var mu sync.Mutex
//
//func GetInsOr() *singleton {
//	if ins == nil {
//		mu.Lock()
//		if ins == nil {
//			ins = &singleton{}
//		}
//		mu.Unlock()
//	}
//	return ins
//}
type singleton struct {
}

var ins *singleton = &singleton{}
var once sync.Once

//使用once.Do可以确保 ins 实例全局只被创建一次，once.Do 函数还可以确保当同时有多个创建动作时，只有一个创建动作在被执行
func GetInsOr() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
