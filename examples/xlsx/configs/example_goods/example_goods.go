// Generate by wctl xlsx. DO NOT EDIT.
package example_goods

import (
	"fmt"
	"sync/atomic"

	"github.com/walleframe/svc_xlsx"
	"github.com/walleframe/svc_xlsx/examples/xlsx/wpb"
)

const (
	ConfigName = "example_goods"
)

var (
	globalData    atomic.Pointer[tableData]
	checkHandlers []func(*wpb.ExampleGoodsContainer) error
)

// Combine struct from data/Examples.xlsx example_goods_cfgs
type tableData struct {
	// table example_goods_cfgs
	ExampleGoods *wpb.ExampleGoodsContainer
}

func init() {
	globalData.Store(new(tableData))
	svc_xlsx.RegAutoConfig("example_goods", "data/Examples.xlsx", "example_goods_cfgs", nil, &loader{})
}

type loader struct{}

// NewContainer 新建container 返回数据指针
func (l *loader) NewContainer() interface{} {
	return new(wpb.ExampleGoodsContainer)
}

// Check 检查数据
func (l *loader) Check(new interface{}) error {
	if new == nil {
		return fmt.Errorf("ptr is nil")
	}
	t, ok := new.(*wpb.ExampleGoodsContainer)
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
	newT, ok := new.(*wpb.ExampleGoodsContainer)
	if !ok {
		return
	}
	newL := &tableData{
		ExampleGoods: newT,
	}

	globalData.Store(newL)
}

// RegisterCheckEntry 注册校验回调(用于更新数据前校验)
func RegisterCheckEntry(h func(*wpb.ExampleGoodsContainer) error) {

	if h == nil {
		panic("empty preload handler")
	}

	checkHandlers = append(checkHandlers, h)
}

// Get 获取全部数据 data/Examples.xlsx example_goods_cfgs
func Get() []*wpb.ExampleGoods {
	data := globalData.Load()
	if data == nil || data.ExampleGoods == nil {
		return nil
	}
	return data.ExampleGoods.Data
}

// Count 获取配置总个数
func Count() int {
	data := globalData.Load()
	if data == nil || data.ExampleGoods == nil {
		return 0
	}
	return len(data.ExampleGoods.Data)
}

// Range 遍历
func Range(filter func(index int, val *wpb.ExampleGoods) bool) {
	data := globalData.Load()
	if data == nil || data.ExampleGoods == nil {
		return
	}
	for index, v := range data.ExampleGoods.Data {
		if !filter(index, v) {
			return
		}
	}
}

// GetByIndex 根据下标获取数据
func GetByIndex(index int) *wpb.ExampleGoods {
	data := globalData.Load()
	if data == nil || data.ExampleGoods == nil {
		return nil
	}
	if index < 0 || index > len(data.ExampleGoods.Data) {
		return nil
	}
	return data.ExampleGoods.Data[index]
}

// GetByFilter 根据过滤器获取批量数据
func GetByFilter(filter func(val *wpb.ExampleGoods) bool) (ret []*wpb.ExampleGoods) {
	for _, v := range Get() {
		if filter(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// GetOneByFilter 根据过滤器获取单个数据
func GetOneByFilter(filter func(val *wpb.ExampleGoods) bool) *wpb.ExampleGoods {
	for _, v := range Get() {
		if filter(v) {
			return v
		}
	}
	return nil
}
