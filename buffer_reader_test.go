package buffer_reader

import (
	"testing"
)

func TestBufferReader(t *testing.T) {
	buf := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07}
	reader := NewBufferReader(buf)

	val := reader.NextUInt8()
	if val != 0x01 {
		t.Errorf("Expected 0x01, but got 0x%02X", val)
	}

	val16 := reader.NextUInt16LE()
	if val16 != 0x0302 {
		t.Errorf("Expected 0x0302, but got 0x%04X", val16)
	}

	val32 := reader.NextUInt32LE()
	if val32 != 0x07060504 {
		t.Errorf("Expected 0x07060504, but got 0x%08X", val32)
	}

	offset := reader.Tell()
	if offset != 7 {
		t.Errorf("Expected offset 7, but got %d", offset)
	}

	reader.Seek(2)
	offset = reader.Tell()
	if offset != 2 {
		t.Errorf("Expected offset 2 after Seek(2), but got %d", offset)
	}

	reader.Move(2)
	offset = reader.Tell()
	if offset != 4 {
		t.Errorf("Expected offset 4 after Move(2), but got %d", offset)
	}
}
