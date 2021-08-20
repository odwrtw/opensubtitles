package opensubtitles

import (
	"bytes"
	"io"
	"testing"
)

// Reference document:
// https://trac.opensubtitles.org/projects/opensubtitles/wiki/HashSourceCodes

// Empty buffer, all the bytes have a zero value, the hash is then equal to the
// size of the file
func generateEmtpyData() io.ReadSeeker {
	data := make([]byte, hashBlockSize)
	for i := 0; i < len(data); i++ {
		data[i] = 0
	}

	return bytes.NewReader(data)
}

func generateTooSmallData() io.ReadSeeker {
	return bytes.NewReader([]byte{0, 0, 0, 0})
}

func TestHash(t *testing.T) {
	tt := []struct {
		name        string
		f           func() io.ReadSeeker
		expected    uint64
		expectedErr error
	}{
		{
			name:     "valid zeroed buffer with the minimal valid size",
			f:        generateEmtpyData,
			expected: hashBlockSize,
		},
		{
			name:        "buffer should be too small",
			f:           generateTooSmallData,
			expectedErr: ErrFileTooSmall,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := tc.f()
			got, err := Hash(r)
			if err != tc.expectedErr {
				t.Fatalf("expected err: %s, got %s", tc.expectedErr, err)
			}

			if got != tc.expected {
				t.Fatalf("invalid hash: expected %d, got %d", tc.expected, got)
			}
		})
	}
}

func TestHashString(t *testing.T) {
	tt := []struct {
		name     string
		hash     uint64
		expected string
	}{
		{
			name:     "non padded hash",
			hash:     10242414353417707026,
			expected: "8e245d9679d31e12",
		},
		{
			name:     "padded hash",
			hash:     72597339223246697,
			expected: "0101eae5380a4769",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := HashString(tc.hash)
			if got != tc.expected {
				t.Fatalf("invalid hash string: expected %s, got %s", tc.expected, got)
			}
		})
	}
}
