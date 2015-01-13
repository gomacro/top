package top

import (
	"github.com/gomacro/heap/8/heap"
)

func Top(ts0 *[1]uintptr, compar func(*uint8, *uint8) int, dst []uint8, src []uint8) {
	incr := int((*ts0)[0])

	copy(dst, src)

	heap.Heapify(ts0, compar, dst, dst)

	for i := (len(dst)/incr); i < (len(src)/incr); i++ {
		if compar(&dst[0], &src[i*incr]) < 0 {
			for q := 0; q < incr; q++ { // copy
				dst[q] = src[incr*i+q]
			}
			heap.Fix(ts0, compar, dst, 0)
		}
	}
}
