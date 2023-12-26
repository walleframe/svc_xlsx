// Generate by wctl xlsx. DO NOT EDIT.
package example_global

import (
	"fmt"
	"sync/atomic"

	"github.com/walleframe/svc_xlsx"
	"github.com/walleframe/svc_xlsx/examples/xlsx/wpb"
)

const (
	ConfigName = "example_global"
)

var (
	globalData    atomic.Pointer[tableData]
	checkHandlers []func(*wpb.ExampleGlobal) error
)

// Combine struct from data/Examples.xlsx example_global_st
type tableData struct {
	// table example_global_st
	ExampleGlobal *wpb.ExampleGlobal
}

func init() {
	globalData.Store(new(tableData))
	svc_xlsx.RegAutoConfig("example_global", "data/Examples.xlsx", "example_global_st", nil, &loader{})
}

type loader struct{}

// NewContainer 新建container 返回数据指针
func (l *loader) NewContainer() interface{} {
	return new(wpb.ExampleGlobal)
}

// Check 检查数据
func (l *loader) Check(new interface{}) error {
	if new == nil {
		return fmt.Errorf("ptr is nil")
	}
	t, ok := new.(*wpb.ExampleGlobal)
	if !ok {
		return fmt.Errorf("error type")
	}

	for _, c := range checkHandlers {
		if err := c(t); err != nil {
			return err
		}
	}
	return nil
}

// Swap 交换内存地址
func (l *loader) Swap(new interface{}) {
	newT, ok := new.(*wpb.ExampleGlobal)
	if !ok {
		return
	}
	newL := &tableData{
		ExampleGlobal: newT,
	}

	globalData.Store(newL)
}

// RegisterCheckEntry 注册校验回调(用于更新数据前校验)
func RegisterCheckEntry(h func(*wpb.ExampleGlobal) error) {

	if h == nil {
		panic("empty preload handler")
	}

	checkHandlers = append(checkHandlers, h)
}

// Get() 获取数据 from data/Examples.xlsx example_global_st
func Get() *wpb.ExampleGlobal {
	data := globalData.Load()
	if data == nil || data.ExampleGlobal == nil {
		return nil
	}
	return data.ExampleGlobal
}
