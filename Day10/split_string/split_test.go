package splitstring

import (
	"reflect"
	"testing"
)

func Test4Split(t *testing.T) {
	//定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	//定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c:d", sep: ":", want: []string{"a", "b", "c", "d"}},
		{input: "a:b:c:d", sep: ",", want: []string{"a:b:c:d"}},
		{input: "abcd", sep: "bc", want: []string{"a", "cd"}},
		// {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有"}},
	}
	//遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want:%v but got:%v\n", tc.want, got)
		}
	}
}

//BenchmarkSplit基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
