package collections

type HeapInterface interface {
	Swap(i, j int)
	Compare(i, j int) bool
	Len() int
}

func Init(heap HeapInterface) {
	n := heap.Len()
	for i := n / 2; i >= 0; i-- {
		adjustToBottom(i, n, heap)
	}
}

func Add(data interface{}) {

}

/**
从数组尾部开始向上调节
*/
func adjustToHead(idx int, heap HeapInterface) {
	var adjustHeap func(idx int)
	adjustHeap = func(idx int) {
		if idx > 0 {
			pidx := (idx - 1) / 2
			if heap.Compare(idx, pidx) {
				heap.Swap(idx, pidx)
				adjustHeap(pidx)
			}
		}
	}
	adjustHeap(idx)
}

/**
从顶部开始往下调整
*/
func adjustToBottom(idx, len int, heap HeapInterface) {
	var adjustBottom func(idx int)
	adjustBottom = func(idx int) {
		tarIdx := idx
		if 2*idx+1 < len {
			if heap.Compare(tarIdx, 2*idx+1) {
				tarIdx = 2*idx + 1
			}
		}
		if 2*idx+2 < len {
			if heap.Compare(tarIdx, 2*idx+2) {
				tarIdx = 2*idx + 2
			}
		}
		if idx != tarIdx {
			heap.Swap(tarIdx, idx)
			adjustBottom(tarIdx)
		}
	}
	adjustBottom(idx)
}
