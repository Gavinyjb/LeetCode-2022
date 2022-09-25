package 饿汉模式

//import ""
type singleton struct {
}

var ins *singleton = &singleton{}

func GetInsOr() *singleton {
	return ins
}
