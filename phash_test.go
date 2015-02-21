package phash

import (
	"testing"
)

var dctTestData = []struct {
	filename string
	hash     uint64
}{
	{"testdata/gif/gopher.gif", 10631249863007147461},
	{"testdata/jpg/gopher.jpg", 10631249863007147461},
	{"testdata/jpg/gopher-blur.jpg", 10631249863007147461},
	{"testdata/jpg/gopher-rotate.jpg", 10631249863007147461},
	{"testdata/jpg/gopher-flip.jpg", 9364791847115625850},
	{"testdata/jpg/gopher-flop.jpg", 17752165113623953620},
	{"testdata/jpg/gopher-shave.jpg", 11459704921468884741},
	{"testdata/png/gopher.png", 10631249863007147461},
}

var hammingTestData = []struct {
	h1, h2   uint64
	distance int
}{
	{10631249863007147461, 10631249863007147461, 0},
	{10631249863007147461, 9364791847115625850, 32},
	{10631249863007147461, 17752165113623953620, 32},
	{10631249863007147461, 11459704921468884741, 16},
}

func TestImageHashDCT(t *testing.T) {
	for _, tc := range dctTestData {
		hash, err := ImageHashDCT(tc.filename)
		if err != nil {
			t.Errorf("ImageHashDCT(%s): %v", tc.filename, err)
		} else if hash != tc.hash {
			t.Errorf("ImageHashDCT(%s): want %d, got %d", tc.filename, tc.hash, hash)
		}
	}
}

func TestHammingDistance(t *testing.T) {
	for _, tc := range hammingTestData {
		distance := HammingDistance(tc.h1, tc.h2)
		if distance != tc.distance {
			t.Errorf("want %d, got %d", tc.distance, distance)
		}
	}
}
