package db

// Filter is as shortcut to map[string]interface{} (aka json)
type Filter map[string]interface{}

func (f Filter) Add(k string, v interface{}) {
	f[k] = v
}

func (f Filter) AddFilter(k, fk string, v interface{}) {
	filter := Filter{}

	if _, ok := f[k]; ok {
		switch f[k].(type) {
		case Filter:
			filter = f[k].(Filter)
		}
	}
	filter[fk] = v
	f.Add(k, filter)
}
