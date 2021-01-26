package gmp4

import (
	"os"
)

type localVideo struct {
	file *os.File
	mvhdBox
}
// NewLocal return a localVideo
func NewLocal(fileName string) (*localVideo, error) {
	f, err := os.Open(fileName)
	return &localVideo{file: f}, err
}
// By reading a part of file,Find & parse mvhd box to get time
func (l *localVideo) collectData() {
	// Read file apart
	buf := make([]byte, 128)
	i := 0
	for {
		_, err := l.file.ReadAt(buf, int64(i))
		if err != nil {
			panic(err)
		}
		switch n, cas := check(buf); cas {
		case "mvhd":
			l.metadataIndex = n
			l.metadata = buf
			return
		case "mdat":
			i = int(byteSliceToInt(buf[n-5 : n-1]))
		default:
			i += 128
		}
	}
	defer l.file.Close()
}
func (l *localVideo)getMvhdBox()mvhdBox{
	return l.mvhdBox
}