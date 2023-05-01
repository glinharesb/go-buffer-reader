package buffer_reader

import (
	"encoding/binary"
)

type BufferReader struct {
	buf    []byte
	offset int
}

func NewBufferReader(buffer []byte) *BufferReader {
	return &BufferReader{buf: buffer}
}

func (r *BufferReader) Seek(pos int) {
	if pos >= 0 && pos <= len(r.buf) {
		r.offset = pos
	}
}

func (r *BufferReader) Move(diff int) {
	if r.offset+diff >= 0 && r.offset+diff <= len(r.buf) {
		r.offset += diff
	}
}

func (r *BufferReader) NextUInt32LE() uint32 {
	if r.offset+4 <= len(r.buf) {
		val := binary.LittleEndian.Uint32(r.buf[r.offset:])
		r.offset += 4
		return val
	}
	return 0
}

func (r *BufferReader) NextUInt16LE() uint16 {
	if r.offset+2 <= len(r.buf) {
		val := binary.LittleEndian.Uint16(r.buf[r.offset:])
		r.offset += 2
		return val
	}
	return 0
}

func (r *BufferReader) NextUInt8() uint8 {
	if r.offset+1 > len(r.buf) {
		return 0
	}

	val := uint8(r.buf[r.offset])
	r.offset++
	return val
}

func (r *BufferReader) Tell() int {
	return r.offset
}
