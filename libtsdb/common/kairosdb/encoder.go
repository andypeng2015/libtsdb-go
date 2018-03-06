package kairosdb

import (
	"strconv"

	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	"github.com/libtsdb/libtsdb-go/libtsdb/util/bytesutil"
)

type TelnetEncoder struct {
	bytesutil.Buffer
}

var _ common.Encoder = (*JsonEncoder)(nil)

// JsonEncoder support mix of single point and series
type JsonEncoder struct {
	bytesutil.Buffer
	finalized bool
}

func NewJsonEncoder() *JsonEncoder {
	e := &JsonEncoder{}
	e.Reset()
	return e
}

func (e *JsonEncoder) Reset() {
	e.Buffer.Reset()
	// start of json array
	e.Buf = append(e.Buf, '[')
}

func (e *JsonEncoder) Bytes() []byte {
	if e.finalized {
		return e.Buffer.Bytes()
	}
	// replace last extra comma with end of json array
	e.Buf[len(e.Buf)-1] = ']'
	e.finalized = true
	return e.Buffer.Bytes()
}

func (e *JsonEncoder) WritePointIntTagged(p *pb.PointIntTagged) {
	e.Buf = append(e.Buf, `{"name":"`...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, `",`...)
	e.Buf = append(e.Buf, `"timestamp":`...)
	e.Buf = strconv.AppendInt(e.Buf, p.Point.T, 10)
	e.Buf = append(e.Buf, `,"value":`...)
	e.Buf = strconv.AppendInt(e.Buf, p.Point.V, 10)
	e.Buf = append(e.Buf, `,"tags":{`...)
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, '"')
		e.Buf = append(e.Buf, tag.K...)
		e.Buf = append(e.Buf, `":"`...)
		e.Buf = append(e.Buf, tag.V...)
		e.Buf = append(e.Buf, `",`...)
	}
	e.Buf[len(e.Buf)-1] = '}'
	e.Buf = append(e.Buf, `},`...)
}

func (e *JsonEncoder) WritePointDoubleTagged(p *pb.PointDoubleTagged) {
	e.Buf = append(e.Buf, `{"name":"`...)
	e.Buf = append(e.Buf, p.Name...)
	e.Buf = append(e.Buf, `",`...)
	e.Buf = append(e.Buf, `"timestamp":`...)
	e.Buf = strconv.AppendInt(e.Buf, p.Point.T, 10)
	e.Buf = append(e.Buf, `,"value":`...)
	// NOTE: the only difference with write int
	e.Buf = strconv.AppendFloat(e.Buf, p.Point.V, 'f', -1, 64)
	e.Buf = append(e.Buf, `,"tags":{`...)
	for _, tag := range p.Tags {
		e.Buf = append(e.Buf, '"')
		e.Buf = append(e.Buf, tag.K...)
		e.Buf = append(e.Buf, `":"`...)
		e.Buf = append(e.Buf, tag.V...)
		e.Buf = append(e.Buf, `",`...)
	}
	e.Buf[len(e.Buf)-1] = '}'
	e.Buf = append(e.Buf, `},`...)
}