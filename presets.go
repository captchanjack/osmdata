package osmdata

import (
	"fmt"

	"github.com/captchanjack/osmdata/helpers"
	nom "github.com/captchanjack/osmdata/nominatim"
	ovp "github.com/captchanjack/osmdata/overpass"
)

type QueryMethod string

const (
	PlaceName   QueryMethod = "PlaceName"
	BoundingBox QueryMethod = "BoundingBox"
	Radius      QueryMethod = "Radius"
	Polygon     QueryMethod = "Polygon"
)

type PresetNetworkType string

const (
	Drive          PresetNetworkType = "Drive"
	DriveMainroads PresetNetworkType = "DriveMainroads"
	DriveService   PresetNetworkType = "DriveService"
	Walk           PresetNetworkType = "Walk"
	Bike           PresetNetworkType = "Bike"
	All            PresetNetworkType = "All"
	AllPrivate     PresetNetworkType = "AllPrivate"
	None           PresetNetworkType = "None"
	Rail           PresetNetworkType = "Rail"
)

var PresetWayTagFilters = map[PresetNetworkType][]ovp.TagFilter{
	Drive: {
		*ovp.NewTagFilter(ovp.Exists, "highway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "area", "yes"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "highway", "cycleway|footway|path|pedestrian|steps|track|corridor|elevator|escalator|proposed|construction|bridleway|abandoned|platform|raceway|service"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "motor_vehicle", "no"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "motorcar", "no"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "access", "private"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "service", "parking|parking_aisle|driveway|private|emergency_access"),
	},
	DriveMainroads: {
		*ovp.NewTagFilter(ovp.Exists, "highway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "area", "yes"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "highway", "cycleway|footway|path|pedestrian|steps|track|corridor|elevator|escalator|proposed|construction|bridleway|abandoned|platform|raceway|service|residential"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "motor_vehicle", "no"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "motorcar", "no"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "access", "private"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "service", "parking|parking_aisle|driveway|private|emergency_access"),
	},
	DriveService: {
		*ovp.NewTagFilter(ovp.Exists, "highway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "area", "yes"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "highway", "cycleway|footway|path|pedestrian|steps|track|corridor|elevator|escalator|proposed|construction|bridleway|abandoned|platform|raceway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "motor_vehicle", "no"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "motorcar", "no"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "access", "private"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "service", "parking|parking_aisle|private|emergency_access"),
	},
	Walk: {
		*ovp.NewTagFilter(ovp.Exists, "highway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "area", "yes"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "highway", "cycleway|motor|proposed|construction|abandoned|platform|raceway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "foot", "no"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "access", "private"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "service", "private"),
	},
	Bike: {
		*ovp.NewTagFilter(ovp.Exists, "highway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "area", "yes"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "highway", "footway|steps|corridor|elevator|escalator|motor|proposed|construction|abandoned|platform|raceway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "bicycle", "no"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "access", "private"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "service", "private"),
	},
	All: {
		*ovp.NewTagFilter(ovp.Exists, "highway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "area", "yes"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "highway", "proposed|construction|abandoned|platform|raceway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "access", "private"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "service", "private"),
	},
	AllPrivate: {
		*ovp.NewTagFilter(ovp.Exists, "highway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "area", "yes"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "highway", "proposed|construction|abandoned|platform|raceway"),
	},
	None: {
		*ovp.NewTagFilter(ovp.Exists, "highway"),
	},
	Rail: {
		*ovp.NewTagFilter(ovp.Exists, "railway"),
		*ovp.NewTagFilter(ovp.RegexNotMatch, "highway", "proposed|construction|abandoned|platform|raceway"),
	},
}

var PresetRelationTagFilters = map[PresetNetworkType][]ovp.TagFilter{
	Drive: {
		*ovp.NewTagFilter(ovp.Exists, "restriction"),
		*ovp.NewTagFilter(ovp.NotExists, "conditional"),
		*ovp.NewTagFilter(ovp.NotExists, "hgv"),
		*ovp.NewTagFilter(ovp.Equals, "type", "restriction"),
	},
	DriveMainroads: {
		*ovp.NewTagFilter(ovp.Exists, "restriction"),
		*ovp.NewTagFilter(ovp.NotExists, "conditional"),
		*ovp.NewTagFilter(ovp.NotExists, "hgv"),
		*ovp.NewTagFilter(ovp.Equals, "type", "restriction"),
	},
	DriveService: {
		*ovp.NewTagFilter(ovp.Exists, "restriction"),
		*ovp.NewTagFilter(ovp.NotExists, "conditional"),
		*ovp.NewTagFilter(ovp.NotExists, "hgv"),
		*ovp.NewTagFilter(ovp.Equals, "type", "restriction"),
	},
	Walk: {},
	Bike: {},
	All: {
		*ovp.NewTagFilter(ovp.Exists, "restriction"),
		*ovp.NewTagFilter(ovp.NotExists, "conditional"),
		*ovp.NewTagFilter(ovp.NotExists, "hgv"),
		*ovp.NewTagFilter(ovp.Equals, "type", "restriction"),
	},
	AllPrivate: {
		*ovp.NewTagFilter(ovp.Exists, "restriction"),
		*ovp.NewTagFilter(ovp.NotExists, "conditional"),
		*ovp.NewTagFilter(ovp.NotExists, "hgv"),
		*ovp.NewTagFilter(ovp.Equals, "type", "restriction"),
	},
	None: {},
	Rail: {},
}

// Get an executable overpass query object.
//
// The args of this function depends on the QueryMethod, the signature should match.
//
// *Note* CSV output format requires tag selection and header options specified in
// outputFormatOptions.
//
//	PlaceName: GetPresetQueryByPlaceName(
//		presetNetworkType PresetNetworkType,
//		includeMetadata bool,
//		outputFormat ovp.OutType,
//		placeName string,
//		outputFormatOptions ...string,
//	) *ovp.StackStatement
//
//	BoundingBox: GetPresetQueryByBoundingBox(
//		presetNetworkType PresetNetworkType,
//		includeMetadata bool,
//		outputFormat ovp.OutType,
//		minLat float64,
//		minLon float64,
//		maxLat float64,
//		maxLon float64,
//		outputFormatOptions ...string,
//	) *ovp.StackStatement
//
//	Radius: GetPresetQueryByRadius(
//		presetNetworkType PresetNetworkType,
//		includeMetadata bool,
//		outputFormat ovp.OutType,
//		radius float64,
//		lat float64,
//		lon float64,
//		outputFormatOptions ...string,
//	) *ovp.StackStatement
//
//	Polygon: GetPresetQueryByPolygon(
//		presetNetworkType PresetNetworkType,
//		includeMetadata bool,
//		outputFormat ovp.OutType,
//		polygonCoordinates *[][][]float64,
//		outputFormatOptions ...string,
//	) *ovp.StackStatement
func GetPresetQuery(queryMethod QueryMethod, args ...interface{}) *ovp.StackStatement {
	var getter interface{}

	if queryMethod == PlaceName {
		getter = GetPresetQueryByPlaceName
	} else if queryMethod == BoundingBox {
		getter = GetPresetQueryByBoundingBox
	} else if queryMethod == Radius {
		getter = GetPresetQueryByRadius
	} else if queryMethod == Polygon {
		getter = GetPresetQueryByPolygon
	}

	result := helpers.ExecVaradicFunction(getter, args...)
	return result[0].Interface().(*ovp.StackStatement)
}

func GetPresetQueryByPolygon(
	presetNetworkType PresetNetworkType,
	includeMetadata bool,
	outputFormat ovp.OutType,
	polygonCoordinates [][][]float64,
	outputFormatOptions ...string,
) *ovp.StackStatement {
	settings := ovp.NewSettingsStatement(*ovp.NewSetting(ovp.Out, outputFormat, outputFormatOptions...))
	stack := ovp.NewStackStatement(settings)

	wayFilters := PresetWayTagFilters[presetNetworkType]
	relationFilters := PresetRelationTagFilters[presetNetworkType]

	var poly *ovp.ElementFilter
	var way ovp.Statement
	var relation ovp.Statement

	union := ovp.NewUnionStatement("_")
	recurse := ovp.NewRecurseStatement(ovp.RecurseDown)

	for _, coords := range polygonCoordinates {
		poly = ovp.NewElementFilter(ovp.PolygonFilter, coords)
		way = ovp.NewElementStatement(ovp.Way, wayFilters, poly)
		relation = ovp.NewElementStatement(ovp.Relation, relationFilters, poly)
		union.Append(way, recurse, relation, recurse)
	}

	body := ovp.NewOutStatement(ovp.Body)

	stack.Append(union, body)

	if includeMetadata {
		stack.Append(ovp.NewOutStatement(ovp.Meta))
	}

	return stack
}

func GetPresetQueryByPlaceName(
	presetNetworkType PresetNetworkType,
	includeMetadata bool,
	outputFormat ovp.OutType,
	placeName string,
	outputFormatOptions ...string,
) *ovp.StackStatement {
	polygon, err := nom.QueryNominatim(placeName)

	if err != nil {
		fmt.Println(fmt.Errorf("encountered error during POST request to Overpass API: %s", err))
	}

	return GetPresetQueryByPolygon(
		presetNetworkType,
		includeMetadata,
		outputFormat,
		polygon,
		outputFormatOptions...
	)
}

func GetPresetQueryByRadius(
	presetNetworkType PresetNetworkType,
	includeMetadata bool,
	outputFormat ovp.OutType,
	radius float64,
	lat float64,
	lon float64,
	outputFormatOptions ...string,
) *ovp.StackStatement {
	settings := ovp.NewSettingsStatement(*ovp.NewSetting(ovp.Out, outputFormat, outputFormatOptions...))
	wayFilters := PresetWayTagFilters[presetNetworkType]
	relationFilters := PresetRelationTagFilters[presetNetworkType]
	around := ovp.NewElementFilter(ovp.AroundFilter, radius, [][]float64{{lon, lat}})
	way := ovp.NewElementStatement(ovp.Way, wayFilters, around)
	relation := ovp.NewElementStatement(ovp.Relation, relationFilters, around)
	recurse := ovp.NewRecurseStatement(ovp.RecurseDown)
	union := ovp.NewUnionStatement("_", way, recurse, relation, recurse)
	body := ovp.NewOutStatement(ovp.Body)
	stack := ovp.NewStackStatement(settings, union, body)

	if includeMetadata {
		stack.Append(ovp.NewOutStatement(ovp.Meta))
	}

	return stack
}

func GetPresetQueryByBoundingBox(
	presetNetworkType PresetNetworkType,
	includeMetadata bool,
	outputFormat ovp.OutType,
	minLat float64,
	minLon float64,
	maxLat float64,
	maxLon float64,
	outputFormatOptions ...string,
) *ovp.StackStatement {
	bbox := ovp.NewElementFilter(ovp.BoundingBoxFilter, minLat, minLon, maxLat, maxLon)
	settings := ovp.NewSettingsStatement(
		*ovp.NewSetting(ovp.Out, outputFormat, outputFormatOptions...),
		*ovp.NewSetting(ovp.GlobalBoundingBox, bbox.FilterStr),
	)
	wayFilters := PresetWayTagFilters[presetNetworkType]
	relationFilters := PresetRelationTagFilters[presetNetworkType]
	way := ovp.NewElementStatement(ovp.Way, wayFilters, bbox)
	relation := ovp.NewElementStatement(ovp.Relation, relationFilters, bbox)
	recurse := ovp.NewRecurseStatement(ovp.RecurseDown)
	union := ovp.NewUnionStatement("_", way, recurse, relation, recurse)
	body := ovp.NewOutStatement(ovp.Body)
	stack := ovp.NewStackStatement(settings, union, body)

	if includeMetadata {
		stack.Append(ovp.NewOutStatement(ovp.Meta))
	}

	return stack
}
