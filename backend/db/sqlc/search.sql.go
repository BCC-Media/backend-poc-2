// Code generated by sqlc. DO NOT EDIT.
// source: search.sql

package db

import (
	"context"
	"time"

	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getMediaSearchDocs = `-- name: GetMediaSearchDocs :many
SELECT 
    m.status, m.type, m.available_from, m.available_to, m.title, m.description, m.long_description, m.image_id, m.translation_id, m.id, m.collectable_type, m.media_type, m.primary_group_id, m.subclipped_media_id, m.reference_media_id, m.sequence_number, m.start_time, m.end_time, m.asset_id, m.agerating, m.created_at, m.updated_at, m.published_time, m.usergroups, m.tags,
    parent.primary_group_id as grandparent_id
FROM admin.media m
LEFT JOIN admin.media parent ON parent.id = m.primary_group_id
`

type GetMediaSearchDocsRow struct {
	Status            int16          `db:"status" json:"status"`
	Type              null_v4.String `db:"type" json:"type"`
	AvailableFrom     null_v4.Time   `db:"available_from" json:"availableFrom"`
	AvailableTo       null_v4.Time   `db:"available_to" json:"availableTo"`
	Title             null_v4.String `db:"title" json:"title"`
	Description       null_v4.String `db:"description" json:"description"`
	LongDescription   null_v4.String `db:"long_description" json:"longDescription"`
	ImageID           null_v4.Int    `db:"image_id" json:"imageID"`
	TranslationID     null_v4.Int    `db:"translation_id" json:"translationID"`
	ID                int64          `db:"id" json:"id"`
	CollectableType   null_v4.String `db:"collectable_type" json:"collectableType"`
	MediaType         null_v4.String `db:"media_type" json:"mediaType"`
	PrimaryGroupID    null_v4.Int    `db:"primary_group_id" json:"primaryGroupID"`
	SubclippedMediaID null_v4.Int    `db:"subclipped_media_id" json:"subclippedMediaID"`
	ReferenceMediaID  null_v4.Int    `db:"reference_media_id" json:"referenceMediaID"`
	SequenceNumber    int16          `db:"sequence_number" json:"sequenceNumber"`
	StartTime         null_v4.Float  `db:"start_time" json:"startTime"`
	EndTime           null_v4.Float  `db:"end_time" json:"endTime"`
	AssetID           null_v4.Int    `db:"asset_id" json:"assetID"`
	Agerating         null_v4.String `db:"agerating" json:"agerating"`
	CreatedAt         time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updatedAt"`
	PublishedTime     null_v4.Time   `db:"published_time" json:"publishedTime"`
	Usergroups        pq.StringArray `db:"usergroups" json:"usergroups"`
	Tags              pq.Int64Array  `db:"tags" json:"tags"`
	GrandparentID     null_v4.Int    `db:"grandparent_id" json:"grandparentID"`
}

func (q *Queries) GetMediaSearchDocs(ctx context.Context) ([]GetMediaSearchDocsRow, error) {
	rows, err := q.db.QueryContext(ctx, getMediaSearchDocs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMediaSearchDocsRow
	for rows.Next() {
		var i GetMediaSearchDocsRow
		if err := rows.Scan(
			&i.Status,
			&i.Type,
			&i.AvailableFrom,
			&i.AvailableTo,
			&i.Title,
			&i.Description,
			&i.LongDescription,
			&i.ImageID,
			&i.TranslationID,
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
			&i.PublishedTime,
			&i.Usergroups,
			&i.Tags,
			&i.GrandparentID,
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