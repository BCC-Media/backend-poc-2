// Code generated by sqlc. DO NOT EDIT.
// source: medias.sql

package db

import (
	"context"
	"time"

	"github.com/lib/pq"
	null_v4 "gopkg.in/guregu/null.v4"
)

const getAsset = `-- name: GetAsset :one
SELECT id, source_id, published_version_id, name FROM asset
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAsset(ctx context.Context, id int64) (Asset, error) {
	row := q.db.QueryRowContext(ctx, getAsset, id)
	var i Asset
	err := row.Scan(
		&i.ID,
		&i.SourceID,
		&i.PublishedVersionID,
		&i.Name,
	)
	return i, err
}

const getMedia = `-- name: GetMedia :one
SELECT status, type, available_from, available_to, title, description, long_description, image_id, translation_id, id, collectable_type, media_type, primary_group_id, subclipped_media_id, reference_media_id, sequence_number, start_time, end_time, asset_id, agerating, created_at, updated_at, published_time, usergroups FROM admin.media
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMedia(ctx context.Context, id int64) (AdminMedium, error) {
	row := q.db.QueryRowContext(ctx, getMedia, id)
	var i AdminMedium
	err := row.Scan(
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
		pq.Array(&i.Usergroups),
	)
	return i, err
}

const insertMedia = `-- name: InsertMedia :one
WITH c AS (
    INSERT INTO collectable (
        type,
        available_from,
        available_to,
        published_time,
        status
    ) VALUES (
        'media',
        $1,
        $2,
        $3,
        $16
    ) RETURNING id, type, available_from, available_to, status, created_at, updated_at, published_time
),
m AS (
    INSERT INTO media (
        id,
        collectable_type,
        media_type,
        primary_group_id,
        subclipped_media_id,
        reference_media_id,
        sequence_number,
        start_time,
        end_time,
        asset_id,
        agerating
    ) SELECT
        c.id,
        c.type,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12
    FROM c RETURNING id, collectable_type, media_type, primary_group_id, subclipped_media_id, reference_media_id, sequence_number, start_time, end_time, asset_id, agerating, created_at, updated_at
),
t AS (
    INSERT INTO media_t (
        media_id,
        language_code,
        title,
        description,
        long_description
    ) SELECT 
        c.id,
        'no',
        $13,
        $14,
        $15
    FROM c RETURNING id, media_id, language_code, title, description, long_description, image_id
)
SELECT 
    c.status,
    c.type,
    c.published_time,
    c.available_from,
    c.available_to,
	t.title,
	t.description,
	t.long_description,
	t.image_id,
	t.id as translation_id,
    m.id, m.collectable_type, m.media_type, m.primary_group_id, m.subclipped_media_id, m.reference_media_id, m.sequence_number, m.start_time, m.end_time, m.asset_id, m.agerating, m.created_at, m.updated_at
    FROM c,t,m
`

type InsertMediaParams struct {
	AvailableFrom     null_v4.Time   `db:"available_from" json:"availableFrom"`
	AvailableTo       null_v4.Time   `db:"available_to" json:"availableTo"`
	PublishedTime     null_v4.Time   `db:"published_time" json:"publishedTime"`
	MediaType         null_v4.String `db:"media_type" json:"mediaType"`
	PrimaryGroupID    null_v4.Int    `db:"primary_group_id" json:"primaryGroupID"`
	SubclippedMediaID null_v4.Int    `db:"subclipped_media_id" json:"subclippedMediaID"`
	ReferenceMediaID  null_v4.Int    `db:"reference_media_id" json:"referenceMediaID"`
	SequenceNumber    int16          `db:"sequence_number" json:"sequenceNumber"`
	StartTime         null_v4.Float  `db:"start_time" json:"startTime"`
	EndTime           null_v4.Float  `db:"end_time" json:"endTime"`
	AssetID           null_v4.Int    `db:"asset_id" json:"assetID"`
	Agerating         null_v4.String `db:"agerating" json:"agerating"`
	Title             null_v4.String `db:"title" json:"title"`
	Description       null_v4.String `db:"description" json:"description"`
	LongDescription   null_v4.String `db:"long_description" json:"longDescription"`
	Status            int16          `db:"status" json:"status"`
}

