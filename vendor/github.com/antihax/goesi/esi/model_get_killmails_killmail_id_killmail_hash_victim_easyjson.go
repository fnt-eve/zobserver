// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson8cf1533DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashVictimList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetKillmailsKillmailIdKillmailHashVictimList, 0, 0)
			} else {
				*out = GetKillmailsKillmailIdKillmailHashVictimList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetKillmailsKillmailIdKillmailHashVictim
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf1533EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashVictimList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashVictimList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8cf1533EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashVictimList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8cf1533EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashVictimList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8cf1533DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashVictimList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8cf1533DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson8cf1533DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashVictim) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "alliance_id":
			out.AllianceId = int32(in.Int32())
		case "character_id":
			out.CharacterId = int32(in.Int32())
		case "corporation_id":
			out.CorporationId = int32(in.Int32())
		case "damage_taken":
			out.DamageTaken = int32(in.Int32())
		case "faction_id":
			out.FactionId = int32(in.Int32())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]GetKillmailsKillmailIdKillmailHashItem, 0, 1)
					} else {
						out.Items = []GetKillmailsKillmailIdKillmailHashItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetKillmailsKillmailIdKillmailHashItem
					easyjson8cf1533DecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "position":
			easyjson8cf1533DecodeGithubComAntihaxGoesiEsi3(in, &out.Position)
		case "ship_type_id":
			out.ShipTypeId = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf1533EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashVictim) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllianceId != 0 {
		const prefix string = ",\"alliance_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AllianceId))
	}
	if in.CharacterId != 0 {
		const prefix string = ",\"character_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CharacterId))
	}
	if in.CorporationId != 0 {
		const prefix string = ",\"corporation_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.CorporationId))
	}
	if in.DamageTaken != 0 {
		const prefix string = ",\"damage_taken\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.DamageTaken))
	}
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	if len(in.Items) != 0 {
		const prefix string = ",\"items\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Items {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson8cf1533EncodeGithubComAntihaxGoesiEsi2(out, v6)
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson8cf1533EncodeGithubComAntihaxGoesiEsi3(out, in.Position)
	}
	if in.ShipTypeId != 0 {
		const prefix string = ",\"ship_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ShipTypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashVictim) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8cf1533EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetKillmailsKillmailIdKillmailHashVictim) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8cf1533EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashVictim) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8cf1533DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetKillmailsKillmailIdKillmailHashVictim) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8cf1533DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson8cf1533DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashPosition) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		case "z":
			out.Z = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf1533EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashPosition) {
	out.RawByte('{')
	first := true
	_ = first
	if in.X != 0 {
		const prefix string = ",\"x\":"
		first = false
		out.RawString(prefix[1:])
		out.Float64(float64(in.X))
	}
	if in.Y != 0 {
		const prefix string = ",\"y\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Y))
	}
	if in.Z != 0 {
		const prefix string = ",\"z\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Z))
	}
	out.RawByte('}')
}
func easyjson8cf1533DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "flag":
			out.Flag = int32(in.Int32())
		case "item_type_id":
			out.ItemTypeId = int32(in.Int32())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]GetKillmailsKillmailIdKillmailHashItemsItem, 0, 2)
					} else {
						out.Items = []GetKillmailsKillmailIdKillmailHashItemsItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v7 GetKillmailsKillmailIdKillmailHashItemsItem
					easyjson8cf1533DecodeGithubComAntihaxGoesiEsi4(in, &v7)
					out.Items = append(out.Items, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "quantity_destroyed":
			out.QuantityDestroyed = int64(in.Int64())
		case "quantity_dropped":
			out.QuantityDropped = int64(in.Int64())
		case "singleton":
			out.Singleton = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf1533EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Flag != 0 {
		const prefix string = ",\"flag\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Flag))
	}
	if in.ItemTypeId != 0 {
		const prefix string = ",\"item_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ItemTypeId))
	}
	if len(in.Items) != 0 {
		const prefix string = ",\"items\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Items {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjson8cf1533EncodeGithubComAntihaxGoesiEsi4(out, v9)
			}
			out.RawByte(']')
		}
	}
	if in.QuantityDestroyed != 0 {
		const prefix string = ",\"quantity_destroyed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.QuantityDestroyed))
	}
	if in.QuantityDropped != 0 {
		const prefix string = ",\"quantity_dropped\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.QuantityDropped))
	}
	if in.Singleton != 0 {
		const prefix string = ",\"singleton\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Singleton))
	}
	out.RawByte('}')
}
func easyjson8cf1533DecodeGithubComAntihaxGoesiEsi4(in *jlexer.Lexer, out *GetKillmailsKillmailIdKillmailHashItemsItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "flag":
			out.Flag = int32(in.Int32())
		case "item_type_id":
			out.ItemTypeId = int32(in.Int32())
		case "quantity_destroyed":
			out.QuantityDestroyed = int64(in.Int64())
		case "quantity_dropped":
			out.QuantityDropped = int64(in.Int64())
		case "singleton":
			out.Singleton = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson8cf1533EncodeGithubComAntihaxGoesiEsi4(out *jwriter.Writer, in GetKillmailsKillmailIdKillmailHashItemsItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Flag != 0 {
		const prefix string = ",\"flag\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Flag))
	}
	if in.ItemTypeId != 0 {
		const prefix string = ",\"item_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ItemTypeId))
	}
	if in.QuantityDestroyed != 0 {
		const prefix string = ",\"quantity_destroyed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.QuantityDestroyed))
	}
	if in.QuantityDropped != 0 {
		const prefix string = ",\"quantity_dropped\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.QuantityDropped))
	}
	if in.Singleton != 0 {
		const prefix string = ",\"singleton\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Singleton))
	}
	out.RawByte('}')
}
