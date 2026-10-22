package models

type Artist struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"-"`
	ConcertDates []string            `json:"-"`
	Relations    map[string][]string `json:"-"`
	Tracks       []Track             `json:"-"`
	ShowTracks   bool                `json:"-"`
}
type LocationTemplate struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations,omitempty"`
}
type Location struct {
	Index []LocationTemplate `json:"index"`
}
type DateTemplate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates,omitempty"`
}
type Date struct {
	Index []DateTemplate `json:"index"`
}
type RelationTemplate struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type Relation struct {
	Index []RelationTemplate `json:"index"`
}
type Track struct {
	TrackName  string `json:"trackName"`
	PreviewURL string `json:"previewUrl"`
	ArtworkURL string `json:"artworkUrl100"`
}
type TrackResponse struct {
	Results []Track `json:"results"`
}
