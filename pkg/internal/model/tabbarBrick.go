package model

type BrickTabbar struct {
	ID       string           `json:"id,omitempty"`
	Variants []TabbarVariants `json:"variants,omitempty"`
}
type TabbarData struct {
	AccessibilityID   string `json:"accessibility_id,omitempty"`
	ImageName         string `json:"image_name,omitempty"`
	IsDefault         bool   `json:"is_default"`
	RootDeeplink      string `json:"root_deeplink,omitempty"`
	SelectedImageName string `json:"selected_image_name,omitempty"`
	TabID             string `json:"tab_id,omitempty"`
	Title             string `json:"title"`
}
type Clients struct {
	ID           int64  `json:"id,omitempty"`
	SinceVersion string `json:"since_version,omitempty"`
}
type TabbarSupport struct {
	Sites     []string  `json:"sites,omitempty"`
	Clients   []Clients `json:"clients,omitempty"`
	Audiences []string  `json:"audiences,omitempty"`
}
type TabbarVariants struct {
	ID      string          `json:"id,omitempty"`
	UIType  string          `json:"ui_type"`
	Data    TabbarData      `json:"data,omitempty"`
	Support []TabbarSupport `json:"support,omitempty"`
}
