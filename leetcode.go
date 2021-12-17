package goproject

import (
	"fmt"
	"testing"
)
func maxSlidingWindow(nums []int, k int) []int {
	res:=make([]int,0,1)
	h:=NewHeap()
	for i:=0;i<k;i++{
		h.Push(nums[i])
	}
	for i:=0;i+k-1<len(nums);i++{
		if i==0{
			res=append(res,h.Peek())
			continue
		}
		if h.Peek()==nums[i-1]{
			fmt.Printf("zhaoxiangyu %+v %+v \n",h.Peek(),i-1)
			h.Pop()
		}
		h.Push(nums[i+k-1])
		res=append(res,h.Peek())
	}
	return res
}

type Heap struct {
	c []int
}

func NewHeap() Heap {
	h:=Heap{c:make([]int,0,1)}
	return h
}

func (h *Heap) Push(x int) {
	//fmt.Println("enter")
	h.c = append(h.c,x)
	idx:=len(h.c)-1
	for idx -1>=0 && h.c[(idx-1)/2]< h.c[idx]{
		h.c[(idx-1)/2],h.c[idx]=h.c[idx],h.c[(idx-1)/2]
		idx = (idx-1)/2
	}
	//fmt.Println(h.c)
}

func (h *Heap) Pop() int{
	fmt.Println("FuncIn")
	res := -1
	if !h.IsEmpty(){
		res= h.c[0]
		h.c[0]=h.c[len(h.c)-1]
		h.c=h.c[:len(h.c)-1]
		idx:=0
		for idx<len(h.c) {
			left:=idx*2+1
			right:=idx*2+2
			if right >= len(h.c) && left >= len(h.c) {
				fmt.Println(h.c)
				fmt.Println(idx)
				fmt.Println("FuncOut1")
				return res
			}
			if right >= len(h.c) && h.c[idx] < h.c[left] {
				h.c[idx], h.c[left] = h.c[left], h.c[idx]
				idx = left
				continue
			}
			if h.c[idx]>=h.c[left] && h.c[idx]>=h.c[right]{
				fmt.Println(h.c)
				fmt.Printf("zhaoxiangyu %+v %+v %+v \n",idx,h.c[idx],h.c[left])
				fmt.Println(res)
				fmt.Println("FuncOut2")
				return res
			}
			if h.c[idx]>=h.c[left]{
				h.c[idx], h.c[right] = h.c[right], h.c[idx]
				idx = right
				continue
			}
			if h.c[idx]>=h.c[right]{
				h.c[idx], h.c[left] = h.c[left], h.c[idx]
				idx = left
				continue
			}
			if h.c[left]<h.c[right]{
				h.c[idx], h.c[right] = h.c[right], h.c[idx]
				idx = right
				continue
			}else {
				h.c[idx], h.c[left] = h.c[left], h.c[idx]
				idx = left
				continue
			}
		}
	}
	fmt.Println(h.c)
	fmt.Println("FuncOut")
	return res
}

func (h *Heap) IsEmpty() bool {
	return len(h.c)==0
}
func (h *Heap) Peek() int{
	if !h.IsEmpty(){
		return h.c[0]
	}
	return -1
}

func Test_Func2(t *testing.T)  {
	nums:=[]int{9,10,9,-7,-4,8,2,-6}
	fmt.Println(maxSlidingWindow(nums,5))
}
