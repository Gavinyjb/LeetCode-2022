package leetcodetopinterviewquestions

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size:=len(nums1)+len(nums2)
	var even bool 
	if size&1==0{  //size为偶数
		even=true  
	}else {
		even =false  //size为奇数
	}

	if (len(nums1)>0) && (len(nums2)>0){
		if even {  //size为偶数
			return float64(findKthNum(nums1,nums2,size/2)+findKthNum(nums1,nums2,size/2+1))/2
		}else {   //size为奇数
			return float64(findKthNum(nums1,nums2,size/2+1))
		}
	}else if(len(nums1)>0){
		if even{
			return float64(nums1[(size-1)/2]+nums1[size/2])/2
		}else {
			return float64(nums1[size/2])
		}
	}else if (len(nums2)>0){
		if even{
			return float64(nums2[(size-1)/2]+nums2[size/2])/2
		}else {
			return float64(nums2[size/2])
		}
	}else{
		return 0
	}
}
//获取有序数组arr1 arr2合并后的第K个数字 
func findKthNum(arr1 ,arr2 []int,kTh int)int  {
	var longs,shorts []int

	if len(arr1)>len(arr2){
		longs,shorts=arr1,arr2
	}else {
		shorts,longs=arr1,arr2
	}

	l:=len(longs)
	s:=len(shorts)
	if kTh<=s{
		return getUpMedian(shorts,0,kTh-1,longs,0,kTh-1)
	}
	if kTh>l{
		if shorts[kTh-l-1]>=longs[l-1]{
			return shorts[kTh-l-1]
		}
		if longs[kTh-s-1]>=shorts[s-1]{
			return longs[kTh-s-1]
		}
		return getUpMedian(shorts,kTh-l,s-1,longs,kTh-s,l-1)
	}
	if longs[kTh-s-1]>=shorts[s-1]{
		return longs[kTh-s-1]
	}
	return getUpMedian(shorts,0,s-1,longs,kTh-s,kTh-1)
}
//a1 a2 为有序数组 
//!!!!!!! e1-s1+1==e2-s2+1 即长度相同 返回上中位数
// 1 2 3 4
// 1'2'3'4' 返回第4小的数
// 1 2 3 4 5
// 1'2'3'4'5' 返回第5小的数
func getUpMedian(a1 []int,s1,e1 int,a2 []int,s2,e2 int) int  {
	mid1,mid2:=0,0
	offset:=0
	for s1<e1{
		mid1=(s1+e1)/2
		mid2=(s2+e2)/2
		offset=((e1-s1+1)&1)^1  //长度为奇数：offset=0  长度为偶数：offset=1
		if a1[mid1]>a2[mid2] {
			// 1 2 3 4
			// 1'2'3'4' 
			// 2>2'则取
			// 1 2    
			//     3'4' 
			e1=mid1
			s2=mid2+offset
		}else if a1[mid1]<a2[mid2] {
			// 1 2 3 4
			// 1'2'3'4' 
			// 2<2'则取
			//     3 4
			// 1'2'
			 s1=mid1+offset
			 e2=mid2
		}else {
			return a1[mid1]
		}
	}
	//两个数组长度为1
	return min(a1[s1],a2[s2])
} 
func min(a,b int)int  {
	if a<b{
		return a
	}
	return b
}