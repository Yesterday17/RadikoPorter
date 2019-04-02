package radiko

import "fmt"

// Record 22
func Record(area, stationID, start, path string) error {
	chunks, err := GetM3U8Chunks(area, stationID, start)
	if err != nil {
		return err
	}

	result := chunks.Download(path)
	if len(result) != 0 {
		return fmt.Errorf("Failed to download the %d aac files", len(result))
	}
	return nil
}
