package intset

import "testing"

func TestAdd(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{1, "{1}"},
		{3, "{1 3}"},
		{4, "{1 3 4}"},
		{4, "{1 3 4}"},
		{3, "{1 3 4}"},
		{10, "{1 3 4 10}"},
		{2, "{1 2 3 4 10}"},
		{100, "{1 2 3 4 10 100}"},
		{10, "{1 2 3 4 10 100}"},
	}

	var x IntSet
	for _, test := range tests {
		x.Add(test.input)
		if x.String() != test.want {
			t.Errorf("intset.Add(%d) = %s, expected = %s", test.input, x.String(), test.want)
		}

	}
}

func TestHas(t *testing.T) {
	var tests = []struct {
		needle int
		search IntSet
		want   bool
	}{
		{1, IntSet{words: bitUint([]int{1, 2, 3, 4, 5})}, true},
		{2, IntSet{words: bitUint([]int{1, 100, 200, 300})}, false},
		{10, IntSet{words: bitUint([]int{10})}, true},
		{1, IntSet{words: bitUint([]int{1, 2, 2, 2, 2})}, true},
		{777, IntSet{words: bitUint([]int{77, 7777, 77777})}, false},
	}

	for _, test := range tests {
		if result := test.search.Has(test.needle); result != test.want {
			t.Errorf("intset.Has(%d) = %v in %v", test.needle, result, test.search.String())
		}

	}
}

func TestUnion(t *testing.T) {
	var tests = []struct {
		ints1 []int
		ints2 []int
		want  string
	}{
		{[]int{1, 2}, []int{1, 2}, "{1 2}"},
		{[]int{1, 2, 3, 4}, []int{1, 2, 5}, "{1 2 3 4 5}"},
		{[]int{1}, []int{10}, "{1 10}"},
		{[]int{1}, []int{}, "{1}"},
		{[]int{10, 11, 9}, []int{5, 4, 3}, "{3 4 5 9 10 11}"},
	}

	// var intset1 IntSet
	// var intset2 IntSet
	for _, test := range tests {
		intset1 := IntSet{[]uint64{}}
		intset2 := IntSet{[]uint64{}}

		for _, val := range test.ints1 {
			intset1.Add(val)
		}

		for _, val := range test.ints2 {
			intset2.Add(val)
		}
		oldvalue := intset1.String()
		intset1.UnionWith(&intset2)
		if intset1.String() != test.want {
			t.Errorf("%v UnionWith (%v) = %v", oldvalue, intset2.String(), intset1.String())
		}

	}
}

func bitUint(x []int) []uint64 {
	var result []uint64
	for _, num := range x {
		result = append(result, 1<<uint(num%64))
	}

	return result
}
