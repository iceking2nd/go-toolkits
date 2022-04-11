package strings

import "strconv"

type StrExt string

func (f StrExt) Exist() bool {
	return string(f) != string(0x1E)
}

func (f StrExt) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

func (f StrExt) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	return uint64(v), err
}

func (f StrExt) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 0)
	return int(v), err
}

func (f StrExt) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

func (f StrExt) Float64() (float64, error) {
	v, err := strconv.ParseFloat(f.String(), 64)
	return float64(v), err
}

func (f StrExt) MustUint8() uint8 {
	v, _ := f.Uint8()
	return v
}

func (f StrExt) MustUint64() uint64 {
	v, _ := f.Uint64()
	return v
}

func (f StrExt) MustInt() int {
	v, _ := f.Int()
	return v
}

func (f StrExt) MustInt64() int64 {
	v, _ := f.Int64()
	return v
}

func (f StrExt) MustFloat64() float64 {
	v, _ := f.Float64()
	return v
}

func (f StrExt) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}
