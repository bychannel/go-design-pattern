package option

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOption(t *testing.T) {
	// 传统模式创建
	// 为什么不使用属性的set方法？换句话说，option模式相对于set()的优点：
	// 如果只想在初始化的时候设置相关字段的话，提供了SetXx之类的方法就可以随时修改字段值，降低了封装性。
	// 如果对外提供的是接口类型而不是直接暴露该结构体的场景下，你就需要为接口类型添加很多的SetXx方法，
	// set方法混合在接口方法中，大大增加了接口的复杂度。
	x := newOption("google", "apple", 1000)
	assert.Equal(t, "google", x.A)
	assert.Equal(t, "apple", x.B)
	assert.Equal(t, 1000, x.C)

	// option模式无参创建，使用默认的参数配置
	x = newOption2()
	assert.Equal(t, "AAA", x.A)
	assert.Equal(t, "BBB", x.B)
	assert.Equal(t, 999, x.C)

	// option模式带参创建，自定义的参数替代默认参数，未给定的参数保持默认参数
	x = newOption2(WithA("baidu"), WithC(111))
	assert.Equal(t, "baidu", x.A)
	assert.Equal(t, "BBB", x.B)
	assert.Equal(t, 111, x.C)
}
