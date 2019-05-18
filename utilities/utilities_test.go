package utilities

import (
	"bytes"
	"testing"
)

type singleToMulti struct {
	in   []byte
	want [][]byte
}

type singleToSingle struct {
	in   []byte
	want []byte
}

type multitoSingle struct {
	in   [][]byte
	want []byte
}

func TestTranspose(t *testing.T) {
	tests := []singleToMulti{
		{
			in:   []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			want: [][]byte{{0, 4, 8, 12}, {1, 5, 9, 13}, {2, 6, 10, 14}, {3, 7, 11, 15}}},
	}
	for _, test := range tests {
		got := Transpose(test.in, 4)
		for i := range got {
			if bytes.Compare(got[i], test.want[i]) != 0 {
				t.Fatalf("Transpose: Not a match with %v, expected %v got %v", test.in, test.want, got)
			}
		}
	}
}

func TestGenerateEqKey(t *testing.T) {
	tests := []singleToSingle{
		{
			in:   []byte{1, 2, 3},
			want: []byte{1, 2, 3, 1, 2, 3, 1, 2}},
	}
	for _, test := range tests {
		got, err := GenerateEqKey(test.in, 8)
		if err != nil {
			t.Fatalf("GenerateEqKey error: %v", err)
		}
		if bytes.Compare(got, test.want) != 0 {
			t.Fatalf("GenerateEqKey: Not a match with %v, expected %v got %v", test.in, test.want, got)
		}
	}
}

func TestChunkify(t *testing.T) {
	tests := []singleToMulti{
		{
			// Regular array
			in:   []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			want: [][]byte{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}},
		},
		{
			// Empty array
			in:   []byte{},
			want: [][]byte{},
		},
	}
	for _, test := range tests {
		got, err := Chunkify(test.in, 4)
		if err != nil {
			t.Fatalf("Chunkify error: %v", err)
		}
		for i := range got {
			if bytes.Compare(got[i], test.want[i]) != 0 {
				t.Fatalf("Chunkify: Not a match with %v, expected %v got %v", test.in, test.want, got)
			}
		}
	}

}

func TestFindKeySize(t *testing.T) {
	// TODO generate test
}

func TestDetectECB(t *testing.T) {
	// TODO generate test
}

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		in1  []byte
		in2  []byte
		want uint64
	}{{
		// single bit difference
		in1:  []byte{0x01},
		in2:  []byte{0x00},
		want: 1,
	},
		{
			// small difference in two byte array
			in1:  []byte{0x01, 0x01},
			in2:  []byte{0x02, 0x02},
			want: 4,
		},
		{
			// all zero array
			in1:  []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			in2:  []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			want: 0,
		},
		{
			// empty array
			in1:  []byte{},
			in2:  []byte{},
			want: 0,
		}}
	for _, test := range tests {
		got, err := HammingDistance(test.in1, test.in2)
		if err != nil {
			t.Fatalf("hammingDistance error: %v", err)
		}
		if got != test.want {
			t.Fatalf("HammingDistance: Not a match with %v & %v, expected %v got %v", test.in1, test.in2, test.want, got)
		}
	}
}

func TestHamUnsafe(t *testing.T) {
	// TODO generate test
}

func BenchmarkChunkify(b *testing.B) {
	
}
