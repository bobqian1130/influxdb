// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: array_cursor.gen.go.tmpl

package reads

import (
	"github.com/influxdata/influxdb/tsdb/cursors"
)

const (
	// MaxPointsPerBlock is the maximum number of points in an encoded
	// block in a TSM file. It should match the value in the tsm1
	// package, but we don't want to import it.
	MaxPointsPerBlock = 1000
)

// ********************
// Float Array Cursor

type floatArrayFilterCursor struct {
	cursors.FloatArrayCursor
	cond expression
	m    *singleValue
	res  *cursors.FloatArray
	tmp  *cursors.FloatArray
}

func newFloatFilterArrayCursor(cur cursors.FloatArrayCursor, cond expression) *floatArrayFilterCursor {
	return &floatArrayFilterCursor{
		FloatArrayCursor: cur,
		cond:             cond,
		m:                &singleValue{},
		res:              cursors.NewFloatArrayLen(MaxPointsPerBlock),
		tmp:              &cursors.FloatArray{},
	}
}

func (c *floatArrayFilterCursor) Stats() cursors.CursorStats { return c.FloatArrayCursor.Stats() }

func (c *floatArrayFilterCursor) Next() *cursors.FloatArray {
	pos := 0
	c.res.Timestamps = c.res.Timestamps[:cap(c.res.Timestamps)]
	c.res.Values = c.res.Values[:cap(c.res.Values)]

	var a *cursors.FloatArray

	if c.tmp.Len() > 0 {
		a = c.tmp
		c.tmp.Timestamps = nil
		c.tmp.Values = nil
	} else {
		a = c.FloatArrayCursor.Next()
	}

LOOP:
	for len(a.Timestamps) > 0 {
		for i, v := range a.Values {
			c.m.v = v
			if c.cond.EvalBool(c.m) {
				c.res.Timestamps[pos] = a.Timestamps[i]
				c.res.Values[pos] = v
				pos++
				if pos >= MaxPointsPerBlock {
					c.tmp.Timestamps = a.Timestamps[i+1:]
					c.tmp.Values = a.Values[i+1:]
					break LOOP
				}
			}
		}
		a = c.FloatArrayCursor.Next()
	}

	c.res.Timestamps = c.res.Timestamps[:pos]
	c.res.Values = c.res.Values[:pos]

	return c.res
}

type floatMultiShardArrayCursor struct {
	cursors.FloatArrayCursor
	cursorContext
	filter *floatArrayFilterCursor
}

func (c *floatMultiShardArrayCursor) reset(cur cursors.FloatArrayCursor, itr cursors.CursorIterator, cond expression) {
	if cond != nil {
		if c.filter == nil {
			c.filter = newFloatFilterArrayCursor(cur, cond)
		}
		cur = c.filter
	}

	c.FloatArrayCursor = cur
	c.itr = itr
	c.err = nil
}

func (c *floatMultiShardArrayCursor) Err() error { return c.err }

func (c *floatMultiShardArrayCursor) Stats() cursors.CursorStats {
	return c.FloatArrayCursor.Stats()
}

func (c *floatMultiShardArrayCursor) Next() *cursors.FloatArray {
	return c.FloatArrayCursor.Next()
}

type floatArraySumCursor struct {
	cursors.FloatArrayCursor
	ts  [1]int64
	vs  [1]float64
	res *cursors.FloatArray
}

func newFloatArraySumCursor(cur cursors.FloatArrayCursor) *floatArraySumCursor {
	return &floatArraySumCursor{
		FloatArrayCursor: cur,
		res:              &cursors.FloatArray{},
	}
}

func (c floatArraySumCursor) Stats() cursors.CursorStats { return c.FloatArrayCursor.Stats() }

func (c floatArraySumCursor) Next() *cursors.FloatArray {
	a := c.FloatArrayCursor.Next()
	if len(a.Timestamps) == 0 {
		return a
	}

	ts := a.Timestamps[0]
	var acc float64

	for {
		for _, v := range a.Values {
			acc += v
		}
		a = c.FloatArrayCursor.Next()
		if len(a.Timestamps) == 0 {
			c.ts[0] = ts
			c.vs[0] = acc
			c.res.Timestamps = c.ts[:]
			c.res.Values = c.vs[:]
			return c.res
		}
	}
}

