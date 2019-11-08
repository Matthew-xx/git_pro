package main

import (
	"testing"
)

func TestSubstr(t *testing.T)  {
	tests := []struct{s string ; ans int}{

		{"abdsssc",4},
		{"dictfffd",5},
		{"",0},
		{"一二三二一",3},
		{"上海自来水来自海上",5},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubstr(tt.s)
		if actual != tt.ans{
			t.Errorf("Got %d for input %s;"+"expected %d",actual,tt.s,tt.ans)
		}
	}
}

//性能测试,再用pprof查看分析消耗较大的程序并进行优化
func BenchmarkSubstr(b *testing.B)  {
	s := "上海自来水来自海上"
	ans := 5

	for i:=0; i<b.N ;i++  { //程序会自己找出N
		actual := lengthOfNonRepeatingSubstr(s)
		if actual != ans {
			b.Errorf("got %d for input %s;" + "ecpected %d",actual,s,ans)
		}
	}
}