package 排序

/**
 * @Author: yirufeng
 * @Date: 2021/4/13 8:28 下午
 * @Desc:

有好几个版本：
1. 最基本的
2. 如果某次冒泡的时候没有交换元素，那么说明这一趟以及后面的元素都是有序的，之后后面我们就不再需要进行冒泡，可以直接退出
3. 在第二种改进的基础上再使用一个变量，表示我们最后交换的位置，下次遍历到这个位置就可以了
 **/

func BubbleSort(nums []int) {
	bubbleSortV2(nums)
}

//最基本的
func bubbleSortV1(nums []int) {
	for i := 1; i < len(nums); i++ { //冒泡的次数
		for j := 0; j < len(nums)-i; j++ {
			//如果前面比后面大就交换
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

//如果某次冒泡的时候没有交换元素，那么说明这一趟以及后面的元素都是有序的，之后后面我们就不再需要进行冒泡，可以直接退出
func bubbleSortV2(nums []int) {
	//表示某一趟是否有元素交换
	flag := true
	for i := 1; i < len(nums); i++ { //冒泡的次数
		flag = true
		for j := 0; j < len(nums)-i; j++ {

			//如果前面比后面大就交换
			if nums[j] > nums[j+1] {
				flag, nums[j], nums[j+1] = false, nums[j+1], nums[j]
			}
		}
		//如果没有元素交换，说明已经有序，直接退出即可
		if flag {
			break
		}
	}
}

//在第二种改进的基础上再使用一个变量，表示我们最后交换的位置，下次遍历到这个位置就可以了
//使用一个变量记录最后发生交换的位置，说明该位置之后都是有序的
func bubbleSortV3(nums []int) {
	//表示某一趟是否有元素交换
	flag := true
	//表示最后发生交换的位置
	//注意点1:这里必须再额外使用临时变量，因为如果只是用一个变量lastSwapPos的话，
	//lastSwapPos的值会一直改变，因此如果一旦变小，我们的for循环就会推出
	lastSwapPos := len(nums) - 1     //注意点2：初始值赋值是len(nums) - 1
	for i := 1; i < len(nums); i++ { //冒泡的次数
		lastSwapPosTemp := len(nums) - i //注意点3：初始值赋值是len(nums) - i
		flag = true
		//下次遍历的时候遍历到我们上次最后发生交换的位置就可以了
		for j := 0; j < lastSwapPos; j++ {
			//如果前面比后面大就交换
			if nums[j] > nums[j+1] {
				flag, nums[j], nums[j+1], lastSwapPosTemp = false, nums[j+1], nums[j], j //注意点4：这里lastSwapPosTemp应该赋值j，因为j与j+1比较过了，所以就认为j之后是已经有序的了
			}
		}
		//如果没有元素交换，说明已经有序，直接退出即可
		if flag {
			break
		}

		lastSwapPos = lastSwapPosTemp
	}
}
