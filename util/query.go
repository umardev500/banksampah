package util

import (
	"fmt"
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

func BuildQuery(baseQuery string, params *types.QueryParam) string {
	var queryBuilder strings.Builder
	queryBuilder.WriteString(baseQuery)

	// Add filter
	if len(params.Filter) > 0 {
		queryBuilder.WriteString(" WHERE ")
		for i, filter := range params.Filter {
			if i > 0 {
				queryBuilder.WriteString(" AND ")
			}
			queryBuilder.WriteString(filter.Field)
			queryBuilder.WriteString(fmt.Sprintf(" %s ", filter.Operator))
			queryBuilder.WriteString(fmt.Sprintf("'%s'", filter.Value))
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

	return queryBuilder.String()
}
