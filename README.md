# osmdata
This is a Go library for downloading OpenStreetMap data using the [Overpass](https://overpass-turbo.eu/) and [Nominatim](https://nominatim.openstreetmap.org/ui/search.html) API.

## Core Features
- Building and executing preset queries against road, rail, walk or bicycle networks, with the flexibility of querying against a polygon, bounding box, radius, or a place name (as free-text)
- Supports XML, JSON and CSV data format and exporting to disk
- Building Overpass queries dynamically with `Statement` structs
- Executing raw Overpass string queries
- Searching the Nominatim API for polygons using free-text

## Usage

### Raw String Queries
Let's start with an simple Overpass query that retrieves all traffic light nodes in an 500m radius area somewhere in Melbourne, Australia:
```
import ovp "github.com/captchanjack/osmdata/overpass"

queryStr := "[out:json];node[\"highway\"=\"traffic_signals\"](around:500.0,-37.740347,144.930127);out body;"

resp, err := ovp.OverpassQuery(queryStr)
```
`resp` can be unmarshalled and should look something like this:
```
{
    "version": 0.6,
    "generator": "Overpass API 0.7.59 e21c39fe",
    "osm3s": {
        "timestamp_osm_base": "2022-11-15T12:23:09Z",
        "copyright": "The data included in this document is from www.openstreetmap.org. The data is made available under ODbL."
    },
    "elements": [
        {
            "type": "node",
            "id": 68717542,
            "lat": -37.7435041,
            "lon": 144.9282267,
            "tags": {
                "highway": "traffic_signals"
            }
        },
        {
            "type": "node",
            "id": 5547839707,
            "lat": -37.7410101,
            "lon": 144.9318164,
            "tags": {
                "highway": "traffic_signals",
                "traffic_signals": "ramp_meter",
                "traffic_signals:direction": "forward"
            }
        }
    ]
}
```

### Preset Queries

Mentioned earlier, this library offers a few useful preset queries against various network types and allows the parameterisation of spatial filters. This example shows how to query all nodes, ways and relations on a road network in Melbourne City, Austrlia:
```
import (
    "fmt"
    osm "github.com/captchanjack/osmdata"
    ovp "github.com/captchanjack/osmdata/overpass"
)

query := osm.GetPresetQuery(
    osm.PlaceName,                  // Query method, i.e.. BoundingBox, Radius, Polygon, PlaceName
    osm.Drive,                      // Preset network type, i.e. Drive, Walk, Bicycle, Rail, etc.
    true,                           // Whether to download metadata i.e. changeset ids etc.
    ovp.JSON,                       // Output format, supports XML, JSON and CSV
    "Melbourne City, Australia"     // Argument(s) specific to the query method, in this case the free-text describing the place name
)

fmt.Println(query.GetCompiled())    // Print the query
resp, err := query.Execute()        // Execute the query
```

### Building Queries

We can build queryies dynamically using `Statement` structs, let's try replicate the raw string query example from earlier:
```
import (
    "fmt"
    ovp "github.com/captchanjack/osmdata/overpass"
)

// Settings statement, output JSON format
// [out:json];
settingsStmt := ovp.NewSettingsStatement(*ovp.NewSetting(ovp.Out, ovp.JSON))

// Main block statement, filtering nodes that within a certain radius and are traffic signals
// node[\"highway\"=\"traffic_signals\"](around:500.0,-37.740347,144.930127);
tagFilters := []ovp.TagFilter{*ovp.NewTagFilter(ovp.Equals, "highway", "traffic_signals")}
elementFilters := ovp.NewElementFilter(ovp.AroundFilter, 500.0, [][]float64{{-37.740347, 144.930127})
nodeStmt := ovp.NewElementStatement(ovp.Node, tagFilters, elementFilters)

// Output statement, exporting the body data
// out body;
bodyStmt := ovp.NewOutStatement(ovp.Body)

// Finally we add all the statements into a stack
stackStmt := ovp.NewStackStatement(settingsStmt, nodeStmt, bodyStmt)

// We can inspect the query and execute it
fmt.Println(stackStmt.GetCompiled())
resp, err := stackStmt.Execute()
```

### Exporting
We can export the downloaded data to a file on disk:
```
import (
    "fmt"
    osm "github.com/captchanjack/osmdata"
    ovp "github.com/captchanjack/osmdata/overpass"
)

query := osm.GetPresetQuery(
    osm.PlaceName,                  // Query method, i.e.. BoundingBox, Radius, Polygon, PlaceName
    osm.Drive,                      // Preset network type, i.e. Drive, Walk, Bicycle, Rail, etc.
    true,                           // Whether to download metadata i.e. changeset ids etc.
    ovp.JSON,                       // Output format, supports XML, JSON and CSV
    "Melbourne City, Australia"     // Argument(s) specific to the query method, in this case the free-text describing the place name
)

resp, err := query.ExecuteAndExport("./test.json")        // Execute the query and export
```
