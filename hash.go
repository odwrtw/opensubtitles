package opensubtitles

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
)

const hashBlockSize = 65536

var uint64Size = int(reflect.TypeOf(uint64(0)).Size())

// ErrFileTooSmall is returned if the file is too small to compute a hash
var ErrFileTooSmall = errors.New("opensubtitles: file too small")

// Hash computes the opensubtitle file hash
func Hash(r io.ReadSeeker) (uint64, error) {
	size, err := r.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}

	if size < hashBlockSize {
		return 0, ErrFileTooSmall
	}

	hash := uint64(size)
	for _, offset := range []int64{0, size - hashBlockSize} {
		_, err = r.Seek(offset, io.SeekStart)
		if err != nil {
			return 0, err
		}

		var tmp uint64
		for i := 0; i < hashBlockSize/uint64Size; i++ {
			if err := binary.Read(r, binary.LittleEndian, &tmp); err != nil {
				return 0, err
			}

			hash += tmp
		}
	}

	return hash, nil
}

// HashString returns the padded hexadecimal representation of a hash
func HashString(hash uint64) string {
	return fmt.Sprintf("%016x", hash)
}
