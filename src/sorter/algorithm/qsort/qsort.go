package qsort

//快速排序
func QuickSort(values []int){
	quickSort(values,0,len(values) -1)
}

func quickSort(values []int, left int, right int) {

	if left >= right {
		return
	}

	p := partition(values,left,right)

	//递归排序
	quickSort(values, left, p-1)   //将左半部分排序
	quickSort(values,p+1,right)   //将右半部分排序
}

//切分
func partition(values []int, left int, right int) int {
	/*
		关键部分，需要满足三个条件：
			1.对于某个p，values[p]已经排定
			2.values[left]到values[p-1]中所有元素都不大于value[p]
			3.values[p+1]到values[right]中所有元素都不小于values[p]
	 */

	 //将数组切分为 values[left ..p-1],values[p],values[p+1...right]
	 i,j := left,right+1    //左右扫描指针
	 value :=values[left]   //切分元素

	 for{
	 	//最左边扫描，找到第一个大于value的位置
	 	for {
			i += 1
			if values[i] > value {
				break
			}
			if i == right {
				break
			}
		}

		//最右边扫描，找到第一个小于value的位置
		for{
			j -= 1
			if values[j] < value {
				break
			}
			if j == left {
				break
			}
		}

		if i >= j {
			break
		}

		//交换i和j位置的元素
		values[i],values[j] = values[j],values[i]
	 }

	 //确定p的正确位置
	 values[left],values[j] = values[j],values[left]

	 return j
}