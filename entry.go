package tiff

import (
	"encoding/json"
	"fmt"
)

/*
Entry structure
  For IFD/Entry:
	Each 12-byte IFD entry has the following format:
	Bytes 0-1:  The Tag that identifies the entry.
	Bytes 2-3:  The entry Type.
	Bytes 4-7:  The number of values, Count of the indicated Type.
	Bytes 8-11: The Value Offset, the file offset (in bytes) of the Value
	            for the entry. The Value is expected to begin on a word
	            boundary; the corresponding Value Offset will thus be an
	            even number. This file offset may point anywhere in the
	            file, even after the image data.

  For IFD8/Entry8:
	Each 20-byte IFD entry has the following format:
	Bytes 0-1:   The Tag that identifies the entry.
	Bytes 2-3:   The entry Type.
	Bytes 4-11:  The number of values, Count of the indicated Type.
	Bytes 12-19: The Value Offset, the file offset (in bytes) of the Value
	             for the entry. The Value is expected to begin on a word
	             boundary; the corresponding Value Offset will thus be an
	             even number. This file offset may point anywhere in the
	             file, even after the image data.
*/

// Entry represents a single entry in an IFD in a TIFF file.  This is the mostly
// uninterpreted core 12 byte data structure only.
type Entry interface {
	TagID() uint16
	TypeID() uint16
	Count() uint32
	ValueOffset() [4]byte
}

// entry represents the data structure of an IFD entry.
type entry struct {
	tagID       uint16  // Bytes 0-1
	typeID      uint16  // Bytes 2-3
	count       uint32  // Bytes 4-7
	valueOffset [4]byte // Bytes 8-11
}

func (e *entry) TagID() uint16 {
	return e.tagID
}

func (e *entry) TypeID() uint16 {
	return e.typeID
}

func (e *entry) Count() uint32 {
	return e.count
}

func (e *entry) ValueOffset() [4]byte {
	return e.valueOffset
}

func (e *entry) String() string {
	return fmt.Sprintf("<TagID: %5d, TypeID: %5d, Count: %d, ValueOffset: %v>", e.tagID, e.typeID, e.count, e.valueOffset)
}

func (e *entry) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Tag         uint16  `json:"tagID"`
		Type        uint16  `json:"typeID"`
		Count       uint32  `json:"count"`
		ValueOffset [4]byte `json:"valueOffset"`
	}{
		Tag:         e.tagID,
		Type:        e.typeID,
		Count:       e.count,
		ValueOffset: e.valueOffset,
	}
	return json.Marshal(tmp)
}

func ParseEntry(br BReader) (out Entry, err error) {
	e := new(entry)
	if err = br.BRead(&e.tagID); err != nil {
		return
	}
	if err = br.BRead(&e.typeID); err != nil {
		return
	}
	if err = br.BRead(&e.count); err != nil {
		return
	}
	if err = br.BRead(&e.valueOffset); err != nil {
		return
	}
	return e, nil
}

// Entry8 represents a single entry in an IFD8 in a BigTIFF file.  This is the
// mostly uninterpreted core 20 byte data structure only.
type Entry8 interface {
	TagID() uint16
	TypeID() uint16
	Count() uint64
	ValueOffset() [8]byte
}

// entry8 represents the data structure of an IFD8 entry.
type entry8 struct {
	tagID       uint16  // Bytes 0-1
	typeID      uint16  // Bytes 2-3
	count       uint64  // Bytes 4-11
	valueOffset [8]byte // Bytes 12-19
}

func (e8 *entry8) TagID() uint16 {
	return e8.tagID
}

func (e8 *entry8) TypeID() uint16 {
	return e8.typeID
}

func (e8 *entry8) Count() uint64 {
	return e8.count
}

func (e8 *entry8) ValueOffset() [8]byte {
	return e8.valueOffset
}

func (e8 *entry8) String() string {
	return fmt.Sprintf("<TagID: %5d, TypeID: %5d, Count: %d, ValueOffset: %v>", e8.tagID, e8.typeID, e8.count, e8.valueOffset)
}

func (e8 *entry8) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Tag         uint16  `json:"tagID"`
		Type        uint16  `json:"typeID"`
		Count       uint64  `json:"count"`
		ValueOffset [8]byte `json:"valueOffset"`
	}{
		Tag:         e8.tagID,
		Type:        e8.typeID,
		Count:       e8.count,
		ValueOffset: e8.valueOffset,
	}
	return json.Marshal(tmp)
}

func ParseEntry8(br BReader) (out Entry8, err error) {
	e := new(entry8)
	if err = br.BRead(&e.tagID); err != nil {
		return
	}
	if err = br.BRead(&e.typeID); err != nil {
		return
	}
	if err = br.BRead(&e.count); err != nil {
		return
	}
	if err = br.BRead(&e.valueOffset); err != nil {
		return
	}
	return e, nil
}
