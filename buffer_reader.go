package buffer_reader

import (
	"encoding/binary"
)

type BufferReader struct {
	Buf    []byte
	Offset int
}

func NewBufferReader(Buffer []byte) *BufferReader {
	return &BufferReader{Buf: Buffer}
}

func (r *BufferReader) Seek(pos int) {
	if pos >= 0 && pos <= len(r.Buf) {
		r.Offset = pos
	}
}

func (r *BufferReader) Move(diff int) {
	if r.Offset+diff >= 0 && r.Offset+diff <= len(r.Buf) {
		r.Offset += diff
	}
}

func (r *BufferReader) NextUInt32LE() uint32 {
	if r.Offset+4 <= len(r.Buf) {
		val := binary.LittleEndian.Uint32(r.Buf[r.Offset:])
		r.Offset += 4
		return val
	}

	return 0
}

func (r *BufferReader) NextUInt16LE() uint16 {
	if r.Offset+2 <= len(r.Buf) {
		val := binary.LittleEndian.Uint16(r.Buf[r.Offset:])
		r.Offset += 2
		return val
	}

	return 0
}

func (r *BufferReader) NextUInt8() uint8 {
	if r.Offset+1 > len(r.Buf) {
		return 0
	}

	val := uint8(r.Buf[r.Offset])
	r.Offset++
	return val
}

func (r *BufferReader) Tell() int {
	return r.Offset
}