type integerFloatCountArrayCursor struct {
	cursors.FloatArrayCursor
}

func (c *integerFloatCountArrayCursor) Stats() cursors.CursorStats {
	return c.FloatArrayCursor.Stats()
}

func (c *integerFloatCountArrayCursor) Next() *cursors.IntegerArray {
	a := c.FloatArrayCursor.Next()
	if len(a.Timestamps) == 0 {
		return &cursors.IntegerArray{}
	}

	ts := a.Timestamps[0]
	var acc int64
	for {
		acc += int64(len(a.Timestamps))
		a = c.FloatArrayCursor.Next()
		if len(a.Timestamps) == 0 {
			res := cursors.NewIntegerArrayLen(1)
			res.Timestamps[0] = ts
			res.Values[0] = acc
			return res
		}
	}
}

type floatEmptyArrayCursor struct {
	res cursors.FloatArray
}

var FloatEmptyArrayCursor cursors.FloatArrayCursor = &floatEmptyArrayCursor{}

func (c *floatEmptyArrayCursor) Err() error                 { return nil }
func (c *floatEmptyArrayCursor) Close()                     {}
func (c *floatEmptyArrayCursor) Stats() cursors.CursorStats { return cursors.CursorStats{} }
func (c *floatEmptyArrayCursor) Next() *cursors.FloatArray  { return &c.res }

// ********************
// Integer Array Cursor

type integerArrayFilterCursor struct {
	cursors.IntegerArrayCursor
	cond expression
	m    *singleValue
	res  *cursors.IntegerArray
	tmp  *cursors.IntegerArray
}

func newIntegerFilterArrayCursor(cur cursors.IntegerArrayCursor, cond expression) *integerArrayFilterCursor {
	return &integerArrayFilterCursor{
		IntegerArrayCursor: cur,
		cond:               cond,
		m:                  &singleValue{},
		res:                cursors.NewIntegerArrayLen(MaxPointsPerBlock),
		tmp:                &cursors.IntegerArray{},
	}
}

func (c *integerArrayFilterCursor) Stats() cursors.CursorStats { return c.IntegerArrayCursor.Stats() }

func (c *integerArrayFilterCursor) Next() *cursors.IntegerArray {
	pos := 0
	c.res.Timestamps = c.res.Timestamps[:cap(c.res.Timestamps)]
	c.res.Values = c.res.Values[:cap(c.res.Values)]

	var a *cursors.IntegerArray

	if c.tmp.Len() > 0 {
		a = c.tmp
		c.tmp.Timestamps = nil
		c.tmp.Values = nil
	} else {
		a = c.IntegerArrayCursor.Next()
	}

LOOP:
	for len(a.Timestamps) > 0 {
		for i, v := range a.Values {
			c.m.v = v
			if c.cond.EvalBool(c.m) {
				c.res.Timestamps[pos] = a.Timestamps[i]
				c.res.Values[pos] = v
				pos++
				if pos >= MaxPointsPerBlock {
					c.tmp.Timestamps = a.Timestamps[i+1:]
					c.tmp.Values = a.Values[i+1:]
					break LOOP
				}
			}
		}
		a = c.IntegerArrayCursor.Next()
	}

	c.res.Timestamps = c.res.Timestamps[:pos]
	c.res.Values = c.res.Values[:pos]

	return c.res
}

type integerMultiShardArrayCursor struct {
	cursors.IntegerArrayCursor
	cursorContext
	filter *integerArrayFilterCursor
}

func (c *integerMultiShardArrayCursor) reset(cur cursors.IntegerArrayCursor, itr cursors.CursorIterator, cond expression) {
	if cond != nil {
		if c.filter == nil {
			c.filter = newIntegerFilterArrayCursor(cur, cond)
		}
		cur = c.filter
	}

	c.IntegerArrayCursor = cur
	c.itr = itr
	c.err = nil
}

func (c *integerMultiShardArrayCursor) Err() error { return c.err }

func (c *integerMultiShardArrayCursor) Stats() cursors.CursorStats {
	return c.IntegerArrayCursor.Stats()
}

