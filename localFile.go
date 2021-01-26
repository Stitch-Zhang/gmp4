package gmp4

import (
	"os"
)

//LocalVideo define a local video
type LocalVideo struct {
	file *os.File
	mvhdBox
}

// NewLocal return a LocalVideo
func NewLocal(fileName string) (*LocalVideo, error) {
	f, err := os.Open(fileName)
	return &LocalVideo{file: f}, err
}

// By reading a part of file,Find & parse mvhd box to get time
func (l *LocalVideo) collectData() {
	defer l.file.Close()
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
}
func (l *LocalVideo) getMvhdBox() mvhdBox {
	return l.mvhdBox
}
