// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
	"github.com/tabbed/pqtype"
	null_v4 "gopkg.in/guregu/null.v4"
)

type AdminCollectable struct {
	ID            int64          `db:"id" json:"id"`
	Type          null_v4.String `db:"type" json:"type"`
	AvailableFrom null_v4.Time   `db:"available_from" json:"availableFrom"`
	AvailableTo   null_v4.Time   `db:"available_to" json:"availableTo"`
	Status        int16          `db:"status" json:"status"`
	CreatedAt     time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updatedAt"`
	PublishedTime null_v4.Time   `db:"published_time" json:"publishedTime"`
	Tags          pq.Int64Array  `db:"tags" json:"tags"`
	Usergroups    pq.StringArray `db:"usergroups" json:"usergroups"`
}

type AdminMedia struct {
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
}

type AdminPage struct {
	Status          int16          `db:"status" json:"status"`
	Type            null_v4.String `db:"type" json:"type"`
	AvailableFrom   null_v4.Time   `db:"available_from" json:"availableFrom"`
	AvailableTo     null_v4.Time   `db:"available_to" json:"availableTo"`
	Title           null_v4.String `db:"title" json:"title"`
	Description     null_v4.String `db:"description" json:"description"`
	TranslationID   null_v4.Int    `db:"translation_id" json:"translationID"`
	ID              int64          `db:"id" json:"id"`
	CollectableType null_v4.String `db:"collectable_type" json:"collectableType"`
	PublishedTime   null_v4.Time   `db:"published_time" json:"publishedTime"`
	Usergroups      pq.StringArray `db:"usergroups" json:"usergroups"`
	Tags            pq.Int64Array  `db:"tags" json:"tags"`
}

type AdminTag struct {
	ID            int64          `db:"id" json:"id"`
	Type          null_v4.String `db:"type" json:"type"`
	TranslationID null_v4.Int    `db:"translation_id" json:"translationID"`
	LanguageCode  null_v4.String `db:"language_code" json:"languageCode"`
	Title         null_v4.String `db:"title" json:"title"`
}

type Agerating struct {
	Code      string         `db:"code" json:"code"`
	Title     null_v4.String `db:"title" json:"title"`
	SortOrder int16          `db:"sort_order" json:"sortOrder"`
}

type AgeratingT struct {
	ID            int32          `db:"id" json:"id"`
	AgeratingCode null_v4.String `db:"agerating_code" json:"ageratingCode"`
	LanguageCode  string         `db:"language_code" json:"languageCode"`
}

type Asset struct {
	ID                 int64          `db:"id" json:"id"`
	SourceID           null_v4.String `db:"source_id" json:"sourceID"`
	PublishedVersionID null_v4.Int    `db:"published_version_id" json:"publishedVersionID"`
	Name               null_v4.String `db:"name" json:"name"`
}

type AssetFormat struct {
	ID      int64          `db:"id" json:"id"`
	Name    null_v4.String `db:"name" json:"name"`
	Url     null_v4.String `db:"url" json:"url"`
	Bitrate null_v4.Int    `db:"bitrate" json:"bitrate"`
}

type AssetVersion struct {
	ID      int64       `db:"id" json:"id"`
	AssetID null_v4.Int `db:"asset_id" json:"assetID"`
}

type AssetVersionFormat struct {
	ID             int64       `db:"id" json:"id"`
	AssetversionID null_v4.Int `db:"assetversion_id" json:"assetversionID"`
	AssetformatID  null_v4.Int `db:"assetformat_id" json:"assetformatID"`
}

type CalendarEvent struct {
	ID               int64          `db:"id" json:"id"`
	EventID          null_v4.String `db:"event_id" json:"eventID"`
	Start            null_v4.Time   `db:"start" json:"start"`
	End              null_v4.Time   `db:"end" json:"end"`
	Locked           bool           `db:"locked" json:"locked"`
	UseImageFromBcco bool           `db:"use_image_from_bcco" json:"useImageFromBcco"`
	ImageID          null_v4.Int    `db:"image_id" json:"imageID"`
	AudienceID       null_v4.Int    `db:"audience_id" json:"audienceID"`
}

