// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: iterator.gen.go.tmpl

package arrowutil

import "github.com/influxdata/flux/array"

type IntIterator struct {
	Values []*array.Int
	i      int
	init   bool
}

func IterateInts(arrs []array.Array) IntIterator {
	if len(arrs) == 0 {
		return IntIterator{}
	}
	values := make([]*array.Int, 0, len(arrs))
	for _, arr := range arrs {
		values = append(values, arr.(*array.Int))
	}
	return IntIterator{Values: values}
}

// Value returns the current value in the iterator.
func (i *IntIterator) Value() int64 {
	vs := i.Values[0]
	return vs.Value(i.i)
}

// IsValid returns if the current value is valid.
func (i *IntIterator) IsValid() bool {
	vs := i.Values[0]
	return vs.IsValid(i.i)
}

// IsNull returns if the current value is null.
func (i *IntIterator) IsNull() bool {
	vs := i.Values[0]
	return vs.IsNull(i.i)
}

// Next will move to the next value. It will return false
// if there are no more values to be read. This will
// initialize the iterator if this is the first time it
// is called and return true if there is at least one element.
func (i *IntIterator) Next() bool {
	if !i.init {
		i.init = true
		return i.peek()
	}
	i.i++
	return i.peek()
}

// IsEmpty returns true if the iterator has no values to read.
func (i *IntIterator) IsEmpty() bool {
	return i.peek()
}

// peek will return whether another value is available.
// It will iterate through the iterators until it finds a valid one.
func (i *IntIterator) peek() bool {
	for len(i.Values) > 0 {
		if i.i < i.Values[0].Len() {
			return true
		}
		i.i = 0
		i.Values = i.Values[1:]
	}
	return false
}

type UintIterator struct {
	Values []*array.Uint
	i      int
	init   bool
}

func IterateUints(arrs []array.Array) UintIterator {
	if len(arrs) == 0 {
		return UintIterator{}
	}
	values := make([]*array.Uint, 0, len(arrs))
	for _, arr := range arrs {
		values = append(values, arr.(*array.Uint))
	}
	return UintIterator{Values: values}
}

// Value returns the current value in the iterator.
func (i *UintIterator) Value() uint64 {
	vs := i.Values[0]
	return vs.Value(i.i)
}

// IsValid returns if the current value is valid.
func (i *UintIterator) IsValid() bool {
	vs := i.Values[0]
	return vs.IsValid(i.i)
}

// IsNull returns if the current value is null.
func (i *UintIterator) IsNull() bool {
	vs := i.Values[0]
	return vs.IsNull(i.i)
}

// Next will move to the next value. It will return false
// if there are no more values to be read. This will
// initialize the iterator if this is the first time it
// is called and return true if there is at least one element.
func (i *UintIterator) Next() bool {
	if !i.init {
		i.init = true
		return i.peek()
	}
	i.i++
	return i.peek()
}

// IsEmpty returns true if the iterator has no values to read.
func (i *UintIterator) IsEmpty() bool {
	return i.peek()
}

// peek will return whether another value is available.
// It will iterate through the iterators until it finds a valid one.
func (i *UintIterator) peek() bool {
	for len(i.Values) > 0 {
		if i.i < i.Values[0].Len() {
			return true
		}
		i.i = 0
		i.Values = i.Values[1:]
	}
	return false
}

type FloatIterator struct {
	Values []*array.Float
	i      int
	init   bool
}

func IterateFloats(arrs []array.Array) FloatIterator {
	if len(arrs) == 0 {
		return FloatIterator{}
	}
	values := make([]*array.Float, 0, len(arrs))
	for _, arr := range arrs {
		values = append(values, arr.(*array.Float))
	}
	return FloatIterator{Values: values}
}

// Value returns the current value in the iterator.
func (i *FloatIterator) Value() float64 {
	vs := i.Values[0]
	return vs.Value(i.i)
}

// IsValid returns if the current value is valid.
func (i *FloatIterator) IsValid() bool {
	vs := i.Values[0]
	return vs.IsValid(i.i)
}

