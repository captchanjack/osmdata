package overpass

type Statement interface {
	compile() string
	GetCompiled() string
	GetSetName() string
}

type ElementStatement struct {
	ElementType   ElementType
	TagFilters    []TagFilter
	ElementFilter *ElementFilter
	compiled      string
}

type UnionStatement struct {
	SetName    string `default:"_"`
	Statements []Statement
	compiled   string
}

type DifferenceStatement struct {
	SetName         string `default:"_"`
	FirstStatement  Statement
	SecondStatement Statement
	compiled        string
}

type SettingsStatement struct {
	Settings []Setting
	compiled string
}

type SetStatement struct {
	SetName  string
	compiled string
}

type OutStatement struct {
	VebosityType VebosityType
	compiled     string
}

type RecurseStatement struct {
	RecurseType RecurseType
	compiled    string
}

type StackStatement struct {
	Statements []Statement
	compiled   string
}

type TagFilter struct {
	TagFilterType TagFilterType
	Key           string
	Value         string
}

type ElementFilter struct {
	ElementFilterType ElementFilterType
	FilterStr         string
}

type Setting struct {
	Key     SettingType
	Value   string
	Options string
}

type ElementType string

const (
	Node     ElementType = "node"
	Way      ElementType = "way"
	Relation ElementType = "relation"
	Area     ElementType = "area"
)

type ElementRecurseType string

const (
	W  ElementRecurseType = "w"  // Forward from ways, e.g. node(w); select child nodes from ways in the input set
	R  ElementRecurseType = "r"  // Forward from relations, e.g. node(r); select node members of relations in the input set
	BN ElementRecurseType = "bn" // Backward from nodes, e.g. way(bn); select parent ways for nodes in the input set
	BW ElementRecurseType = "bw" // Backward from ways, e.g. rel(bw); select relations that have way members in the input set
	BR ElementRecurseType = "br" // Backward from relations, e.g. rel(br); select parent relations from relations in the input set
)

type SettingType string

const (
	Out               SettingType = "out"     // The out setting defines the output format used to return OSM data, default value is xml, see OutType
	Timeout           SettingType = "timeout" // The timeout: setting has one parameter, a non-negative integer. Default value is 180
	Maxsize           SettingType = "maxsize" // The maxsize: setting has one parameter, a non-negative integer. Default value is 536870912 (512 MB)
	GlobalBoundingBox SettingType = "bbox"    // The 'bounding box' defines the map area that the query will include
	Date              SettingType = "date"    // date is a global setting which modifies an Overpass QL query to examine attic data, and return results based on the OpenStreetMap database as of the date specified
	Diff              SettingType = "diff"    // The diff setting lets the database determine the difference of two queries at different points in time
)

type OutType string

const (
	XML  OutType = "xml"  // Default format
	JSON OutType = "json" // Not GeoJSON
	CSV  OutType = "csv"  // Must specify as [out:csv( fieldname_1 [,fieldname_n ...] [; csv-headerline [; csv-separator-character ] ] )]
)

type VebosityType string

const (
	Ids       VebosityType = "ids"        // Print only the ids of the elements in the set
	IdsNoIds  VebosityType = "ids noids"  // The modificator noids lets all ids to be omitted from the statement
	Skel      VebosityType = "skel"       // Print the minimum information necessary for geometry (nodes: id and coordinates, ways: id and the ids of its member nodes, relations: id of the relation, and the id, type, and role of all of its members)
	SkelNoIds VebosityType = "skel noids" // The modificator noids lets all ids to be omitted from the statement
	Body      VebosityType = "body"       // Print all information necessary to use the data. These are also tags for all elements and the roles for relation members
	BodyNoIds VebosityType = "body noids" // The modificator noids lets all ids to be omitted from the statement
	Tags      VebosityType = "tags"       // Print only ids and tags for each element and not coordinates or members
	TagsNoIds VebosityType = "tags noids" // The modificator noids lets all ids to be omitted from the statement
	Meta      VebosityType = "meta"       // Print everything known about the elements. meta includes everything output by body for each OSM element, as well as the version, changeset id, timestamp, and the user data of the user that last touched the object. Derived elements' metadata attributes are also missing for derived elements
	MetaNoIds VebosityType = "meta noids" // The modificator noids lets all ids to be omitted from the statement
	Geom      VebosityType = "geom"       // Add the full geometry to each object. This adds coordinates to each node, to each node member of a way or relation, and it adds a sequence of "nd" members with coordinates to all relations
	BB        VebosityType = "bb"         // Adds only the bounding box of each element to the element. For nodes this is equivalent to "geom". For ways it is the enclosing bounding box of all nodes. For relations it is the enclosing bounding box of all node and way members, relations as members have no effect
	Center    VebosityType = "center"     // This adds only the center of the above mentioned bounding box to ways and relations. Note: The center point is not guaranteed to lie inside the polygon
	ASC       VebosityType = "asc"        // Sort by object id
	QT        VebosityType = "qt"         // Sort by quadtile index; this is roughly geographical and significantly faster than order by ids (derived elements generated by make or convert statements without any geometry will be grouped separately, only sorted by id)
	Count     VebosityType = "count"      // prints only the total counts of elements in the input set by type (nodes, ways, relations, areas). It cannot be combined with anything else
)

type RecurseType string

const (
	RecurseDown          RecurseType = "<"  // All nodes that are part of a way which appears in the input set; plus all nodes and ways that are members of a relation which appears in the input set; plus all nodes that are part of a way which appears in the result set
	RecurseDownRelations RecurseType = "<<" // Continues to follow the membership links including nodes in ways until for every object in its input or result set all the members of that object are in the result set as well.
	RecurseUp            RecurseType = ">"  // All ways that have a node which appears in the input set; plus all relations that have a node or way which appears in the input set; plus all relations that have a way which appears in the result set
	RecurseUpRelations   RecurseType = ">>" // Continues to follow backlinks onto the found relations until it contains all relations that point to an object in the input or result set
)

type TagFilterType string

const (
	Equals        TagFilterType = "="  // The most common variant selects all elements where the tag with the given key has a specific value
	NotEquals     TagFilterType = "!=" // The most common variant selects all elements where the tag with the given key does not have a specific value
	Exists        TagFilterType = ""   // The second variant selects all elements that have a tag with a certain key and an arbitrary value
	NotExists     TagFilterType = "!"  // This variant selects all element, that don't have a tag with a certain key and an arbitrary value
	RegexMatch    TagFilterType = "~"  // The third variant selects all elements that have a tag with a certain key and a value that matches some regular expression
	RegexNotMatch TagFilterType = "!~" // The third variant selects all elements that have a tag with a certain key and a value that does not matche some regular expression
)

type ElementFilterType string

const (
	BoundingBoxFilter ElementFilterType = "BoundingBoxFilter" // The bounding box query filter selects all elements within a rectangular bounding box, (south,west,north,east), e.g. node(50.6,7.0,50.8,7.3);
	RecurseFilter     ElementFilterType = "RecurseFilter"     // The recurse filter selects all elements that are members of an element from the input set or have an element of the input set as member, depending on the given parameter, e.g. node(w)
	IDFilter          ElementFilterType = "IDFilter"          // The id-query filter selects the element of given type with given id
	AroundFilter      ElementFilterType = "AroundFilter"      // The around filter selects all elements within a certain radius in metres around the elements in the input set
	PolygonFilter     ElementFilterType = "PolygonFilter"     // The third variant selects all elements that have a tag with a certain key and a value that matches some regular expression
	AreaFilter        ElementFilterType = "AreaFilter"        // The pivot filter selects the element of the chosen type that defines the outline of the given area
)
