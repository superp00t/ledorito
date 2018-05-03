package ledorito

import (
	"fmt"

	"github.com/pierrec/lz4"
	"github.com/superp00t/etc"
)

type Chunk struct {
	Offset uint32
	Size   uint32
	ZSize  uint32
	Data   []byte
}

func DecodeChunk(e *etc.Buffer) (*Chunk, error) {
	c := &Chunk{}
	c.Offset = uint32(e.Rpos())
	c.Size = e.ReadUint32()
	c.ZSize = e.ReadUint32()

	if c.Size == 0 {
		e.Jump(paddingAlign(e.Rpos(), 2))
		return c, nil
	}

	zbuffer := e.ReadBytes(int(c.ZSize))

	dst := make([]byte, c.Size)
	bytesWritten, err := lz4.UncompressBlock(zbuffer, dst, 0)
	if err != nil {
		return nil, err
	}

	if bytesWritten != int(c.Size) {
		return nil, fmt.Errorf("invalid block read")
	}

	c.Data = dst

	return c, nil
}

func paddingAlign(num, alignTo int64) int64 {
	if (num % alignTo) == 0 {
		return 0
	} else {
		return alignTo - (num % alignTo)
	}
}
