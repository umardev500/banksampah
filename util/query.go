package util

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/umardev500/banksampah/types"
)

func NewQueryParams(page, limit int, filter []types.Filter, order types.Order) *types.QueryParam {
	return &types.QueryParam{
		Pagination: types.Pagination{
			Page:   page,
			Limit:  limit,
			Offset: (page - 1) * limit,
			Total:  0,
		},
		Filter: filter,
		Order:  order,
	}
}

func BuildUpdateQuery(baseQuery string, payload interface{}, filter []types.Filter) (string, []interface{}) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(baseQuery)

	var args []interface{}
	argIndex := 1

	v := reflect.ValueOf(payload)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldTag := v.Type().Field(i).Tag.Get("db")
		if !field.IsZero() {
			if argIndex > 1 {
				queryBuilder.WriteString(",")
			}
			fieldValue := field.Interface()
			args = append(args, fieldValue)
			queryBuilder.WriteString(fmt.Sprintf(" %s = $%d", fieldTag, argIndex))
			argIndex++
		}
	}

	if argIndex == 1 {
		return "", nil
	}

	// Add filter
	if len(filter) > 0 {
		queryBuilder.WriteString(" WHERE ")
		for i, filter := range filter {
			if i > 0 {
				queryBuilder.WriteString(fmt.Sprintf(" %s ", filter.LogicalOperator))
			}
			queryBuilder.WriteString(filter.Field)
			queryBuilder.WriteString(fmt.Sprintf(" %s ", filter.Operator))
			queryBuilder.WriteString(fmt.Sprintf("$%d", argIndex))
			args = append(args, filter.Value)
			argIndex++
		}
	}

	return queryBuilder.String(), args
}

func BuildQuery(baseQuery string, params *types.QueryParam) (string, []interface{}) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(baseQuery)

	var args []interface{}
	argIndex := 1

	// Add filter
	if len(params.Filter) > 0 {
		queryBuilder.WriteString(" WHERE ")
		for i, filter := range params.Filter {
			if i > 0 {
				queryBuilder.WriteString(fmt.Sprintf(" %s ", filter.LogicalOperator))
			}
			queryBuilder.WriteString(filter.Field)
			queryBuilder.WriteString(fmt.Sprintf(" %s ", filter.Operator))
			queryBuilder.WriteString(fmt.Sprintf("$%d", argIndex))
			args = append(args, filter.Value)
			argIndex++
		}
	}

	// Add ordering
	if params.Order.Field != "" {
		queryBuilder.WriteString(" ORDER BY ")
		queryBuilder.WriteString(params.Order.Field)
		queryBuilder.WriteString(" ")
		queryBuilder.WriteString(params.Order.Dir)
	}

	// Add pagination
	if params.Pagination.Page > 0 && params.Pagination.Limit > 0 {
		queryBuilder.WriteString(" LIMIT ")
		queryBuilder.WriteString(fmt.Sprintf("%d", params.Pagination.Limit))
		queryBuilder.WriteString(" OFFSET ")
		queryBuilder.WriteString(fmt.Sprintf("%d", params.Pagination.Offset))
	}

	return queryBuilder.String(), args
}
