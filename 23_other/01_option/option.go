package option

// 选项或配置创建模式，一般用来初始化配置项
type Option struct {
	A string
	B string
	C int
}

// 传统的创建模式
func newOption(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}

// 定义设置option的方法类型
// 参数是option结构体，实际上传入的是默认的option属性
type OptionFunc func(*Option)

// 定义属性的设置方法，那个属性需要设置就创建对应的设置方法
func WithA(a string) OptionFunc {
	return func(o *Option) {
		o.A = a
	}
}

func WithC(c int) OptionFunc {
	return func(o *Option) {
		o.C = c
	}
}

var defaultOption = &Option{
	A: "AAA",
	B: "BBB",
	C: 999,
}

// option模式下的创建方法
func newOption2(opts ...OptionFunc) *Option {
	for _, o := range opts {
		o(defaultOption)
	}
	return defaultOption
}
