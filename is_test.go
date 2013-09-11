package objx

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIsNil(t *testing.T) {

	n := New(nil)

	assert.True(t, n.IsNil())

	n.obj = "something"

	assert.False(t, n.IsNil())

}

func TestIsKind(t *testing.T) {

	var o *O

	o = New(bool(true))
	assert.True(t, o.IsKind(reflect.Bool))

	o = New(bool(false))
	assert.True(t, o.IsKind(reflect.Bool))

	o = New(int(1))
	assert.True(t, o.IsKind(reflect.Int))

	o = New(int8(1))
	assert.True(t, o.IsKind(reflect.Int8))

	o = New(int16(1))
	assert.True(t, o.IsKind(reflect.Int16))

	o = New(int32(1))
	assert.True(t, o.IsKind(reflect.Int32))

	o = New(int64(1))
	assert.True(t, o.IsKind(reflect.Int64))

	o = New(uint(1))
	assert.True(t, o.IsKind(reflect.Uint))

	o = New(uint8(1))
	assert.True(t, o.IsKind(reflect.Uint8))

	o = New(uint16(1))
	assert.True(t, o.IsKind(reflect.Uint16))

	o = New(uint32(1))
	assert.True(t, o.IsKind(reflect.Uint32))

	o = New(uint64(1))
	assert.True(t, o.IsKind(reflect.Uint64))

	o = New(float32(1))
	assert.True(t, o.IsKind(reflect.Float32))

	o = New(float64(1))
	assert.True(t, o.IsKind(reflect.Float64))

	o = New(complex64(1))
	assert.True(t, o.IsKind(reflect.Complex64))

	o = New(complex128(1))
	assert.True(t, o.IsKind(reflect.Complex128))

	o = New(string("1"))
	assert.True(t, o.IsKind(reflect.String))

	o = New(func() {})
	assert.True(t, o.IsKind(reflect.Func))

}

func TestIsSpecificTypes(t *testing.T) {

	var o *O

	o = New(bool(true))
	assert.True(t, o.IsBool())

	o = New(bool(false))
	assert.True(t, o.IsBool())

	o = New(int(1))
	assert.True(t, o.IsInt())

	o = New(int8(1))
	assert.True(t, o.IsInt8())

	o = New(int16(1))
	assert.True(t, o.IsInt16())

	o = New(int32(1))
	assert.True(t, o.IsInt32())

	o = New(int64(1))
	assert.True(t, o.IsInt64())

	o = New(uint(1))
	assert.True(t, o.IsUint())

	o = New(uint8(1))
	assert.True(t, o.IsUint8())

	o = New(uint16(1))
	assert.True(t, o.IsUint16())

	o = New(uint32(1))
	assert.True(t, o.IsUint32())

	o = New(uint64(1))
	assert.True(t, o.IsUint64())

	o = New(float32(1))
	assert.True(t, o.IsFloat32())

	o = New(float64(1))
	assert.True(t, o.IsFloat64())

	o = New(complex64(1))
	assert.True(t, o.IsComplex64())

	o = New(complex128(1))
	assert.True(t, o.IsComplex128())

	o = New(string("1"))
	assert.True(t, o.IsString())

	o = New(func() {})
	assert.True(t, o.IsFunc())

	o = New(uintptr(1))
	assert.True(t, o.IsUintPtr())

}
