package geojson

type GeoJSON struct {
	Type       string      `json:"type"`
	Properties interface{} `json:"properties"`
	Geometry   interface{} `json:"geometry"`
}
