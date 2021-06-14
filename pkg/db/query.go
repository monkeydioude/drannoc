package db

type Query struct {
	Filters Filter
	Options FindOptions
}

type KVBag struct {
	K string
	V interface{}
}

//NewQuery creates a new Query struct
func NewQuery() Query {
	return Query{
		Filters: make(Filter),
	}
}

// WithFilters create a new Query with filters
func (q Query) WithFilters(kvs ...KVBag) Query {
	for _, kv := range kvs {
		q.Filters.Add(kv.K, kv.V)
	}

	return q
}

// WithProjs create a new Query with projections
func (q Query) WithProjs(kvs ...KVBag) Query {
	for _, kv := range kvs {
		q.Options.Proj(kv.K, kv.V)
	}

	return q
}

func KV(k string, v interface{}) KVBag {
	return KVBag{
		K: k,
		V: v,
	}
}