type CalendareventT struct {
	ID              int64          `db:"id" json:"id"`
	CalendareventID int64          `db:"calendarevent_id" json:"calendareventID"`
	LanguageCode    string         `db:"language_code" json:"languageCode"`
	Name            null_v4.String `db:"name" json:"name"`
	ImageID         null_v4.Int    `db:"image_id" json:"imageID"`
}

type Collectable struct {
	ID            int64          `db:"id" json:"id"`
	Type          null_v4.String `db:"type" json:"type"`
	AvailableFrom null_v4.Time   `db:"available_from" json:"availableFrom"`
	AvailableTo   null_v4.Time   `db:"available_to" json:"availableTo"`
	Status        int16          `db:"status" json:"status"`
	CreatedAt     time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updatedAt"`
	PublishedTime null_v4.Time   `db:"published_time" json:"publishedTime"`
}

type Collection struct {
	ID          int64                 `db:"id" json:"id"`
	Type        null_v4.String        `db:"type" json:"type"`
	QueryID     null_v4.Int           `db:"query_id" json:"queryID"`
	QueryParams pqtype.NullRawMessage `db:"query_params" json:"queryParams"`
	PageID      null_v4.Int           `db:"page_id" json:"pageID"`
}

type CollectionCollectable struct {
	ID            int64       `db:"id" json:"id"`
	CollectionID  null_v4.Int `db:"collection_id" json:"collectionID"`
	CollectableID null_v4.Int `db:"collectable_id" json:"collectableID"`
	SortOrder     int16       `db:"sort_order" json:"sortOrder"`
}

type Faq struct {
	ID       int16 `db:"id" json:"id"`
	Category int16 `db:"category" json:"category"`
}

type FaqT struct {
	ID           int16          `db:"id" json:"id"`
	FaqID        int16          `db:"faq_id" json:"faqID"`
	LanguageCode string         `db:"language_code" json:"languageCode"`
	Question     null_v4.String `db:"question" json:"question"`
	Answer       null_v4.String `db:"answer" json:"answer"`
}

type Faqcategory struct {
	ID        int16 `db:"id" json:"id"`
	SortOrder int16 `db:"sort_order" json:"sortOrder"`
	Status    int16 `db:"status" json:"status"`
}

type FaqcategoryT struct {
	ID            int16  `db:"id" json:"id"`
	FaqcategoryID int16  `db:"faqcategory_id" json:"faqcategoryID"`
	LanguageCode  string `db:"language_code" json:"languageCode"`
	Title         int16  `db:"title" json:"title"`
}

type Image struct {
	ID  int64  `db:"id" json:"id"`
	Url string `db:"url" json:"url"`
}

type Language struct {
	Code string `db:"code" json:"code"`
	Name string `db:"name" json:"name"`
}

type Media struct {
	ID                int64          `db:"id" json:"id"`
	CollectableType   null_v4.String `db:"collectable_type" json:"collectableType"`
	MediaType         null_v4.String `db:"media_type" json:"mediaType"`
	PrimaryGroupID    null_v4.Int    `db:"primary_group_id" json:"primaryGroupID"`
	SubclippedMediaID null_v4.Int    `db:"subclipped_media_id" json:"subclippedMediaID"`
	ReferenceMediaID  null_v4.Int    `db:"reference_media_id" json:"referenceMediaID"`
	SequenceNumber    int16          `db:"sequence_number" json:"sequenceNumber"`
	// for subclips and markers
	StartTime null_v4.Float `db:"start_time" json:"startTime"`
	// for subclips
	EndTime   null_v4.Float  `db:"end_time" json:"endTime"`
	AssetID   null_v4.Int    `db:"asset_id" json:"assetID"`
	Agerating null_v4.String `db:"agerating" json:"agerating"`
	CreatedAt time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time      `db:"updated_at" json:"updatedAt"`
}

