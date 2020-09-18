package query

import (
	"github.com/olivere/elastic"
)

func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	equalsQueries := make([]elastic.Query, 0)
	for _, equal := range q.Equals {
		equalsQueries = append(equalsQueries, elastic.NewMatchQuery(equal.Field, equal.Value))
	}
	query.Must(equalsQueries...)
	return query
}
