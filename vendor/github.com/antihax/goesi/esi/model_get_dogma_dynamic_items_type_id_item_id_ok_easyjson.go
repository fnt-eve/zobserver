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

func easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetDogmaDynamicItemsTypeIdItemIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetDogmaDynamicItemsTypeIdItemIdOkList, 0, 1)
			} else {
				*out = GetDogmaDynamicItemsTypeIdItemIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetDogmaDynamicItemsTypeIdItemIdOk
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
func easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetDogmaDynamicItemsTypeIdItemIdOkList) {
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
func (v GetDogmaDynamicItemsTypeIdItemIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetDogmaDynamicItemsTypeIdItemIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetDogmaDynamicItemsTypeIdItemIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetDogmaDynamicItemsTypeIdItemIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetDogmaDynamicItemsTypeIdItemIdOk) {
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
		case "created_by":
			out.CreatedBy = int32(in.Int32())
		case "dogma_attributes":
			if in.IsNull() {
				in.Skip()
				out.DogmaAttributes = nil
			} else {
				in.Delim('[')
				if out.DogmaAttributes == nil {
					if !in.IsDelim(']') {
						out.DogmaAttributes = make([]GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute, 0, 8)
					} else {
						out.DogmaAttributes = []GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute{}
					}
				} else {
					out.DogmaAttributes = (out.DogmaAttributes)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute
					easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.DogmaAttributes = append(out.DogmaAttributes, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "dogma_effects":
			if in.IsNull() {
				in.Skip()
				out.DogmaEffects = nil
			} else {
				in.Delim('[')
				if out.DogmaEffects == nil {
					if !in.IsDelim(']') {
						out.DogmaEffects = make([]GetDogmaDynamicItemsTypeIdItemIdDogmaEffect, 0, 8)
					} else {
						out.DogmaEffects = []GetDogmaDynamicItemsTypeIdItemIdDogmaEffect{}
					}
				} else {
					out.DogmaEffects = (out.DogmaEffects)[:0]
				}
				for !in.IsDelim(']') {
					var v5 GetDogmaDynamicItemsTypeIdItemIdDogmaEffect
					easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi3(in, &v5)
					out.DogmaEffects = append(out.DogmaEffects, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "mutator_type_id":
			out.MutatorTypeId = int32(in.Int32())
		case "source_type_id":
			out.SourceTypeId = int32(in.Int32())
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
func easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetDogmaDynamicItemsTypeIdItemIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CreatedBy != 0 {
		const prefix string = ",\"created_by\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.CreatedBy))
	}
	if len(in.DogmaAttributes) != 0 {
		const prefix string = ",\"dogma_attributes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.DogmaAttributes {
				if v6 > 0 {
					out.RawByte(',')
				}
				easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi2(out, v7)
			}
			out.RawByte(']')
		}
	}
	if len(in.DogmaEffects) != 0 {
		const prefix string = ",\"dogma_effects\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.DogmaEffects {
				if v8 > 0 {
					out.RawByte(',')
				}
				easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi3(out, v9)
			}
			out.RawByte(']')
		}
	}
	if in.MutatorTypeId != 0 {
		const prefix string = ",\"mutator_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.MutatorTypeId))
	}
	if in.SourceTypeId != 0 {
		const prefix string = ",\"source_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.SourceTypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetDogmaDynamicItemsTypeIdItemIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetDogmaDynamicItemsTypeIdItemIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetDogmaDynamicItemsTypeIdItemIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetDogmaDynamicItemsTypeIdItemIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi3(in *jlexer.Lexer, out *GetDogmaDynamicItemsTypeIdItemIdDogmaEffect) {
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
		case "effect_id":
			out.EffectId = int32(in.Int32())
		case "is_default":
			out.IsDefault = bool(in.Bool())
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
func easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi3(out *jwriter.Writer, in GetDogmaDynamicItemsTypeIdItemIdDogmaEffect) {
	out.RawByte('{')
	first := true
	_ = first
	if in.EffectId != 0 {
		const prefix string = ",\"effect_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.EffectId))
	}
	if in.IsDefault {
		const prefix string = ",\"is_default\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsDefault))
	}
	out.RawByte('}')
}
func easyjson5a3b9194DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute) {
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
		case "attribute_id":
			out.AttributeId = int32(in.Int32())
		case "value":
			out.Value = float32(in.Float32())
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
func easyjson5a3b9194EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AttributeId != 0 {
		const prefix string = ",\"attribute_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AttributeId))
	}
	if in.Value != 0 {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Value))
	}
	out.RawByte('}')
}