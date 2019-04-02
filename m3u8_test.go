package radiko

import (
	"testing"
)

func TestGetM3U8Chunks(t *testing.T) {
	// FIXME: Replace radio here as radio bangumi expires from time to time
	chunks, err := GetM3U8Chunks("JP13", "QRR", "20190331220000")
	if err != nil {
		t.Errorf("Error when trying to get m3u8 chunks: %s", err)
	}

	if len(chunks) == 0 {
		t.Errorf("length(chunk) should not be 0!")
	}

	if chunks[0] != "https://media.radiko.jp/sound/b/QRR/20190331/20190331_220000_WNH24.aac" {
		t.Errorf("Wrong download link of m3u8!")
	}
}
