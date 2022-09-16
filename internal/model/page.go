package model

type Page struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Data *Data  `json:"data"`
}

type Data struct {
	PageID string  `json:"page_id,omitempty"`
	Events []Event `json:"events"`
}

type Event struct {
	Type     string                 `json:"type"`
	Data     *DataInterface         `json:"data,omitempty"`
	Tracking map[string]interface{} `json:"tracking,omitempty"`
}

type DataInterface struct {
	Bricks  []PageBrick `json:"bricks,omitempty"`
	Mode    string      `json:"mode,omitempty"`
	BrickID string      `json:"brick_id,omitempty"`
}

type PageBrick struct {
	BrickDTO
	Bricks []PageBrick `json:"bricks,omitempty"`
}
