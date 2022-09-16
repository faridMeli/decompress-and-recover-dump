package model

type Collection struct {
	Title       string   `json:"title"`
	Color       string   `json:"color"`
	ShortcutIDs []string `json:"shortcut_ids"`
	ID          string   `json:"id"`
	SiteID      string   `json:"site_id"`
}
