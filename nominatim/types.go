package nominatim

import "encoding/json"

type GeoJSONPoint struct {
	Coordinates []float64
}

type GeoJSONLine struct {
	Coordinates [][]float64
}

type GeoJSONPolygon struct {
	Coordinates [][][]float64
}

type GeoJSONMultiPolygon struct {
	Coordinates [][][][]float64
}

type GeoJSON struct {
	Type         string          `json:"type"`
	Coordinates  json.RawMessage `json:"coordinates"`
	Point        GeoJSONPoint
	Line         GeoJSONLine
	Polygon      GeoJSONPolygon
	MultiPolygon GeoJSONMultiPolygon
}

type NominatimItem struct {
	OSMType     int       `json:"osm_type"`
	OSMID       int       `json:"osm_id"`
	GeoJSON     GeoJSON   `json:"geojson"`
	BoundingBox []float64 `json:"boundingbox"`
	DisplayName string    `json:"display_name"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"lon"`
}