func (c *integerMultiShardArrayCursor) Next() *cursors.IntegerArray {
	return c.IntegerArrayCursor.Next()
}

type integerArraySumCursor struct {
	cursors.IntegerArrayCursor
	ts  [1]int64
	vs  [1]int64
	res *cursors.IntegerArray
}

func newIntegerArraySumCursor(cur cursors.IntegerArrayCursor) *integerArraySumCursor {
	return &integerArraySumCursor{
		IntegerArrayCursor: cur,
		res:                &cursors.IntegerArray{},
	}
}

func (c integerArraySumCursor) Stats() cursors.CursorStats { return c.IntegerArrayCursor.Stats() }

func (c integerArraySumCursor) Next() *cursors.IntegerArray {
	a := c.IntegerArrayCursor.Next()
	if len(a.Timestamps) == 0 {
		return a
	}

	ts := a.Timestamps[0]
	var acc int64

	for {
		for _, v := range a.Values {
			acc += v
		}
		a = c.IntegerArrayCursor.Next()
		if len(a.Timestamps) == 0 {
			c.ts[0] = ts
			c.vs[0] = acc
			c.res.Timestamps = c.ts[:]
			c.res.Values = c.vs[:]
			return c.res
		}
	}
}

type integerIntegerCountArrayCursor struct {
	cursors.IntegerArrayCursor
}

func (c *integerIntegerCountArrayCursor) Stats() cursors.CursorStats {
	return c.IntegerArrayCursor.Stats()
}

func (c *integerIntegerCountArrayCursor) Next() *cursors.IntegerArray {
	a := c.IntegerArrayCursor.Next()
	if len(a.Timestamps) == 0 {
		return &cursors.IntegerArray{}
	}

	ts := a.Timestamps[0]
	var acc int64
	for {
		acc += int64(len(a.Timestamps))
		a = c.IntegerArrayCursor.Next()
		if len(a.Timestamps) == 0 {
			res := cursors.NewIntegerArrayLen(1)
			res.Timestamps[0] = ts
			res.Values[0] = acc
			return res
		}
	}
}

type integerEmptyArrayCursor struct {
	res cursors.IntegerArray
}

var IntegerEmptyArrayCursor cursors.IntegerArrayCursor = &integerEmptyArrayCursor{}

func (c *integerEmptyArrayCursor) Err() error                  { return nil }
func (c *integerEmptyArrayCursor) Close()                      {}
func (c *integerEmptyArrayCursor) Stats() cursors.CursorStats  { return cursors.CursorStats{} }
func (c *integerEmptyArrayCursor) Next() *cursors.IntegerArray { return &c.res }

// ********************
// Unsigned Array Cursor

type unsignedArrayFilterCursor struct {
	cursors.UnsignedArrayCursor
	cond expression
	m    *singleValue
	res  *cursors.UnsignedArray
	tmp  *cursors.UnsignedArray
}

func newUnsignedFilterArrayCursor(cur cursors.UnsignedArrayCursor, cond expression) *unsignedArrayFilterCursor {
	return &unsignedArrayFilterCursor{
		UnsignedArrayCursor: cur,
		cond:                cond,
		m:                   &singleValue{},
		res:                 cursors.NewUnsignedArrayLen(MaxPointsPerBlock),
		tmp:                 &cursors.UnsignedArray{},
	}
}

func (c *unsignedArrayFilterCursor) Stats() cursors.CursorStats { return c.UnsignedArrayCursor.Stats() }

func (c *unsignedArrayFilterCursor) Next() *cursors.UnsignedArray {
	pos := 0
	c.res.Timestamps = c.res.Timestamps[:cap(c.res.Timestamps)]
	c.res.Values = c.res.Values[:cap(c.res.Values)]

	var a *cursors.UnsignedArray

	if c.tmp.Len() > 0 {
		a = c.tmp
		c.tmp.Timestamps = nil
		c.tmp.Values = nil
	} else {
		a = c.UnsignedArrayCursor.Next()
	}

LOOP:
	for len(a.Timestamps) > 0 {
		for i, v := range a.Values {
			c.m.v = v
			if c.cond.EvalBool(c.m) {
				c.res.Timestamps[pos] = a.Timestamps[i]
				c.res.Values[pos] = v
				pos++
				if pos >= MaxPointsPerBlock {
					c.tmp.Timestamps = a.Timestamps[i+1:]
					c.tmp.Values = a.Values[i+1:]
					break LOOP
				}
			}
		}
		a = c.UnsignedArrayCursor.Next()
	}

	c.res.Timestamps = c.res.Timestamps[:pos]
	c.res.Values = c.res.Values[:pos]

	return c.res
}

