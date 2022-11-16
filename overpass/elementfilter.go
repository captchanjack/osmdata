package overpass

import (
	"fmt"

	"github.com/captchanjack/osmdata/helpers"
)

// Returns a pointer to a new ElementFilter, this constructor will also compile the filter string
// needed for the final query.
//
// The args of this function depends on the elementFilterType, the signature should match:
//
//	BoundingBoxFilter: GetBoundingBoxFilterStr(south float64, west float64, north float64, east float64) string
//	RecurseFilter: GetRecurseFilterStr(elementRecurseType ElementRecurseType) string
//	IDFilter: GetIDFilterStr(id int, ids ...int) string
//	AroundFilter: GetAroundFilterStr(radius float64, coordinates [][]float64, inputSetName ...string) string
//	PolygonFilter: GetPolygonFilterStr(coordinates [][]float64) string
//	AreaFilter: GetAroundFilterStr(radius float64, coordinates []float64, setName ...string) string
func NewElementFilter(elementFilterType ElementFilterType, args ...interface{}) *ElementFilter {
	f := new(ElementFilter)
	f.ElementFilterType = elementFilterType

	var getter interface{}

	if elementFilterType == BoundingBoxFilter {
		getter = GetBoundingBoxFilterStr
	} else if elementFilterType == RecurseFilter {
		getter = GetRecurseFilterStr
	} else if elementFilterType == IDFilter {
		getter = GetIDFilterStr
	} else if elementFilterType == AroundFilter {
		getter = GetAroundFilterStr
	} else if elementFilterType == PolygonFilter {
		getter = GetPolygonFilterStr
	} else if elementFilterType == AreaFilter {
		getter = GetAroundFilterStr
	}

	result := helpers.ExecVaradicFunction(getter, args...)
	f.FilterStr = result[0].Interface().(string)

	return f
}

func GetBoundingBoxFilterStr(south float64, west float64, north float64, east float64) string {
	return fmt.Sprintf("%f,%f,%f,%f", south, west, north, east)
}

func GetRecurseFilterStr(elementRecurseType ElementRecurseType) string {
	return fmt.Sprintf("%s)", elementRecurseType)
}

func GetIDFilterStr(id int, ids ...int) string {
	idArray := append([]int{id}, ids...)
	return fmt.Sprintf("id:%s", helpers.JoinArrInt(idArray))
}

// Input coordinates for the polygon are in GeoJSON polygon format
// i.e. [[lon1, lat1], [lon2, lat2], ...]
// OSM expects string format "lat1,lon1,lat2,lon2, ..."
func GetAroundFilterStr(radius float64, coordinates [][]float64, inputSetName ...string) string {
	prefix := "around"
	if len(inputSetName) > 0 {
		prefix = fmt.Sprintf("%s.%s", prefix, inputSetName[0])
	}
	converted := convertPolygonCooridnates(coordinates)
	return fmt.Sprintf("%s:%v,%s", prefix, radius, helpers.JoinArrFloat64(converted))
}

// Input coordinates for the polygon are in GeoJSON polygon format
// i.e. [[lon1, lat1], [lon2, lat2], ...]
// OSM expects string format "lat1 lon1 lat2 lon2 ..."
func GetPolygonFilterStr(coordinates [][]float64) string {
	converted := convertPolygonCooridnates(coordinates)
	return fmt.Sprintf("poly:\"%s\"", helpers.JoinArrFloat64(converted, " "))
}

func GetAreaFilterStr(id int, ids ...int) string {
	idArray := append([]int{id}, ids...)
	return fmt.Sprintf("area:%s", helpers.JoinArrInt(idArray))
}

// Converts polygon coordinates conforming to GeoJSON [[lon1, lat1], [lon2, lat2], ...]
// to OSM format [lat1, lon1, lat2, lon2, ...]
func convertPolygonCooridnates(x [][]float64) []float64 {
	converted := make([]float64, len(x)*2)
	for i, v := range x {
		converted[2*i] = v[1]   // lat
		converted[2*i+1] = v[0] // lon
	}
	return converted
}