type InsertMediaRow struct {
	Status            int16          `db:"status" json:"status"`
	Type              null_v4.String `db:"type" json:"type"`
	PublishedTime     null_v4.Time   `db:"published_time" json:"publishedTime"`
	AvailableFrom     null_v4.Time   `db:"available_from" json:"availableFrom"`
	AvailableTo       null_v4.Time   `db:"available_to" json:"availableTo"`
	Title             null_v4.String `db:"title" json:"title"`
	Description       null_v4.String `db:"description" json:"description"`
	LongDescription   null_v4.String `db:"long_description" json:"longDescription"`
	ImageID           null_v4.Int    `db:"image_id" json:"imageID"`
	TranslationID     int64          `db:"translation_id" json:"translationID"`
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
}

func (q *Queries) InsertMedia(ctx context.Context, arg InsertMediaParams) (InsertMediaRow, error) {
	row := q.db.QueryRowContext(ctx, insertMedia,
		arg.AvailableFrom,
		arg.AvailableTo,
		arg.PublishedTime,
		arg.MediaType,
		arg.PrimaryGroupID,
		arg.SubclippedMediaID,
		arg.ReferenceMediaID,
		arg.SequenceNumber,
		arg.StartTime,
		arg.EndTime,
		arg.AssetID,
		arg.Agerating,
		arg.Title,
		arg.Description,
		arg.LongDescription,
		arg.Status,
	)
	var i InsertMediaRow
	err := row.Scan(
		&i.Status,
		&i.Type,
		&i.PublishedTime,
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
	)
	return i, err
}

const upsertMedia = `-- name: UpsertMedia :one
WITH c AS (
    UPDATE collectable c1
    SET
        available_from = $2,
        available_to = $3,
        published_time = $4,
        status = $5
    WHERE c1.id = $6
    RETURNING c1.id, c1.type, c1.available_from, c1.available_to, c1.status, c1.created_at, c1.updated_at, c1.published_time
),
m AS (
    UPDATE media m1
    SET
        media_type = $7,
        primary_group_id = $8,
        subclipped_media_id = $9,
        reference_media_id = $10,
        sequence_number = $11,
        start_time = $12,
        end_time = $13,
        asset_id = $14,
        agerating = $15
    WHERE m1.id = $6
    RETURNING m1.id, m1.collectable_type, m1.media_type, m1.primary_group_id, m1.subclipped_media_id, m1.reference_media_id, m1.sequence_number, m1.start_time, m1.end_time, m1.asset_id, m1.agerating, m1.created_at, m1.updated_at
),
t AS (
    INSERT INTO media_t (
        media_id,
        language_code,
        title,
        description,
        long_description
    ) SELECT 
        $6,
        'no',
        $16,
        $17,
        $18
    ON CONFLICT (media_id, language_code)
        DO UPDATE SET
            title = $16,
            description = $17,
            long_description = $18
    RETURNING id, media_id, language_code, title, description, long_description, image_id
),
ug_del AS (
    DELETE FROM usergroup_collectable
    where collectable_id=$6 and usergroup_id not in (select ug.id from unnest($1::text[]) as ug(id))
),
ug_ins AS (
    INSERT INTO usergroup_collectable(usergroup_id, collectable_id)
    select ug.id, $6
    from unnest($1::text[]) as ug(id)
    ON CONFLICT (usergroup_id, collectable_id) DO NOTHING
)
SELECT 
    m.id, m.collectable_type, m.media_type, m.primary_group_id, m.subclipped_media_id, m.reference_media_id, m.sequence_number, m.start_time, m.end_time, m.asset_id, m.agerating, m.created_at, m.updated_at,
    c.id, c.type, c.available_from, c.available_to, c.status, c.created_at, c.updated_at, c.published_time,
    $1::text[] as usergroups,
	t.title,
	t.description,
	t.long_description,
	t.image_id
    FROM c,t,m
`

