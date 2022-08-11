package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	keyList []int
	hashmap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.hashmap[val]; ok {
		return false
	}
	rs.hashmap[val] = len(rs.keyList)
	rs.keyList = append(rs.keyList, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	id, ok := rs.hashmap[val]
	if !ok {
		return false
	}
	last := len(rs.keyList) - 1
	rs.keyList[id] = rs.keyList[last]
	rs.hashmap[rs.keyList[id]] = id
	rs.keyList = rs.keyList[:last]
	delete(rs.hashmap, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.keyList[rand.Intn(len(rs.keyList))]
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
