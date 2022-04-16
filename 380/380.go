package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	hashMap map[int]int
	keyList []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hashMap: make(map[int]int, 0),
		keyList: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashMap[val]; ok {
		return false
	} else {
		this.hashMap[val] = len(this.keyList)
		this.keyList = append(this.keyList, val)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.hashMap[val]
	if !ok {
		return false
	} else {
		delete(this.hashMap, val)
		last := len(this.keyList) - 1
		this.keyList[id] = this.keyList[last]
		this.hashMap[this.keyList[id]] = id
		this.keyList = this.keyList[:last]
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(this.keyList))
	return this.keyList[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
func main() {
	val, val2, val3, val4 := 1, 2, 3, 4
	obj := Constructor()
	param := obj.Insert(val)
	fmt.Println(obj, param)
	obj.Insert(val2)
	param2 := obj.GetRandom()
	fmt.Println(obj, param2)
	param = obj.Remove(val)
	fmt.Println(obj, param)
	param = obj.Remove(val3)
	fmt.Println(obj, param)
	param = obj.Insert(val4)
	fmt.Println(obj, param)

}
