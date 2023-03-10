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

func easyjson13ace497DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdClonesOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdClonesOkList, 0, 0)
			} else {
				*out = GetCharactersCharacterIdClonesOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdClonesOk
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
func easyjson13ace497EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdClonesOkList) {
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
func (v GetCharactersCharacterIdClonesOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson13ace497EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdClonesOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson13ace497EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson13ace497DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson13ace497DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson13ace497DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdClonesOk) {
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
		case "home_location":
			easyjson13ace497DecodeGithubComAntihaxGoesiEsi2(in, &out.HomeLocation)
		case "jump_clones":
			if in.IsNull() {
				in.Skip()
				out.JumpClones = nil
			} else {
				in.Delim('[')
				if out.JumpClones == nil {
					if !in.IsDelim(']') {
						out.JumpClones = make([]GetCharactersCharacterIdClonesJumpClone, 0, 0)
					} else {
						out.JumpClones = []GetCharactersCharacterIdClonesJumpClone{}
					}
				} else {
					out.JumpClones = (out.JumpClones)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdClonesJumpClone
					easyjson13ace497DecodeGithubComAntihaxGoesiEsi3(in, &v4)
					out.JumpClones = append(out.JumpClones, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "last_clone_jump_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastCloneJumpDate).UnmarshalJSON(data))
			}
		case "last_station_change_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastStationChangeDate).UnmarshalJSON(data))
			}
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
func easyjson13ace497EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdClonesOk) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"home_location\":"
		first = false
		out.RawString(prefix[1:])
		easyjson13ace497EncodeGithubComAntihaxGoesiEsi2(out, in.HomeLocation)
	}
	if len(in.JumpClones) != 0 {
		const prefix string = ",\"jump_clones\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.JumpClones {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson13ace497EncodeGithubComAntihaxGoesiEsi3(out, v6)
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"last_clone_jump_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastCloneJumpDate).MarshalJSON())
	}
	if true {
		const prefix string = ",\"last_station_change_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastStationChangeDate).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdClonesOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson13ace497EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdClonesOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson13ace497EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson13ace497DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson13ace497DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson13ace497DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetCharactersCharacterIdClonesJumpClone) {
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
		case "implants":
			if in.IsNull() {
				in.Skip()
				out.Implants = nil
			} else {
				in.Delim('[')
				if out.Implants == nil {
					if !in.IsDelim(']') {
						out.Implants = make([]int32, 0, 16)
					} else {
						out.Implants = []int32{}
					}
				} else {
					out.Implants = (out.Implants)[:0]
				}
				for !in.IsDelim(']') {
					var v7 int32
					v7 = int32(in.Int32())
					out.Implants = append(out.Implants, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "jump_clone_id":
			out.JumpCloneId = int32(in.Int32())
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "location_type":
			out.LocationType = string(in.String())
		case "name":
			out.Name = string(in.String())
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
func easyjson13ace497EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetCharactersCharacterIdClonesJumpClone) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Implants) != 0 {
		const prefix string = ",\"implants\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v8, v9 := range in.Implants {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v9))
			}
			out.RawByte(']')
		}
	}
	if in.JumpCloneId != 0 {
		const prefix string = ",\"jump_clone_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.JumpCloneId))
	}
	if in.LocationId != 0 {
		const prefix string = ",\"location_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LocationId))
	}
	if in.LocationType != "" {
		const prefix string = ",\"location_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LocationType))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}
func easyjson13ace497DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdClonesHomeLocation) {
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
		case "location_id":
			out.LocationId = int64(in.Int64())
		case "location_type":
			out.LocationType = string(in.String())
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
func easyjson13ace497EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdClonesHomeLocation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LocationId != 0 {
		const prefix string = ",\"location_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.LocationId))
	}
	if in.LocationType != "" {
		const prefix string = ",\"location_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LocationType))
	}
	out.RawByte('}')
}