type unsignedMultiShardArrayCursor struct {
	cursors.UnsignedArrayCursor
	cursorContext
	filter *unsignedArrayFilterCursor
}

func (c *unsignedMultiShardArrayCursor) reset(cur cursors.UnsignedArrayCursor, itr cursors.CursorIterator, cond expression) {
	if cond != nil {
		if c.filter == nil {
			c.filter = newUnsignedFilterArrayCursor(cur, cond)
		}
		cur = c.filter
	}

	c.UnsignedArrayCursor = cur
	c.itr = itr
	c.err = nil
}

func (c *unsignedMultiShardArrayCursor) Err() error { return c.err }

func (c *unsignedMultiShardArrayCursor) Stats() cursors.CursorStats {
	return c.UnsignedArrayCursor.Stats()
}

func (c *unsignedMultiShardArrayCursor) Next() *cursors.UnsignedArray {
	return c.UnsignedArrayCursor.Next()
}

type unsignedArraySumCursor struct {
	cursors.UnsignedArrayCursor
	ts  [1]int64
	vs  [1]uint64
	res *cursors.UnsignedArray
}

func newUnsignedArraySumCursor(cur cursors.UnsignedArrayCursor) *unsignedArraySumCursor {
	return &unsignedArraySumCursor{
		UnsignedArrayCursor: cur,
		res:                 &cursors.UnsignedArray{},
	}
}

func (c unsignedArraySumCursor) Stats() cursors.CursorStats { return c.UnsignedArrayCursor.Stats() }

func (c unsignedArraySumCursor) Next() *cursors.UnsignedArray {
	a := c.UnsignedArrayCursor.Next()
	if len(a.Timestamps) == 0 {
		return a
	}

	ts := a.Timestamps[0]
	var acc uint64

	for {
		for _, v := range a.Values {
			acc += v
		}
		a = c.UnsignedArrayCursor.Next()
		if len(a.Timestamps) == 0 {
			c.ts[0] = ts
			c.vs[0] = acc
			c.res.Timestamps = c.ts[:]
			c.res.Values = c.vs[:]
			return c.res
		}
	}
}

type integerUnsignedCountArrayCursor struct {
	cursors.UnsignedArrayCursor
}

func (c *integerUnsignedCountArrayCursor) Stats() cursors.CursorStats {
	return c.UnsignedArrayCursor.Stats()
}

func (c *integerUnsignedCountArrayCursor) Next() *cursors.IntegerArray {
	a := c.UnsignedArrayCursor.Next()
	if len(a.Timestamps) == 0 {
		return &cursors.IntegerArray{}
	}

	ts := a.Timestamps[0]
	var acc int64
	for {
		acc += int64(len(a.Timestamps))
		a = c.UnsignedArrayCursor.Next()
		if len(a.Timestamps) == 0 {
			res := cursors.NewIntegerArrayLen(1)
			res.Timestamps[0] = ts
			res.Values[0] = acc
			return res
		}
	}
}

type unsignedEmptyArrayCursor struct {
	res cursors.UnsignedArray
}

var UnsignedEmptyArrayCursor cursors.UnsignedArrayCursor = &unsignedEmptyArrayCursor{}

func (c *unsignedEmptyArrayCursor) Err() error                   { return nil }
func (c *unsignedEmptyArrayCursor) Close()                       {}
func (c *unsignedEmptyArrayCursor) Stats() cursors.CursorStats   { return cursors.CursorStats{} }
func (c *unsignedEmptyArrayCursor) Next() *cursors.UnsignedArray { return &c.res }

// ********************
// String Array Cursor

