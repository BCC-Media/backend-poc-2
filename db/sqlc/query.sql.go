// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
)

const getAllMedias = `-- name: GetAllMedias :many
SELECT id, collectable_type, media_type, primary_group_id, subclipped_media_id, reference_media_id, sequence_number, start_time, end_time, asset_id, agerating, created_at, updated_at FROM media
ORDER BY name
`

func (q *Queries) GetAllMedias(ctx context.Context) ([]Media, error) {
	rows, err := q.db.QueryContext(ctx, getAllMedias)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Media
	for rows.Next() {
		var i Media
		if err := rows.Scan(
			&i.ID,
			&i.CollectableType,
			&i.MediaType,
			&i.PrimaryGroupID,
			&i.SubclippedMediaID,
			&i.ReferenceMediaID,
			&i.SequenceNumber,
			&i.StartTime,
			&i.EndTime,
			&i.AssetID,
			&i.Agerating,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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

const getMedia = `-- name: GetMedia :one
SELECT id, collectable_type, media_type, primary_group_id, subclipped_media_id, reference_media_id, sequence_number, start_time, end_time, asset_id, agerating, created_at, updated_at FROM media
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMedia(ctx context.Context, id int64) (Media, error) {
	row := q.db.QueryRowContext(ctx, getMedia, id)
	var i Media
	err := row.Scan(
		&i.ID,
		&i.CollectableType,
		&i.MediaType,
		&i.PrimaryGroupID,
		&i.SubclippedMediaID,
		&i.ReferenceMediaID,
		&i.SequenceNumber,
		&i.StartTime,
		&i.EndTime,
		&i.AssetID,
		&i.Agerating,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getMediasWithFilter = `-- name: GetMediasWithFilter :many
SELECT id, collectable_type, media_type, primary_group_id, subclipped_media_id, reference_media_id, sequence_number, start_time, end_time, asset_id, agerating, created_at, updated_at FROM media

ORDER BY name
`

func (q *Queries) GetMediasWithFilter(ctx context.Context) ([]Media, error) {
	rows, err := q.db.QueryContext(ctx, getMediasWithFilter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Media
	for rows.Next() {
		var i Media
		if err := rows.Scan(
			&i.ID,
			&i.CollectableType,
			&i.MediaType,
			&i.PrimaryGroupID,
			&i.SubclippedMediaID,
			&i.ReferenceMediaID,
			&i.SequenceNumber,
			&i.StartTime,
			&i.EndTime,
			&i.AssetID,
			&i.Agerating,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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