package utils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/spf13/cast"
)

type TSafeText string

func NewTSafeText(value interface{}) TSafeText {
	val := TSafeText(cast.ToString(value))
	return val
}

func (this TSafeText) MarshalJSON() ([]byte, error) {
	var buff bytes.Buffer
	var valBytes []byte

	valRawStr := cast.ToString(this)

	valBytes, _ = json.Marshal(valRawStr)
	buff.Write(valBytes)

	return buff.Bytes(), nil
}

func (this *TSafeText) UnmarshalJSON(b []byte) error {
	var result interface{}
	err := json.Unmarshal(b, &result)
	if err != nil {
		return err
	}

	*this = TSafeText(cast.ToString(result))

	return nil
}

func (this TSafeText) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(cast.ToString(this), start)

	return nil
}

func (this *TSafeText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var val interface{}
	d.DecodeElement(&val, &start)

	*this = TSafeText(cast.ToString(val))

	return nil
}

func (this TSafeText) IsNil() bool {
	return string(this) == ""
}

func (this TSafeText) ToString() string {
	return string(this)
}

func (this TSafeText) ToValueOfString() *string {
	if this.IsNil() {
		return nil
	}

	returnValue := this.ToString()
	return &returnValue
}

func (this TSafeText) ToJsonNumber() json.Number {
	return json.Number(this.ToString())
}

func (this TSafeText) ToValueOfJsonNumber() *json.Number {
	if this.IsNil() {
		return nil
	}

	returnValue := this.ToJsonNumber()
	return &returnValue
}

func (this TSafeText) ToInt64() int64 {
	return cast.ToInt64(this.ToFloat64())
}

func (this TSafeText) ToValueOfInt64() *int64 {
	if this.IsNil() {
		return nil
	}

	returnValue := this.ToInt64()
	return &returnValue
}

func (this TSafeText) ToInt() int {
	return cast.ToInt(this.ToFloat64())
}

func (this TSafeText) ToValueOfInt() *int {
	if this.IsNil() {
		return nil
	}

	returnValue := this.ToInt()
	return &returnValue
}

func (this TSafeText) ToFloat32() float32 {
	return cast.ToFloat32(this.ToString())
}

func (this TSafeText) ToValueOfFloat32() *float32 {
	if this.IsNil() {
		return nil
	}

	returnValue := this.ToFloat32()
	return &returnValue
}

func (this TSafeText) ToFloat64() float64 {
	return cast.ToFloat64(this.ToString())
}

func (this TSafeText) ToValueOfFloat64() *float64 {
	if this.IsNil() {
		return nil
	}

	returnValue := this.ToFloat64()
	return &returnValue
}

func (this TSafeText) ToSliceOfString() []string {
	return strings.Split(this.ToString(), ",")
}

func (this TSafeText) ToSliceOfInt64() []int64 {
	var arrInt64 []int64
	for _, v := range strings.Split(this.ToString(), ",") {
		int64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		arrInt64 = append(arrInt64, int64)
	}
	return arrInt64
}
