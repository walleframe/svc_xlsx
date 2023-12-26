// Code generated by protoc-gen-gopb. DO NOT EDIT.
// Code generated by wpb. DO NOT EDIT.

package wpb

import (
	"errors"

	"github.com/walleframe/walle/util/protowire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// example_goods generate from example_goods_cfgs in data/Examples.xlsx
type ExampleGoods struct {
	// GoodsID 物品ID
	GoodsId uint32 `json:"goods_id,omitempty"`
	// GoodsType 物品类型
	GoodsType int32 `json:"goods_type,omitempty"`
	// DiamondPrice 出售钻石价格
	DiamondPrice int64 `json:"diamond_price,omitempty"`
	// GoldPrice 出售金币价格
	GoldPrice int64 `json:"gold_price,omitempty"`
}

func (x *ExampleGoods) Reset() {
	*x = ExampleGoods{}
}

// MarshalObject marshal data to []byte
func (x *ExampleGoods) MarshalObject() (data []byte, err error) {
	data = make([]byte, 0, x.MarshalSize())
	return x.MarshalObjectTo(data)
}

// MarshalSize calc marshal data need space
func (x *ExampleGoods) MarshalSize() (size int) {
	if x.GoodsId != 0 {
		// 1 = protowire.SizeTag(1)
		size += 1 + protowire.SizeVarint(uint64(x.GoodsId))
	}
	if x.GoodsType != 0 {
		// 1 = protowire.SizeTag(2)
		size += 1 + protowire.SizeVarint(uint64(x.GoodsType))
	}
	if x.DiamondPrice != 0 {
		// 1 = protowire.SizeTag(3)
		size += 1 + protowire.SizeVarint(uint64(x.DiamondPrice))
	}
	if x.GoldPrice != 0 {
		// 1 = protowire.SizeTag(4)
		size += 1 + protowire.SizeVarint(uint64(x.GoldPrice))
	}
	return
}

// MarshalObjectTo marshal data to []byte
func (x *ExampleGoods) MarshalObjectTo(buf []byte) (data []byte, err error) {
	data = buf
	if x.GoodsId != 0 {
		// data = protowire.AppendTag(data, 1, protowire.VarintType) => 00001000
		data = append(data, 0x8)
		data = protowire.AppendVarint(data, uint64(x.GoodsId))
	}
	if x.GoodsType != 0 {
		// data = protowire.AppendTag(data, 2, protowire.VarintType) => 00010000
		data = append(data, 0x10)
		data = protowire.AppendVarint(data, uint64(x.GoodsType))
	}
	if x.DiamondPrice != 0 {
		// data = protowire.AppendTag(data, 3, protowire.VarintType) => 00011000
		data = append(data, 0x18)
		data = protowire.AppendVarint(data, uint64(x.DiamondPrice))
	}
	if x.GoldPrice != 0 {
		// data = protowire.AppendTag(data, 4, protowire.VarintType) => 00100000
		data = append(data, 0x20)
		data = protowire.AppendVarint(data, uint64(x.GoldPrice))
	}
	return
}

// UnmarshalObject unmarshal data from []byte
func (x *ExampleGoods) UnmarshalObject(data []byte) (err error) {
	index := 0
	for index < len(data) {
		num, typ, cnt := protowire.ConsumeTag(data[index:])
		if num == 0 {
			err = errors.New("invalid tag")
			return
		}

		index += cnt
		switch num {
		case 1:
			v, cnt := protowire.ConsumeVarint(data[index:])
			if cnt < 1 {
				err = errors.New("parse ExampleGoods.GoodsId ID:1 : invalid varint value")
				return
			}
			index += cnt
			x.GoodsId = uint32(v)
		case 2:
			v, cnt := protowire.ConsumeVarint(data[index:])
			if cnt < 1 {
				err = errors.New("parse ExampleGoods.GoodsType ID:2 : invalid varint value")
				return
			}
			index += cnt
			x.GoodsType = int32(v)
		case 3:
			v, cnt := protowire.ConsumeVarint(data[index:])
			if cnt < 1 {
				err = errors.New("parse ExampleGoods.DiamondPrice ID:3 : invalid varint value")
				return
			}
			index += cnt
			x.DiamondPrice = int64(v)
		case 4:
			v, cnt := protowire.ConsumeVarint(data[index:])
			if cnt < 1 {
				err = errors.New("parse ExampleGoods.GoldPrice ID:4 : invalid varint value")
				return
			}
			index += cnt
			x.GoldPrice = int64(v)
		default: // skip fields
			cnt = protowire.ConsumeFieldValue(num, typ, data[index:])
			if cnt < 0 {
				return protowire.ParseError(cnt)
			}
			index += cnt
		}
	}

	return
}

func (x *ExampleGoods) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddUint32("GoodsId", x.GoodsId)
	enc.AddInt32("GoodsType", x.GoodsType)
	enc.AddInt64("DiamondPrice", x.DiamondPrice)
	enc.AddInt64("GoldPrice", x.GoldPrice)
	return nil
}

