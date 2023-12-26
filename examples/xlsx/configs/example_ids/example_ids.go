// Generate by wctl xlsx. DO NOT EDIT.
package example_ids

import (
	"fmt"
	"sync/atomic"

	"github.com/walleframe/svc_xlsx"
	"github.com/walleframe/svc_xlsx/examples/xlsx/wpb"
)

const (
	ConfigName = "example_ids"
)

type KeyType = uint32

var (
	globalData    atomic.Pointer[tableData]
	checkHandlers []func(*wpb.ExampleIdsContainer) error
)

// Combine struct from data/Examples.xlsx example_ids_cfgs
type tableData struct {
	// table example_ids_cfgs
	ExampleIds *wpb.ExampleIdsContainer
	MapData    map[KeyType]*wpb.ExampleIds
}

func init() {
	globalData.Store(new(tableData))
	svc_xlsx.RegAutoConfig("example_ids", "data/Examples.xlsx", "example_ids_cfgs", nil, &loader{})
}

type loader struct{}

// NewContainer 新建container 返回数据指针
func (l *loader) NewContainer() interface{} {
	return new(wpb.ExampleIdsContainer)
}

// Check 检查数据
func (l *loader) Check(new interface{}) error {
	if new == nil {
		return fmt.Errorf("ptr is nil")
	}
	t, ok := new.(*wpb.ExampleIdsContainer)
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
	newT, ok := new.(*wpb.ExampleIdsContainer)
	if !ok {
		return
	}
	newL := &tableData{
		ExampleIds: newT, MapData: make(map[KeyType]*wpb.ExampleIds, len(newT.Data)),
	}

	for _, item := range newT.Data {
		newL.MapData[item.GoodsId] = item
	}

	globalData.Store(newL)
}

// RegisterCheckEntry 注册校验回调(用于更新数据前校验)
func RegisterCheckEntry(h func(*wpb.ExampleIdsContainer) error) {

	if h == nil {
		panic("empty preload handler")
	}

	checkHandlers = append(checkHandlers, h)
}

// Get 获取全部数据 data/Examples.xlsx example_ids_cfgs
func Get() []*wpb.ExampleIds {
	data := globalData.Load()
	if data == nil || data.ExampleIds == nil {
		return nil
	}
	return data.ExampleIds.Data
}

// GetByID 根据索引获取数据
func GetByID(id KeyType) *wpb.ExampleIds {
	data := globalData.Load()
	if data == nil || data.MapData == nil {
		return nil
	}
	v, ok := data.MapData[id]
	if !ok {
		return nil
	}
	return v
}

// Count 获取配置总个数
func Count() int {
	data := globalData.Load()
	if data == nil || data.ExampleIds == nil {
		return 0
	}
	return len(data.ExampleIds.Data)
}

// Range 遍历
func Range(filter func(index int, val *wpb.ExampleIds) bool) {
	data := globalData.Load()
	if data == nil || data.ExampleIds == nil {
		return
	}
	for index, v := range data.ExampleIds.Data {
		if !filter(index, v) {
			return
		}
	}
}

// GetByIndex 根据下标获取数据
func GetByIndex(index int) *wpb.ExampleIds {
	data := globalData.Load()
	if data == nil || data.ExampleIds == nil {
		return nil
	}
	if index < 0 || index > len(data.ExampleIds.Data) {
		return nil
	}
	return data.ExampleIds.Data[index]
}

// GetByFilter 根据过滤器获取批量数据
func GetByFilter(filter func(val *wpb.ExampleIds) bool) (ret []*wpb.ExampleIds) {
	for _, v := range Get() {
		if filter(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// GetOneByFilter 根据过滤器获取单个数据
func GetOneByFilter(filter func(val *wpb.ExampleIds) bool) *wpb.ExampleIds {
	for _, v := range Get() {
		if filter(v) {
			return v
		}
	}
	return nil
}