type stringArrayFilterCursor struct {
	cursors.StringArrayCursor
	cond expression
	m    *singleValue
	res  *cursors.StringArray
	tmp  *cursors.StringArray
}

func newStringFilterArrayCursor(cur cursors.StringArrayCursor, cond expression) *stringArrayFilterCursor {
	return &stringArrayFilterCursor{
		StringArrayCursor: cur,
		cond:              cond,
		m:                 &singleValue{},
		res:               cursors.NewStringArrayLen(MaxPointsPerBlock),
		tmp:               &cursors.StringArray{},
	}
}

func (c *stringArrayFilterCursor) Stats() cursors.CursorStats { return c.StringArrayCursor.Stats() }

func (c *stringArrayFilterCursor) Next() *cursors.StringArray {
	pos := 0
	c.res.Timestamps = c.res.Timestamps[:cap(c.res.Timestamps)]
	c.res.Values = c.res.Values[:cap(c.res.Values)]

	var a *cursors.StringArray

	if c.tmp.Len() > 0 {
		a = c.tmp
		c.tmp.Timestamps = nil
		c.tmp.Values = nil
	} else {
		a = c.StringArrayCursor.Next()
	}

LOOP:
	for len(a.Timestamps) > 0 {
		for i, v := range a.Values {
			c.m.v = v
			if c.cond.EvalBool(c.m) {
				c.res.Timestamps[pos] = a.Timestamps[i]
				c.res.Values[pos] = v
				pos++
				if pos >= MaxPointsPerBlock {
					c.tmp.Timestamps = a.Timestamps[i+1:]
					c.tmp.Values = a.Values[i+1:]
					break LOOP
				}
			}
		}
		a = c.StringArrayCursor.Next()
	}

	c.res.Timestamps = c.res.Timestamps[:pos]
	c.res.Values = c.res.Values[:pos]

	return c.res
}

type stringMultiShardArrayCursor struct {
	cursors.StringArrayCursor
	cursorContext
	filter *stringArrayFilterCursor
}

func (c *stringMultiShardArrayCursor) reset(cur cursors.StringArrayCursor, itr cursors.CursorIterator, cond expression) {
	if cond != nil {
		if c.filter == nil {
			c.filter = newStringFilterArrayCursor(cur, cond)
		}
		cur = c.filter
	}

	c.StringArrayCursor = cur
	c.itr = itr
	c.err = nil
}

func (c *stringMultiShardArrayCursor) Err() error { return c.err }

func (c *stringMultiShardArrayCursor) Stats() cursors.CursorStats {
	return c.StringArrayCursor.Stats()
}

func (c *stringMultiShardArrayCursor) Next() *cursors.StringArray {
	return c.StringArrayCursor.Next()
}

type integerStringCountArrayCursor struct {
	cursors.StringArrayCursor
}

func (c *integerStringCountArrayCursor) Stats() cursors.CursorStats {
	return c.StringArrayCursor.Stats()
}

func (c *integerStringCountArrayCursor) Next() *cursors.IntegerArray {
	a := c.StringArrayCursor.Next()
	if len(a.Timestamps) == 0 {
		return &cursors.IntegerArray{}
	}

	ts := a.Timestamps[0]
	var acc int64
	for {
		acc += int64(len(a.Timestamps))
		a = c.StringArrayCursor.Next()
		if len(a.Timestamps) == 0 {
			res := cursors.NewIntegerArrayLen(1)
			res.Timestamps[0] = ts
			res.Values[0] = acc
			return res
		}
	}
}

type stringEmptyArrayCursor struct {
	res cursors.StringArray
}

var StringEmptyArrayCursor cursors.StringArrayCursor = &stringEmptyArrayCursor{}

func (c *stringEmptyArrayCursor) Err() error                 { return nil }
func (c *stringEmptyArrayCursor) Close()                     {}
func (c *stringEmptyArrayCursor) Stats() cursors.CursorStats { return cursors.CursorStats{} }
func (c *stringEmptyArrayCursor) Next() *cursors.StringArray { return &c.res }

// ********************
// Boolean Array Cursor

