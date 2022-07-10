package option

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

type OptionFunc func(*Option)

func WithA(a string) OptionFunc {
	return func(o *Option) {
		o.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(o *Option) {
		o.B = b
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
func newOption2(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}
	return
}
