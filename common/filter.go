package common

import (
	"fmt"
	"strings"
)

// SubQuery supports equality-based filtering, containing:
// key, based on which field the filtering is performed;
// op, only ops in validOps are supported;
// value, the given value used to compare with data persisted.
type SubQuery struct {
	Key   string
	Op    string
	Value string
}

var validFilterKeys map[string]bool = map[string]bool{
	"title":             true,
	"version":           true,
	"maintainers.name":  true,
	"maintainers.email": true,
	"company":           true,
	"website":           true,
	"source":            true,
	"license":           true,
}

var validOps []string = []string{"!=", "==", "="} // orders matter!

// divider is used to separate subqueries.
const divider = ","

// QueryError provides a formatted error with complete query string,
// the subquery that has issue, and the cause.
func QueryError(query, subquery, msg string) error {
	return fmt.Errorf("Query parsing error:\n\t-Query: %v\n\tSubQuery: %v\n\tReason: %v", query, subquery, msg)
}

// ParseQuery checks the query from client and generates a slice of SubQuery.
// For each subquery, it will check: (1) if it only contains one operator.
// (2) if it uses a valid filtering key. (3) if it contains a valid operators.
// If either one is not satisfied, an error will be returned.
func ParseQuery(query string) ([]SubQuery, error) {
	subqueries := strings.Split(query, divider)
	res := make([]SubQuery, 0) // assume in most cases we only has one subquery.

	for _, sub := range subqueries {
		valid := false
		var givenKey, givenOp, givenValue string

		for _, op := range validOps {
			if strings.Contains(sub, op) {
				parts := strings.Split(sub, op)
				if len(parts) != 2 {
					return res, QueryError(query, sub, "more than one operator.")
				}

				givenKey, givenOp, givenValue = strings.Trim(parts[0], " "), op, strings.Trim(parts[1], " ")
				if _, ok := validFilterKeys[givenKey]; !ok {
					return res, QueryError(query, sub, givenKey+" is not a valid key.")
				}

				// subquery parse succeeded.
				valid = true
				res = append(res, SubQuery{Key: givenKey, Op: givenOp, Value: givenValue})
				break
			}
		}

		// no op matched.
		if !valid {
			return res, QueryError(query, sub, "no valid op is found.")
		}
	}

	return res, nil
}

// ExecuteQuery executes a given list of subqueries on a metadata and return a
// bool representing if the metadata satisfies these queries.
func ExecuteQuery(metadata Metadata, subqueries []SubQuery) bool {
	for _, sub := range subqueries {
		if strings.Contains(sub.Key, "maintainers") {
			if !maintainerCompare(metadata.Maintainers, sub.Key, sub.Op, sub.Value) {
				return false
			}
		} else {
			var actual string
			switch sub.Key {
			case "title":
				actual = metadata.Title
			case "version":
				actual = metadata.Version
			case "company":
				actual = metadata.Company
			case "website":
				actual = metadata.Website
			case "source":
				actual = metadata.Source
			case "license":
				actual = metadata.License
			}

			if !generalCompare(actual, sub.Op, sub.Value) {
				return false
			}
		}
	}

	// all subqueries are satisfied, return true
	return true
}

// generalCompare is a helper function that does the actual comparision given
// all necessary components.
func generalCompare(actual, op, given string) bool {
	fmt.Printf("%v\t%v\t%v\n", actual, op, given)
	switch op {
	case "!=":
		return actual != given
	case "=", "==":
		return actual == given
	}
	return false
}

// maintainerCompare is similar to generalCompare. Since Maintainer is a struct,
// we need another wrapper on top of generalCompare to obtain the actual value.
// Also since a metadata has multiple Maintainer, we assume the comparision
// returns true as long as it returns true on one Maintainer.
func maintainerCompare(maintainers []Maintainer, key, op, given string) bool {
	key = strings.Split(key, ".")[1]
	for _, m := range maintainers {
		var actual string
		if key == "name" {
			actual = m.Name
		} else {
			actual = m.Email
		}

		if generalCompare(actual, op, given) {
			return true
		}
	}

	return false
}
