package graphql

import (
	"encoding/json"
	"time"
)

type Query struct {
	Data interface{} `graphql:"data"`
}

type GqlBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

func NewBoolStruct(x bool) GqlBool {
	nw := NewBool(x)
	return *nw
}

func NewBool(x bool) *GqlBool { return &GqlBool{Bool: x, Valid: true} }

func (nb GqlBool) MarshalJSON() ([]byte, error) {
	if nb.Valid {
		return json.Marshal(nb.Bool)
	}
	return json.Marshal(nil)
}

func (nb *GqlBool) UnmarshalJSON(data []byte) error {
	var b *bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	if b != nil {
		nb.Valid = true
		nb.Bool = *b
	} else {
		nb.Valid = false
	}
	return nil
}

type GqlFloat64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
}

func NewFloat64Struct(x float64) GqlFloat64 {
	nw := NewFloat64(x)
	return *nw
}

func NewFloat64(x float64) *GqlFloat64 { return &GqlFloat64{Float64: x, Valid: true} }

func (nf GqlFloat64) MarshalJSON() ([]byte, error) {
	if nf.Valid {
		return json.Marshal(nf.Float64)
	}
	return json.Marshal(nil)
}

func (nf *GqlFloat64) UnmarshalJSON(data []byte) error {
	var f *float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	if f != nil {
		nf.Valid = true
		nf.Float64 = *f
	} else {
		nf.Valid = false
	}
	return nil
}

type GqlInt64 struct {
	Int64 int64
	Valid bool // Valid is true if Int64 is not NULL
}

func NewInt64Struct(x int64) GqlInt64 {
	nw := NewInt64(x)
	return *nw
}

func NewInt64(x int64) *GqlInt64 { return &GqlInt64{Int64: x, Valid: true} }

func (ni GqlInt64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return json.Marshal(nil)
}

func (ni *GqlInt64) UnmarshalJSON(data []byte) error {
	var i *int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	if i != nil {
		ni.Valid = true
		ni.Int64 = *i
	} else {
		ni.Valid = false
	}
	return nil
}

type GqlString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

func NewStringStruct(x string) GqlString {
	nw := NewString(x)
	return *nw
}

func NewString(x string) *GqlString { return &GqlString{String: x, Valid: true} }

func (ns GqlString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

func (ns *GqlString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false
	}
	return nil
}

type GqlTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func NewTimeStruct(x time.Time) GqlTime {
	nw := NewTime(x)
	return *nw
}

func NewTime(x time.Time) *GqlTime { return &GqlTime{Time: x, Valid: true} }

func (nt GqlTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time)
	}
	return json.Marshal(nil)
}

func (nt *GqlTime) UnmarshalJSON(data []byte) error {
	var t *time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	if t != nil {
		nt.Valid = true
		nt.Time = *t
	} else {
		nt.Valid = false
	}
	return nil
}

type GqlPoint struct {
	//this does not need an ID as it is a scalar type in GraphQL

	Latitude  *GqlFloat64 `json:"latitude,omitempty"`
	Longitude *GqlFloat64 `json:"longitude,omitempty"`
}

func NewPointStruct(lat, lng float64) GqlPoint {
	nw := NewPoint(lat, lng)
	return *nw
}

func NewPoint(lat, lng float64) *GqlPoint {
	return &GqlPoint{Latitude: NewFloat64(lat), Longitude: NewFloat64(lng)}
}
