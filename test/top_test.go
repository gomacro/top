// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

package top_test

import (
	"fmt"
	"github.com/gomacro/sort/compare"
	"github.com/gomacro/top/unsafe/top"
	"sort"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////

type int64slice []int64

func (p int64slice) Len() int           { return len(p) }
func (p int64slice) Less(i, j int) bool { return p[i] < p[j] }
func (p int64slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

////////////////////////////////////////////////////////////////////////////////

func random(prng *[2]uint64) uint64 {
	s1 := prng[0]
	s0 := prng[1]
	prng[0] = s0
	s1 ^= s1 << 23 // a
	prng[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26))
	return prng[1] + s0 // b, c
}

func fillu(data []int64, seed *[2]uint64) {
	for i := 0; i < len(data); i++ {
		data[i] = int64(random(seed))
	}
}

const debug = true
const shortdatasize = 100
const datasize = 10000000

func BenchmarkMyTopLarge_Random(b *testing.B) {
	b.StopTimer()

	fmt.Printf("")

	var seed = [2]uint64{0x13371337, 0x1337beef}
	n := datasize
	if testing.Short() {
		n /= shortdatasize
	}
	data := make([]int64, n)
	sample := make([]int64, 100)
	fillu(data, &seed)
	if sort.IsSorted(int64slice(data)) {
		b.Fatalf("terrible rand.rand")
	}

	for n := 0; n < b.N; n++ {

		b.StartTimer()

		top.Top(compare.Int64, sample, data)

		b.StopTimer()

		if false {

			if !sort.IsSorted(int64slice(sample)) {
				b.Errorf("sort didn't sort - 1M ints")
			}
		}
		fillu(data, &seed)

	}
}

func TestCustom0(t *testing.T) {
	type Movie struct {
		Title     string
		BoxOffice int64
	}

	compar_asc := func(l, r *Movie) int {
		return compare.Int64(&l.BoxOffice, &r.BoxOffice)
	}

	compar_desc := func(l, r *Movie) int {
		return compare.Int64(&r.BoxOffice, &l.BoxOffice)
	}

	data := []Movie {
		{"Iron Man", 5000000},
		{"Independence Day", 1000000},
		{"Fargo", 3000000},
		{"Django Unchained", 9000000},
		{"WALL-E", 4000000},
	}

	sample_asc := make([]Movie, 3)
	sample_desc := make([]Movie, 3)

	top.Top(compar_asc, sample_asc, data)
	top.Top(compar_desc, sample_desc, data)

	for _, movie := range sample_asc {
		if movie == data[1] || movie == data[2] {
			t.Errorf("found %v in asc", movie)
		}
	}
	for _, movie := range sample_desc {
		if movie == data[0] || movie == data[3] {
			t.Errorf("found %v in desc", movie)
		}
	}
}