type UpsertMediaParams struct {
	Usergroups        []string       `db:"usergroups" json:"usergroups"`
	AvailableFrom     null_v4.Time   `db:"available_from" json:"availableFrom"`
	AvailableTo       null_v4.Time   `db:"available_to" json:"availableTo"`
	PublishedTime     null_v4.Time   `db:"published_time" json:"publishedTime"`
	Status            int16          `db:"status" json:"status"`
	ID                int64          `db:"id" json:"id"`
	MediaType         null_v4.String `db:"media_type" json:"mediaType"`
	PrimaryGroupID    null_v4.Int    `db:"primary_group_id" json:"primaryGroupID"`
	SubclippedMediaID null_v4.Int    `db:"subclipped_media_id" json:"subclippedMediaID"`
	ReferenceMediaID  null_v4.Int    `db:"reference_media_id" json:"referenceMediaID"`
	SequenceNumber    int16          `db:"sequence_number" json:"sequenceNumber"`
	StartTime         null_v4.Float  `db:"start_time" json:"startTime"`
	EndTime           null_v4.Float  `db:"end_time" json:"endTime"`
	AssetID           null_v4.Int    `db:"asset_id" json:"assetID"`
	Agerating         null_v4.String `db:"agerating" json:"agerating"`
	Title             null_v4.String `db:"title" json:"title"`
	Description       null_v4.String `db:"description" json:"description"`
	LongDescription   null_v4.String `db:"long_description" json:"longDescription"`
}

type UpsertMediaRow struct {
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
	ID_2              int64          `db:"id_2" json:"id2"`
	Type              null_v4.String `db:"type" json:"type"`
	AvailableFrom     null_v4.Time   `db:"available_from" json:"availableFrom"`
	AvailableTo       null_v4.Time   `db:"available_to" json:"availableTo"`
	Status            int16          `db:"status" json:"status"`
	CreatedAt_2       time.Time      `db:"created_at_2" json:"createdAt2"`
	UpdatedAt_2       time.Time      `db:"updated_at_2" json:"updatedAt2"`
	PublishedTime     null_v4.Time   `db:"published_time" json:"publishedTime"`
	Usergroups        []string       `db:"usergroups" json:"usergroups"`
	Title             null_v4.String `db:"title" json:"title"`
	Description       null_v4.String `db:"description" json:"description"`
	LongDescription   null_v4.String `db:"long_description" json:"longDescription"`
	ImageID           null_v4.Int    `db:"image_id" json:"imageID"`
}

func (q *Queries) UpsertMedia(ctx context.Context, arg UpsertMediaParams) (UpsertMediaRow, error) {
	row := q.db.QueryRowContext(ctx, upsertMedia,
		pq.Array(arg.Usergroups),
		arg.AvailableFrom,
		arg.AvailableTo,
		arg.PublishedTime,
		arg.Status,
		arg.ID,
		arg.MediaType,
		arg.PrimaryGroupID,
		arg.SubclippedMediaID,
		arg.ReferenceMediaID,
		arg.SequenceNumber,
		arg.StartTime,
		arg.EndTime,
		arg.AssetID,
		arg.Agerating,
		arg.Title,
		arg.Description,
		arg.LongDescription,
	)
	var i UpsertMediaRow
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
		&i.ID_2,
		&i.Type,
		&i.AvailableFrom,
		&i.AvailableTo,
		&i.Status,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.PublishedTime,
		pq.Array(&i.Usergroups),
		&i.Title,
		&i.Description,
		&i.LongDescription,
		&i.ImageID,
	)
	return i, err
}
