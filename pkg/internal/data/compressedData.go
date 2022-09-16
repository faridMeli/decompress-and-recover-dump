package data

import "time"

type DataCompressed struct {
	Item struct {
		Key struct {
			S string `json:"S"`
		} `json:"key"`
		Metadata struct {
			S string `json:"S"`
		} `json:"metadata"`
		Version struct {
			N string `json:"N"`
		} `json:"version"`
		LastUpdated struct {
			S time.Time `json:"S"`
		} `json:"last_updated"`
		DateCreated struct {
			S time.Time `json:"S"`
		} `json:"date_created"`
		CompressedValue struct {
			B string `json:"B"`
		} `json:"compressed_value"`
		DecompressedData struct {
			S interface{} `json:"S"`
		} `json:"decompressed_value"`
		LastUpdatedMicros struct {
			N string `json:"N"`
		} `json:"last_updated_micros"`
	} `json:"Item"`
}
