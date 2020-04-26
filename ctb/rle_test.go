package ctb

import (
	"image"
	"testing"
)

func TestRleEncodeGraymap(t *testing.T) {
	rect := image.Rect(0, 0, 8, 21)
	gray := &image.Gray{
		Pix: []uint8{
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 00
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // 08
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 10
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 18
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 20
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 28
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 30
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 38
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 40
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 48
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 50
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 58
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 60
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 68
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 70
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 78
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 80
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // 88
			0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0xff, // 90
			0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, // 98
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // a0
		},
		Stride: rect.Size().X,
		Rect:   rect,
	}

	out_rle := []byte{
		0xff, 0x08,
		0x80, 0x08,
		0xff, 0x80, 0x85,
		0x80, 0x02,
		0xff, 0x02,
		0x00, 0x7f, 0x00, 0x7f, 0x00, 0x7f,
		0x80, 0x09,
	}

	out_hash := uint64(0x1be2583a56fdcbe9)
	out_bits := uint(146)

	rle, hash, bits := rleEncodeGraymap(gray)

	if bits != out_bits {
		t.Errorf("expected %v, got %v", out_bits, bits)
	}

	if out_hash != hash {
		t.Errorf("expected %#v, got %#v", out_hash, hash)
	}

	if len(rle) != len(out_rle) {
		t.Fatalf("expected %v, got %v", len(out_rle), len(rle))
	}

	for n, b := range out_rle {
		if rle[n] != b {
			t.Errorf("%v: expected %#v, got %#v", n, b, rle[n])
		}
	}

	// All empty
	rect = image.Rect(0, 0, 127, 4)
	gray.Rect = rect
	gray.Stride = rect.Size().X
	gray.Pix = make([]byte, rect.Size().X*rect.Size().Y)

	rle, hash, bits = rleEncodeGraymap(gray)

	out_rle = []byte{0x80, 0x81, 0xfc}
	out_hash = uint64(0x6af46758cc323d17)
	out_bits = 0

	if out_bits != bits {
		t.Errorf("expected %v, got %v", out_bits, bits)
	}

	if out_hash != hash {
		t.Errorf("expected %#v, got %#v", out_hash, hash)
	}

	if len(rle) != len(out_rle) {
		t.Fatalf("expected %v, got %v [% 02x]", len(out_rle), len(rle), rle)
	}

	for n, b := range out_rle {
		if rle[n] != b {
			t.Errorf("%v: expected %#v, got %#v", n, b, rle[n])
		}
	}

}