type booleanArrayFilterCursor struct {
	cursors.BooleanArrayCursor
	cond expression
	m    *singleValue
	res  *cursors.BooleanArray
	tmp  *cursors.BooleanArray
}

func newBooleanFilterArrayCursor(cur cursors.BooleanArrayCursor, cond expression) *booleanArrayFilterCursor {
	return &booleanArrayFilterCursor{
		BooleanArrayCursor: cur,
		cond:               cond,
		m:                  &singleValue{},
		res:                cursors.NewBooleanArrayLen(MaxPointsPerBlock),
		tmp:                &cursors.BooleanArray{},
	}
}

func (c *booleanArrayFilterCursor) Stats() cursors.CursorStats { return c.BooleanArrayCursor.Stats() }

func (c *booleanArrayFilterCursor) Next() *cursors.BooleanArray {
	pos := 0
	c.res.Timestamps = c.res.Timestamps[:cap(c.res.Timestamps)]
	c.res.Values = c.res.Values[:cap(c.res.Values)]

	var a *cursors.BooleanArray

	if c.tmp.Len() > 0 {
		a = c.tmp
		c.tmp.Timestamps = nil
		c.tmp.Values = nil
	} else {
		a = c.BooleanArrayCursor.Next()
	}

LOOP:
	for len(a.Timestamps) > 0 {
		for i, v := range a.Values {
			c.m.v = v
			if c.cond.EvalBool(c.m) {
				c.res.Timestamps[pos] = a.Timestamps[i]
				c.res.Values[pos] = v
				pos++
				if pos >= MaxPointsPerBlock {
					c.tmp.Timestamps = a.Timestamps[i+1:]
					c.tmp.Values = a.Values[i+1:]
					break LOOP
				}
			}
		}
		a = c.BooleanArrayCursor.Next()
	}

	c.res.Timestamps = c.res.Timestamps[:pos]
	c.res.Values = c.res.Values[:pos]

	return c.res
}

type booleanMultiShardArrayCursor struct {
	cursors.BooleanArrayCursor
	cursorContext
	filter *booleanArrayFilterCursor
}

func (c *booleanMultiShardArrayCursor) reset(cur cursors.BooleanArrayCursor, itr cursors.CursorIterator, cond expression) {
	if cond != nil {
		if c.filter == nil {
			c.filter = newBooleanFilterArrayCursor(cur, cond)
		}
		cur = c.filter
	}

	c.BooleanArrayCursor = cur
	c.itr = itr
	c.err = nil
}

func (c *booleanMultiShardArrayCursor) Err() error { return c.err }

func (c *booleanMultiShardArrayCursor) Stats() cursors.CursorStats {
	return c.BooleanArrayCursor.Stats()
}

func (c *booleanMultiShardArrayCursor) Next() *cursors.BooleanArray {
	return c.BooleanArrayCursor.Next()
}

type integerBooleanCountArrayCursor struct {
	cursors.BooleanArrayCursor
}

func (c *integerBooleanCountArrayCursor) Stats() cursors.CursorStats {
	return c.BooleanArrayCursor.Stats()
}

func (c *integerBooleanCountArrayCursor) Next() *cursors.IntegerArray {
	a := c.BooleanArrayCursor.Next()
	if len(a.Timestamps) == 0 {
		return &cursors.IntegerArray{}
	}

	ts := a.Timestamps[0]
	var acc int64
	for {
		acc += int64(len(a.Timestamps))
		a = c.BooleanArrayCursor.Next()
		if len(a.Timestamps) == 0 {
			res := cursors.NewIntegerArrayLen(1)
			res.Timestamps[0] = ts
			res.Values[0] = acc
			return res
		}
	}
}

type booleanEmptyArrayCursor struct {
	res cursors.BooleanArray
}

var BooleanEmptyArrayCursor cursors.BooleanArrayCursor = &booleanEmptyArrayCursor{}

func (c *booleanEmptyArrayCursor) Err() error                  { return nil }
func (c *booleanEmptyArrayCursor) Close()                      {}
func (c *booleanEmptyArrayCursor) Stats() cursors.CursorStats  { return cursors.CursorStats{} }
func (c *booleanEmptyArrayCursor) Next() *cursors.BooleanArray { return &c.res }
