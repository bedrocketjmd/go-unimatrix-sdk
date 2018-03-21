package unimatrix

import (
	"strconv"
)

type Query struct {
	parameters map[string][]string
}

func NewQuery() *Query {
	return &Query{parameters: map[string][]string{}}
}

func (query *Query) Parameters() map[string][]string {
	return query.parameters
}

func (query *Query) Where(parameter, value string) *Query {
	query.parameters[parameter] = []string{value}
	return query
}

func (query *Query) WhereArray(parameter string, values []string) *Query {
	parameter += "[]"
	query.parameters[parameter] = values
	return query
}

func (query *Query) Count(value int) *Query {
	query.parameters["count"] = []string{strconv.Itoa(value)}
	return query
}

func (query *Query) Offset(value int) *Query {
	query.parameters["offset"] = []string{strconv.Itoa(value)}
	return query
}

func (query *Query) Sort(sort_by, sort_direction string) *Query {
	query.parameters["sort_by"] = []string{sort_by}
	query.parameters["sort_direction"] = []string{sort_direction}
	return query
}

func (query *Query) Include(parameters ...string) *Query {
	includeParameter := "include"
	for _, parameter := range parameters {
		includeParameter += "[" + parameter + "]"
	}
	query.parameters[includeParameter] = []string{"true"}
	return query
}
