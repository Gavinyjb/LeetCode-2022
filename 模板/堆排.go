package 模板

func heapSort(a []int)[]int  {
	//1. 无序数组
	//2. 将无序数组a构建为一个大根堆
	for i:=len(a)>>1-1;i>=1;i--{
		sink(a,i,len(a))
	}
	//3. 交换a[0]和a[len(a)-1]
	//3. 然后那前面这段数组继续下沉保持堆结构，如此循环即可
	for i:=len(a)-1;i>=1
}