// IsNull returns if the current value is null.
func (i *FloatIterator) IsNull() bool {
	vs := i.Values[0]
	return vs.IsNull(i.i)
}

// Next will move to the next value. It will return false
// if there are no more values to be read. This will
// initialize the iterator if this is the first time it
// is called and return true if there is at least one element.
func (i *FloatIterator) Next() bool {
	if !i.init {
		i.init = true
		return i.peek()
	}
	i.i++
	return i.peek()
}

// IsEmpty returns true if the iterator has no values to read.
func (i *FloatIterator) IsEmpty() bool {
	return i.peek()
}

// peek will return whether another value is available.
// It will iterate through the iterators until it finds a valid one.
func (i *FloatIterator) peek() bool {
	for len(i.Values) > 0 {
		if i.i < i.Values[0].Len() {
			return true
		}
		i.i = 0
		i.Values = i.Values[1:]
	}
	return false
}

type BooleanIterator struct {
	Values []*array.Boolean
	i      int
	init   bool
}

func IterateBooleans(arrs []array.Array) BooleanIterator {
	if len(arrs) == 0 {
		return BooleanIterator{}
	}
	values := make([]*array.Boolean, 0, len(arrs))
	for _, arr := range arrs {
		values = append(values, arr.(*array.Boolean))
	}
	return BooleanIterator{Values: values}
}

// Value returns the current value in the iterator.
func (i *BooleanIterator) Value() bool {
	vs := i.Values[0]
	return vs.Value(i.i)
}

// IsValid returns if the current value is valid.
func (i *BooleanIterator) IsValid() bool {
	vs := i.Values[0]
	return vs.IsValid(i.i)
}

// IsNull returns if the current value is null.
func (i *BooleanIterator) IsNull() bool {
	vs := i.Values[0]
	return vs.IsNull(i.i)
}

// Next will move to the next value. It will return false
// if there are no more values to be read. This will
// initialize the iterator if this is the first time it
// is called and return true if there is at least one element.
func (i *BooleanIterator) Next() bool {
	if !i.init {
		i.init = true
		return i.peek()
	}
	i.i++
	return i.peek()
}

// IsEmpty returns true if the iterator has no values to read.
func (i *BooleanIterator) IsEmpty() bool {
	return i.peek()
}

// peek will return whether another value is available.
// It will iterate through the iterators until it finds a valid one.
func (i *BooleanIterator) peek() bool {
	for len(i.Values) > 0 {
		if i.i < i.Values[0].Len() {
			return true
		}
		i.i = 0
		i.Values = i.Values[1:]
	}
	return false
}

type StringIterator struct {
	Values []*array.String
	i      int
	init   bool
}

func IterateStrings(arrs []array.Array) StringIterator {
	if len(arrs) == 0 {
		return StringIterator{}
	}
	values := make([]*array.String, 0, len(arrs))
	for _, arr := range arrs {
		values = append(values, arr.(*array.String))
	}
	return StringIterator{Values: values}
}

// Value returns the current value in the iterator.
func (i *StringIterator) Value() string {
	vs := i.Values[0]
	return vs.Value(i.i)
}

// IsValid returns if the current value is valid.
func (i *StringIterator) IsValid() bool {
	vs := i.Values[0]
	return vs.IsValid(i.i)
}

// IsNull returns if the current value is null.
func (i *StringIterator) IsNull() bool {
	vs := i.Values[0]
	return vs.IsNull(i.i)
}

// Next will move to the next value. It will return false
// if there are no more values to be read. This will
// initialize the iterator if this is the first time it
// is called and return true if there is at least one element.
func (i *StringIterator) Next() bool {
	if !i.init {
		i.init = true
		return i.peek()
	}
	i.i++
	return i.peek()
}

// IsEmpty returns true if the iterator has no values to read.
func (i *StringIterator) IsEmpty() bool {
	return i.peek()
}

// peek will return whether another value is available.
// It will iterate through the iterators until it finds a valid one.
func (i *StringIterator) peek() bool {
	for len(i.Values) > 0 {
		if i.i < i.Values[0].Len() {
			return true
		}
		i.i = 0
		i.Values = i.Values[1:]
	}
	return false
}