type MediaT struct {
	ID              int64          `db:"id" json:"id"`
	MediaID         null_v4.Int    `db:"media_id" json:"mediaID"`
	LanguageCode    string         `db:"language_code" json:"languageCode"`
	Title           null_v4.String `db:"title" json:"title"`
	Description     null_v4.String `db:"description" json:"description"`
	LongDescription null_v4.String `db:"long_description" json:"longDescription"`
	ImageID         null_v4.Int    `db:"image_id" json:"imageID"`
}

type Notification struct {
	ID            int64          `db:"id" json:"id"`
	CohortID      string         `db:"cohort_id" json:"cohortID"`
	Action        null_v4.String `db:"action" json:"action"`
	InternalUrl   null_v4.String `db:"internal_url" json:"internalUrl"`
	ExternalUrl   null_v4.String `db:"external_url" json:"externalUrl"`
	Badge         int16          `db:"badge" json:"badge"`
	ScheduledTime null_v4.Time   `db:"scheduled_time" json:"scheduledTime"`
}

type NotificationAction struct {
	Code string         `db:"code" json:"code"`
	Name null_v4.String `db:"name" json:"name"`
}

type NotificationT struct {
	ID             int64          `db:"id" json:"id"`
	NotificationID null_v4.Int    `db:"notification_id" json:"notificationID"`
	LanguageCode   string         `db:"language_code" json:"languageCode"`
	Title          null_v4.String `db:"title" json:"title"`
	Message        null_v4.String `db:"message" json:"message"`
}

type Page struct {
	ID              int64          `db:"id" json:"id"`
	CollectableType null_v4.String `db:"collectable_type" json:"collectableType"`
}

type PageSection struct {
	ID        int64       `db:"id" json:"id"`
	PageID    null_v4.Int `db:"page_id" json:"pageID"`
	SectionID null_v4.Int `db:"section_id" json:"sectionID"`
}

type PageT struct {
	ID           int64          `db:"id" json:"id"`
	LanguageCode string         `db:"language_code" json:"languageCode"`
	PageID       null_v4.Int    `db:"page_id" json:"pageID"`
	Title        null_v4.String `db:"title" json:"title"`
	Description  null_v4.String `db:"description" json:"description"`
}

type Query struct {
	ID            int64           `db:"id" json:"id"`
	Template      json.RawMessage `db:"template" json:"template"`
	Parameters    json.RawMessage `db:"parameters" json:"parameters"`
	SystemDefined bool            `db:"system_defined" json:"systemDefined"`
}

type Section struct {
	ID              int64                 `db:"id" json:"id"`
	Type            string                `db:"type" json:"type"`
	CollectionID    null_v4.Int           `db:"collection_id" json:"collectionID"`
	VisibilityRules pqtype.NullRawMessage `db:"visibility_rules" json:"visibilityRules"`
}

type SectionItem struct {
	ID            int64          `db:"id" json:"id"`
	SectionID     int64          `db:"section_id" json:"sectionID"`
	Type          string         `db:"type" json:"type"`
	CollectableID null_v4.Int    `db:"collectable_id" json:"collectableID"`
	Url           null_v4.String `db:"url" json:"url"`
	CollectionID  null_v4.Int    `db:"collection_id" json:"collectionID"`
	SortOrder     int16          `db:"sort_order" json:"sortOrder"`
}

type SectionItemT struct {
	ID            int64          `db:"id" json:"id"`
	LanguageCode  string         `db:"language_code" json:"languageCode"`
	SectionItemID null_v4.Int    `db:"section_item_id" json:"sectionItemID"`
	Title         null_v4.String `db:"title" json:"title"`
}

type SectionT struct {
	ID           int64          `db:"id" json:"id"`
	SectionID    null_v4.Int    `db:"section_id" json:"sectionID"`
	LanguageCode string         `db:"language_code" json:"languageCode"`
	Title        null_v4.String `db:"title" json:"title"`
}

type SectionitemUsergroup struct {
	ID            int64          `db:"id" json:"id"`
	SectionitemID null_v4.Int    `db:"sectionitem_id" json:"sectionitemID"`
	UsergroupID   null_v4.String `db:"usergroup_id" json:"usergroupID"`
}

