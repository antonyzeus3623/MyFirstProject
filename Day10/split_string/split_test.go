package splitstring

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("abcbdef", "b")
	want := []string{"a", "c", "def"}
	if !reflect.DeepEqual(got, want) {
		//测试用例失败了
		t.Errorf("want:%v but got:%v\n", want, got)
	}
}

func Test2Split(t *testing.T) {
	got := Split("a:b:c:d", ":")
	want := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(got, want) {
		//测试用例失败了
		t.Errorf("want:%v but got:%v\n", want, got)
	}
}

func Test3Split(t *testing.T) {
	got := Split("abagababahkaj", "b")
	want := []string{"a", "aga", "a", "ahkaj"}
	if !reflect.DeepEqual(got, want) {
		//测试用例失败了
		t.Fatalf("want:%v but got:%v\n", want, got)
	}
}
