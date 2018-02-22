package unimatrix

import (
  "net/url"
)

type Query struct {
  parameters url.Values
}

func NewQuery() *Query {
  parameters := url.Values{}
  return &Query{
    parameters: parameters,
  }
}

func (query *Query) Parameters() string {
  return query.parameters.Encode()
}

func (query *Query) Where(parameter, value string) *Query {
  query.parameters.Add(parameter, value)
  return query
}

func (query *Query) WhereArray(parameter string, values []string) *Query {
  parameter += "[]"
  for _, value := range values {
    query.parameters.Add(parameter, value)
  }
  return query
}

func (query *Query) Include(parameters ...string) *Query {
  includeParameter := "include"
  for _, parameter := range parameters {
    includeParameter += "[" + parameter + "]"
  }
  query.parameters.Add(includeParameter, "true")
  return query
}
