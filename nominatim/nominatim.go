package nominatim

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/captchanjack/osmdata/helpers"
)

const nominatimSearchEndpoint = "https://nominatim.openstreetmap.org/search"

func getNominatimSearchParams(placeName string) map[string]string {
	return map[string]string{
		"format":          "json",
		"limit":           "5",
		"dedupe":          "0",
		"polygon_geojson": "1",
		"q":               placeName,
	}
}

type GeneralMap map[string]interface{}

func QueryNominatim(placeName string) (coordinates [][][]float64, err error) {
	url := helpers.FormatHTTPGetURL(
		nominatimSearchEndpoint,
		getNominatimSearchParams(placeName),
	)

	resp, err := http.Get(url)

	if err != nil {
		return [][][]float64{}, fmt.Errorf("encountered error during GET request to Nominatim API: %w", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return [][][]float64{}, fmt.Errorf("encountered error during GET request to Nominatim API: %w", err)
	}

	var result []NominatimItem
	err = json.Unmarshal(body, &result)

	for i := range result {
		item := &result[i]

		switch item.GeoJSON.Type {

		case "Polygon":
			fmt.Printf("Using Polygon for %s\n", item.DisplayName)
			err = json.Unmarshal(item.GeoJSON.Coordinates, &item.GeoJSON.Polygon.Coordinates)

			if err != nil {
				fmt.Println(fmt.Errorf("%s: failed to convert %s: %s", item.DisplayName, item.GeoJSON.Type, err))
				continue
			}

			return item.GeoJSON.Polygon.Coordinates, nil

		case "MultiPolygon":
			fmt.Printf("Using Polygon for %s\n", item.DisplayName)
			err = json.Unmarshal(item.GeoJSON.Coordinates, &item.GeoJSON.MultiPolygon.Coordinates)

			if err != nil {
				fmt.Println(fmt.Errorf("%s: failed to convert %s: %s", item.DisplayName, item.GeoJSON.Type, err))
				continue
			}

			// First element is the polygon, subsequent elements are holes
			return item.GeoJSON.MultiPolygon.Coordinates[0], nil

		default:
			fmt.Printf("unsupported GeoJSON type '%s'\n", item.GeoJSON.Type)
			continue

		}
	}

	return [][][]float64{}, fmt.Errorf("could not find suitable polygon for input place name '%s'", placeName)
}