type Systemdatum struct {
	ID                  int16       `db:"id" json:"id"`
	Live                bool        `db:"live" json:"live"`
	FullMaintenanceMode bool        `db:"full_maintenance_mode" json:"fullMaintenanceMode"`
	MetaImageID         null_v4.Int `db:"meta_image_id" json:"metaImageID"`
}

type Tag struct {
	ID int64 `db:"id" json:"id"`
	// speaker, composer, etc.
	Type null_v4.String `db:"type" json:"type"`
}

type TagCollectable struct {
	ID            int64       `db:"id" json:"id"`
	CollectableID null_v4.Int `db:"collectable_id" json:"collectableID"`
	TagID         null_v4.Int `db:"tag_id" json:"tagID"`
}

type TagT struct {
	ID           int64          `db:"id" json:"id"`
	TagID        null_v4.Int    `db:"tag_id" json:"tagID"`
	LanguageCode string         `db:"language_code" json:"languageCode"`
	Title        null_v4.String `db:"title" json:"title"`
}

type Tvguideentry struct {
	ID                int64        `db:"id" json:"id"`
	Start             null_v4.Time `db:"start" json:"start"`
	End               null_v4.Time `db:"end" json:"end"`
	EventID           null_v4.Int  `db:"event_id" json:"eventID"`
	MediaID           null_v4.Int  `db:"media_id" json:"mediaID"`
	ImageID           null_v4.Int  `db:"image_id" json:"imageID"`
	UseImageFromMedia bool         `db:"use_image_from_media" json:"useImageFromMedia"`
}

type TvguideentryT struct {
	ID             int64          `db:"id" json:"id"`
	TvguideentryID null_v4.Int    `db:"tvguideentry_id" json:"tvguideentryID"`
	LanguageCode   string         `db:"language_code" json:"languageCode"`
	ImageID        null_v4.Int    `db:"image_id" json:"imageID"`
	Title          null_v4.String `db:"title" json:"title"`
	Description    null_v4.String `db:"description" json:"description"`
}

type UserDataDevice struct {
	ID        int64          `db:"id" json:"id"`
	UserID    string         `db:"user_id" json:"userID"`
	PushToken null_v4.String `db:"push_token" json:"pushToken"`
	App       string         `db:"app" json:"app"`
}

type UserDataList struct {
	ID     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	UserID string `db:"user_id" json:"userID"`
}

type UserDataListMedium struct {
	ID        int64       `db:"id" json:"id"`
	MediaID   int64       `db:"media_id" json:"mediaID"`
	ListID    null_v4.Int `db:"list_id" json:"listID"`
	SortOrder int16       `db:"sort_order" json:"sortOrder"`
}

type UserDataMediaProgress struct {
	ID          int64          `db:"id" json:"id"`
	MediaID     null_v4.Int    `db:"media_id" json:"mediaID"`
	UserID      null_v4.String `db:"user_id" json:"userID"`
	Seconds     null_v4.Int    `db:"seconds" json:"seconds"`
	LastUpdated null_v4.Time   `db:"last_updated" json:"lastUpdated"`
}

type UserDataUserList struct {
	ID        int64       `db:"id" json:"id"`
	UserID    string      `db:"user_id" json:"userID"`
	ListID    null_v4.Int `db:"list_id" json:"listID"`
	SortOrder int16       `db:"sort_order" json:"sortOrder"`
}

type UserDataUsergroupUser struct {
	ID          int64  `db:"id" json:"id"`
	UsergroupID string `db:"usergroup_id" json:"usergroupID"`
	UserID      string `db:"user_id" json:"userID"`
}

type Usergroup struct {
	ID string `db:"id" json:"id"`
}

type UsergroupCollectable struct {
	ID            int64  `db:"id" json:"id"`
	UsergroupID   string `db:"usergroup_id" json:"usergroupID"`
	CollectableID int64  `db:"collectable_id" json:"collectableID"`
}