type ZapArrayExampleGoods []*ExampleGoods

func (x ZapArrayExampleGoods) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, v := range x {
		ae.AppendObject(v)
	}
	return nil
}

func LogArrayExampleGoods(name string, v []*ExampleGoods) zap.Field {
	return zap.Array(name, ZapArrayExampleGoods(v))
}

// example_goods generate from example_goods_cfgs in data/Examples.xlsx
type ExampleGoodsContainer struct {
	Data []*ExampleGoods `json:"data,omitempty"` // table
}

func (x *ExampleGoodsContainer) Reset() {
	*x = ExampleGoodsContainer{}
}

// MarshalObject marshal data to []byte
func (x *ExampleGoodsContainer) MarshalObject() (data []byte, err error) {
	data = make([]byte, 0, x.MarshalSize())
	return x.MarshalObjectTo(data)
}

// MarshalSize calc marshal data need space
func (x *ExampleGoodsContainer) MarshalSize() (size int) {
	if x.Data != nil {
		// 1 = protowire.SizeTag(1)
		size += 1 * len(x.Data)
		for k := 0; k < len(x.Data); k++ {
			size += protowire.SizeBytes(x.Data[k].MarshalSize())
		}
	}
	return
}

// MarshalObjectTo marshal data to []byte
func (x *ExampleGoodsContainer) MarshalObjectTo(buf []byte) (data []byte, err error) {
	data = buf
	if x.Data != nil {
		for _, item := range x.Data {
			// data = protowire.AppendTag(data, 1, protowire.BytesType) => 00001010
			data = append(data, 0xa)
			data = protowire.AppendVarint(data, uint64(item.MarshalSize()))
			data, err = item.MarshalObjectTo(data)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalObject unmarshal data from []byte
func (x *ExampleGoodsContainer) UnmarshalObject(data []byte) (err error) {
	index := 0
	for index < len(data) {
		num, typ, cnt := protowire.ConsumeTag(data[index:])
		if num == 0 {
			err = errors.New("invalid tag")
			return
		}

		index += cnt
		switch num {
		case 1:
			if typ != protowire.BytesType {
				err = errors.New("parse ExampleGoodsContainer.Data ID:1 : invalid repeated tag value")
				return
			}
			buf, cnt := protowire.ConsumeBytes(data[index:])
			if buf == nil {
				err = errors.New("parse ExampleGoodsContainer.Data ID:1 : invalid len value")
				return
			}
			index += cnt
			if x.Data == nil {
				x.Data = make([]*ExampleGoods, 0, 2)
			}
			item := &ExampleGoods{}
			err = item.UnmarshalObject(buf)
			if err != nil {
				return
			}
			x.Data = append(x.Data, item)
		default: // skip fields
			cnt = protowire.ConsumeFieldValue(num, typ, data[index:])
			if cnt < 0 {
				return protowire.ParseError(cnt)
			}
			index += cnt
		}
	}

	return
}

func (x *ExampleGoodsContainer) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddArray("Data", zapcore.ArrayMarshalerFunc(func(ae zapcore.ArrayEncoder) error {
		for _, v := range x.Data {
			ae.AppendObject(v)
		}
		return nil
	}))
	return nil
}

type ZapArrayExampleGoodsContainer []*ExampleGoodsContainer

func (x ZapArrayExampleGoodsContainer) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, v := range x {
		ae.AppendObject(v)
	}
	return nil
}

func LogArrayExampleGoodsContainer(name string, v []*ExampleGoodsContainer) zap.Field {
	return zap.Array(name, ZapArrayExampleGoodsContainer(v))
}
