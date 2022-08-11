package main

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		num := 0
		for _, v := range account {
			num += v
		}
		if num > max {
			max = num
		}
	}
	return max
}
func main() {

}
