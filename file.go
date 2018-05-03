package ledorito

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/superp00t/etc"
)

type DATFile struct {
	Header      DATHeader
	TOC_Entries []uint32

	file *etc.Buffer
}

func Extract(path, outpath string) error {
	ctl, err := etc.FileController(path)
	if err != nil {
		return err
	}

	df := new(DATFile)
	df.file = ctl
	df.Header = DecodeDATHeader(df.file)
	fmt.Println(spew.Sdump(df.Header))
	df.file.Seek(int64(df.Header.TOC_Pointer))

	for i := uint32(0); i < df.Header.TOC_EntryCount; i++ {
		ent := df.file.ReadUint32()
		fmt.Println("Entry", ent)
		if ent == 0xFFFFFFFF {
			continue
		}
		df.TOC_Entries = append(df.TOC_Entries, ent)
	}

	df.file.Seek(0)

	for j := 0; j < len(df.TOC_Entries); j++ {
		curEntry := df.TOC_Entries[j]
		nextEntry := uint32(0)
		if (j + 1) < len(df.TOC_Entries) {
			nextEntry = df.TOC_Entries[j+1]
		} else {
			nextEntry = df.Header.TOC_Pointer
		}

		df.file.Seek(int64(curEntry))

		curChunk := Chunk{}
		var subChunks []*Chunk

		for df.file.Rpos() < int64(nextEntry-16) {
			subChunk, err := DecodeChunk(df.file)
			if err != nil {
				return err
			}
			subChunks = append(subChunks, subChunk)
		}

		if len(subChunks) == 0 {
			continue
		}

		curChunk.Offset = subChunks[0].Offset

		for i := 0; i < len(subChunks); i++ {
			if subChunks[i].Data == nil {
				continue
			}

			curChunk.Data = append(curChunk.Data, subChunks[i].Data...)
			curChunk.Size += subChunks[i].Size
			curChunk.ZSize += subChunks[i].ZSize
		}

		sp := filepath.SplitList(outpath)
		outFilePath := filepath.Join(append(sp, fmt.Sprintf("entry_%d.datchunk", j))...)

		err := ioutil.WriteFile(outFilePath, curChunk.Data, 0700)
		if err != nil {
			return err
		}

		fmt.Println("Extracted", outFilePath)
	}

	return nil
}
