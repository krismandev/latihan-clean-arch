package helper

import "strings"

func AppendComma(baseQuery *string, baseParam *[]interface{}, appendedQuery string, value string) {
	if len(*baseQuery) > 0 {
		*baseQuery += " , "
	}

	if len(value) > 0 {
		*baseQuery += appendedQuery
		*baseParam = append(*baseParam, value)
	} else {
		*baseQuery += strings.ReplaceAll(appendedQuery, "?", "NULL")
	}
}

func AppendCommaRaw(baseQuery *string, appendedQuery string) {
	if len(appendedQuery) > 0 {
		if len(*baseQuery) > 0 {
			*baseQuery += " , "
		}
		*baseQuery += appendedQuery
	}
}
