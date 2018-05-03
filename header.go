package ledorito

import "github.com/superp00t/etc"

type DATHeader struct {
	BlankMagic                   uint32
	TOC_Pointer                  uint32
	TOC_EntryCount               uint32
	UnkA, UnkB, UnkC, UnkD, UnkE uint32
}

func DecodeDATHeader(e *etc.Buffer) DATHeader {
	d := DATHeader{}
	d.BlankMagic = e.ReadUint32()
	d.TOC_Pointer = e.ReadUint32()
	d.TOC_EntryCount = e.ReadUint32()
	d.UnkA = e.ReadUint32()
	d.UnkB = e.ReadUint32()
	d.UnkC = e.ReadUint32()
	d.UnkD = e.ReadUint32()
	d.UnkE = e.ReadUint32()

	return d
}
