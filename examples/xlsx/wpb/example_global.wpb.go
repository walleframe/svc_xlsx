// Code generated by protoc-gen-gopb. DO NOT EDIT.
// Code generated by wpb. DO NOT EDIT.

package wpb

import (
	"errors"

	"github.com/walleframe/walle/util/protowire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// example_global generate from example_global_st in data/Examples.xlsx
type ExampleGlobal struct {
	// register_limit 注册限制
	RegisterLimit int32 `json:"register_limit,omitempty"`
	// friend_limit 好友限制
	FriendLimit int32 `json:"friend_limit,omitempty"`
	// register_gold 注册赠送
	RegisterGold int64 `json:"register_gold,omitempty"`
}

func (x *ExampleGlobal) Reset() {
	*x = ExampleGlobal{}
}

// MarshalObject marshal data to []byte
func (x *ExampleGlobal) MarshalObject() (data []byte, err error) {
	data = make([]byte, 0, x.MarshalSize())
	return x.MarshalObjectTo(data)
}

// MarshalSize calc marshal data need space
func (x *ExampleGlobal) MarshalSize() (size int) {
	if x.RegisterLimit != 0 {
		// 1 = protowire.SizeTag(1)
		size += 1 + protowire.SizeVarint(uint64(x.RegisterLimit))
	}
	if x.FriendLimit != 0 {
		// 1 = protowire.SizeTag(2)
		size += 1 + protowire.SizeVarint(uint64(x.FriendLimit))
	}
	if x.RegisterGold != 0 {
		// 1 = protowire.SizeTag(3)
		size += 1 + protowire.SizeVarint(uint64(x.RegisterGold))
	}
	return
}

// MarshalObjectTo marshal data to []byte
func (x *ExampleGlobal) MarshalObjectTo(buf []byte) (data []byte, err error) {
	data = buf
	if x.RegisterLimit != 0 {
		// data = protowire.AppendTag(data, 1, protowire.VarintType) => 00001000
		data = append(data, 0x8)
		data = protowire.AppendVarint(data, uint64(x.RegisterLimit))
	}
	if x.FriendLimit != 0 {
		// data = protowire.AppendTag(data, 2, protowire.VarintType) => 00010000
		data = append(data, 0x10)
		data = protowire.AppendVarint(data, uint64(x.FriendLimit))
	}
	if x.RegisterGold != 0 {
		// data = protowire.AppendTag(data, 3, protowire.VarintType) => 00011000
		data = append(data, 0x18)
		data = protowire.AppendVarint(data, uint64(x.RegisterGold))
	}
	return
}

// UnmarshalObject unmarshal data from []byte
func (x *ExampleGlobal) UnmarshalObject(data []byte) (err error) {
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
				err = errors.New("parse ExampleGlobal.RegisterLimit ID:1 : invalid varint value")
				return
			}
			index += cnt
			x.RegisterLimit = int32(v)
		case 2:
			v, cnt := protowire.ConsumeVarint(data[index:])
			if cnt < 1 {
				err = errors.New("parse ExampleGlobal.FriendLimit ID:2 : invalid varint value")
				return
			}
			index += cnt
			x.FriendLimit = int32(v)
		case 3:
			v, cnt := protowire.ConsumeVarint(data[index:])
			if cnt < 1 {
				err = errors.New("parse ExampleGlobal.RegisterGold ID:3 : invalid varint value")
				return
			}
			index += cnt
			x.RegisterGold = int64(v)
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

func (x *ExampleGlobal) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("RegisterLimit", x.RegisterLimit)
	enc.AddInt32("FriendLimit", x.FriendLimit)
	enc.AddInt64("RegisterGold", x.RegisterGold)
	return nil
}

type ZapArrayExampleGlobal []*ExampleGlobal

func (x ZapArrayExampleGlobal) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, v := range x {
		ae.AppendObject(v)
	}
	return nil
}

func LogArrayExampleGlobal(name string, v []*ExampleGlobal) zap.Field {
	return zap.Array(name, ZapArrayExampleGlobal(v))
}
