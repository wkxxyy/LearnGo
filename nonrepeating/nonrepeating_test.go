package main

import "testing"

func TestSubstr(t *testing.T) { //测试是否是对的
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwkew", 3},

		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lenthOfNoneRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("Got %d for input %s; "+"expected %d ", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) { //测试性能
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d ", len(s))
	ans := 8
	b.ResetTimer() //这个之前的就是不算时间的

	for i := 0; i < b.N; i++ {

		actual := lenthOfNoneRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input %s; "+"expected %d ", actual, s, ans)
		}
	}
}
