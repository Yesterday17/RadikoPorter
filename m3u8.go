package radiko

import (
	"context"
	"time"

	"github.com/yyoshiki41/go-radiko"
)

// GetM3U8Chunks returns chunk, err
func GetM3U8Chunks(area, stationID, start string) (M3U8ChunkList, error) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	client, err := radiko.New("")
	if err != nil {
		return nil, err
	}

	client.SetAreaID(area)
	_, err = client.AuthorizeToken(ctx)
	if err != nil {
		return nil, err
	}

	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	startTime, err := time.ParseInLocation("20060102150405", start, location)
	if err != nil {
		return nil, err
	}

	uri, err := client.TimeshiftPlaylistM3U8(ctx, stationID, startTime)
	if err != nil {
		return nil, err
	}

	chunks, err := radiko.GetChunklistFromM3U8(uri)
	if err != nil {
		return nil, err
	}

	return chunks, nil
}
