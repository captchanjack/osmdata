package overpass

func NewTagFilter(tagFilterType TagFilterType, key string, value ...string) *TagFilter {
	f := new(TagFilter)
	f.TagFilterType = tagFilterType
	f.Key = key

	_value := ""
	if len(value) > 0 {
		_value = value[0]
	}
	f.Value = _value

	return f
}
