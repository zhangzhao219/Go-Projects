package main

import "fmt"

func quicksort(arr []int) {
	// 1. 获取数组的标志位，判断是否符合返回条件
	L := 0            // 数组最左侧标志位
	R := len(arr) - 1 // 数组最右侧标志位
	// 如果这个数组只剩下一个数字（L=R）或者为空(L>R)，说明不需要排序，直接退出即可
	if L >= R {
		return
	}

	// 2. 设置移动的双指针和中轴
	start, end := L, R // 移动的双指针
	key := arr[start]  // 选取划分中轴

	// 3. 正式算法
	// 如果start == end，说明区间范围已经是[start, start]，说明当前只剩下start一个数字位没有判断过
	// 这个数字位应该刚好是划分的中轴，因此直接赋值退出即可
	for start < end {
		// 从右边开始找，找到第一个小于中轴的数字，直接将这个数字覆盖到start的位置上
		for start < end && arr[end] >= key {
			end--
		}
		arr[start] = arr[end]
		// 然后从左边开始找，找到第一个大于中轴的数字，直接将这个数字覆盖到end的位置上
		for start < end && arr[start] <= key {
			start++
		}
		arr[end] = arr[start]
	}
	arr[start] = key

	// 4. 递归对左右两侧进行排序
	quicksort(arr[:start])
	quicksort(arr[start+1:])
}

func mergesort(arr []int) {
	// 1. 判断返回条件
	// 如果就剩下一个数字或者不剩下数字了，说明不需要再次拆分，可以返回了
	if len(arr) <= 1 {
		return
	}
	// 2. 递归拆分
	mid := len(arr) / 2
	mergesort(arr[:mid])
	mergesort(arr[mid:])
	// 3. 合并，现在有两个有序数组（或者只有一个元素或者为空），需要将这两个数组合并起来
	// 创建一个新的空间，双指针进行判断即可
	temparr := make([]int, len(arr))
	tindex, lindex, rindex := 0, 0, mid
	for tindex < len(arr) {
		if lindex < mid && rindex < len(arr) {
			if arr[lindex] < arr[rindex] {
				temparr[tindex] = arr[lindex]
				lindex++
			} else {
				temparr[tindex] = arr[rindex]
				rindex++
			}
		} else if lindex < mid {
			copy(temparr[tindex:], arr[lindex:])
			break
		} else if rindex < len(arr) {
			copy(temparr[tindex:], arr[rindex:])
			break
		}
		tindex++
	}
	copy(arr, temparr)
}

// 堆相关操作

// 堆的数据结构
// 索引从0开始，则i的左孩子是2i+1，右孩子是2i+2，i的父亲是(i-1)/2
// 用切片表示堆，size控制堆的大小，排序过程中有用
type Heap struct {
	heap []int
	size int
}

// 返回带有数字的堆的实例
func NewHeap(arr []int) *Heap {
	return &Heap{
		heap: arr,
		size: len(arr),
	}
}

// 创建大顶堆
func (h *Heap) CreateHeap() {
	// 叶子节点一定满足堆的性质，因此从后往前数第一个非叶子节点开始向下调整即可
	lastfather := (h.size - 1 - 1) / 2
	for i := lastfather; i >= 0; i-- {
		h.AdjustDown(i)
	}
}

// 堆排序
func (h *Heap) Sort() {
	// 不断删除元素（不是真的删除，仅仅是移动元素到切片的末尾）
	for i := h.size - 1; i > 0; i-- {
		// 1. 交换元素
		t := h.heap[0]
		h.heap[0] = h.heap[i]
		h.heap[i] = t
		// 2. 减小数组大小
		h.size -= 1
		// 3. 堆顶元素向下调整
		h.AdjustDown(0)
	}
	fmt.Println(h.heap)
}

// 删除堆顶元素
func (h *Heap) Delete() int {
	// 1. 保留堆顶的元素
	num := h.heap[0]
	// 2. 替换堆顶元素为最后面的元素
	h.heap[0] = h.heap[h.size-1]
	// 3. 减小堆的尺寸
	h.size -= 1
	h.heap = h.heap[:h.size]
	// 4. 堆顶元素向下调整
	h.AdjustDown(0)
	return num
}

// 插入元素
func (h *Heap) Insert(num int) {
	// 1. 更新堆的尺寸
	h.heap = append(h.heap, num)
	h.size += 1
	// 2. 新结点向上调整
	h.AdjustUp(h.size - 1)
}

// 向下调整
func (h *Heap) AdjustDown(index int) {

	k := h.heap[index] // 记录这个结点
	lchild := 2*index + 1
	rchild := 2*index + 2
	for lchild < h.size {
		// 1. 找到左右孩子中最大的那个
		var mchild int
		if rchild < h.size && h.heap[rchild] > h.heap[lchild] {
			mchild = rchild
		} else {
			mchild = lchild
		}
		// 2. 如果还不比最大的孩子大，直接返回
		if h.heap[mchild] < k {
			break
		}
		// 3. 否则赋值，判断下一组
		h.heap[index] = h.heap[mchild]
		index = mchild
		lchild = 2*mchild + 1
		rchild = 2*mchild + 2
	}
	h.heap[index] = k
}

// 向上调整
func (h *Heap) AdjustUp(index int) {
	k := h.heap[index] // 记录这个结点
	father := (index - 1) / 2
	// 注意-1/2也为0，因此需要判断是不是都为0
	for father >= 0 && index != father {
		if h.heap[father] >= k {
			break
		}
		h.heap[index] = h.heap[father]
		index = father
		father = (index - 1) / 2
	}
	h.heap[index] = k
}

// 打印堆
func (h *Heap) PrintHeap() {
	for i := 0; i < h.size; i++ {
		fmt.Printf("%d ", h.heap[i])
	}
	fmt.Println()
}

func main() {
	demolist := []int{1, 2, 9, 6, 7, 8, 4, 2, 1, 4, 5, 4, 1, 5, 2, 3, 6, 0, 3, 2, 5, 8, 9, 6, 3, 7, 8, 9, 6, 3, 0}
	quicksort(demolist)
	fmt.Println(demolist)
	mergesort(demolist)
	fmt.Println(demolist)
	h := NewHeap(demolist)
	h.CreateHeap()
	h.PrintHeap()
	fmt.Println(h.Delete())
	h.PrintHeap()
	h.Insert(10)
	h.PrintHeap()
	h.Sort()
}
