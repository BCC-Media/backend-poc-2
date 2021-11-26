// Code generated by sqlc. DO NOT EDIT.
// source: languages.sql

package db

import (
	"context"
)

const getLanguages = `-- name: GetLanguages :many
SELECT code, name from public.language
`

func (q *Queries) GetLanguages(ctx context.Context) ([]Language, error) {
	rows, err := q.db.QueryContext(ctx, getLanguages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Language
	for rows.Next() {
		var i Language
		if err := rows.Scan(&i.Code, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
