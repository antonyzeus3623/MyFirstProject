package snow

// GetNum 获取输入数字
type GetNum struct {
	Target int
	Self   int
}

// NewGetNum 结构体初始化
func NewGetNum(a int, b int) *GetNum {
	return &GetNum{
		Target: a,
		Self:   b,
	}
}

// Add 加
func (s *GetNum) Add(new *GetNum) int {
	Sum := new.Target + new.Self
	return Sum
}

// Minus 减
func (s *GetNum) Minus(new *GetNum) int {
	Sum := new.Target - new.Self
	return Sum
}

// Take 乘法
func (s *GetNum) Take(new *GetNum) int {
	Sum := new.Target * new.Self
	return Sum
}

// Division 除法
func (s *GetNum) Division(new *GetNum) int {
	Sum := new.Target / new.Self
	return Sum
}
