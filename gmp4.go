// This package was written based on ISO/IEC 14496
// Gain video time,only supported mp4 now
package gmp4

import (
	"bytes"
	"encoding/binary"
	"log"
)

type mvhdBox struct {
	metadata        []byte
	metadataIndex   int
	totalTimeSecond uint32
	done            bool
}

type video interface {
	collectData()
	getMvhdBox() mvhdBox
}

// Translate byte to int
func byteSliceToInt(raw []byte) (result uint32) {
	if len(raw) != 4 {
		log.Panicln("gmp4.go:data is invalid")
		return 0
	}
	buf := bytes.NewBuffer(raw)
	err := binary.Read(buf, binary.BigEndian, &result)
	if err != nil {
		log.Panicln("gmp4.go:data read failed!")
		return 0
	}
	return
}

// check find the index of mvhd box
func check(raw []byte) (int, string) {
	for i := 0; i < len(raw); i++ {
		//lmvhd stand for mvhd box
		if string(raw[i:i+5]) == "lmvhd" {
			return i, "mvhd"
		} else if string(raw[i:i+3]) == "dat" {
			return i, "mdat"
		}
	}
	return 0, ""
}

// GetDuration get video time(second)
// by downloading mvhd box of video and parse it
func GetDuration(m video) uint32 {
	m.collectData()
	data := m.getMvhdBox()
	if data.done {
		return data.totalTimeSecond
	}
	m.collectData()
	data.totalTimeSecond = byteSliceToInt(data.metadata[data.metadataIndex+21:data.metadataIndex+25]) / byteSliceToInt(data.metadata[data.metadataIndex+17:data.metadataIndex+21])
	data.done = true
	return data.totalTimeSecond
}
