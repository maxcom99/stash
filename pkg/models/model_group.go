package models

import (
	"context"
	"time"
)

type Group struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Aliases  string `json:"aliases"`
	Duration *int   `json:"duration"`
	Date     *Date  `json:"date"`
	// Rating expressed in 1-100 scale
	Rating    *int      `json:"rating"`
	StudioID  *int      `json:"studio_id"`
	Director  string    `json:"director"`
	Synopsis  string    `json:"synopsis"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	URLs   RelatedStrings `json:"urls"`
	TagIDs RelatedIDs     `json:"tag_ids"`
}

func NewGroup() Group {
	currentTime := time.Now()
	return Group{
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

func (m *Group) LoadURLs(ctx context.Context, l URLLoader) error {
	return m.URLs.load(func() ([]string, error) {
		return l.GetURLs(ctx, m.ID)
	})
}

func (m *Group) LoadTagIDs(ctx context.Context, l TagIDLoader) error {
	return m.TagIDs.load(func() ([]int, error) {
		return l.GetTagIDs(ctx, m.ID)
	})
}

type GroupPartial struct {
	Name     OptionalString
	Aliases  OptionalString
	Duration OptionalInt
	Date     OptionalDate
	// Rating expressed in 1-100 scale
	Rating    OptionalInt
	StudioID  OptionalInt
	Director  OptionalString
	Synopsis  OptionalString
	URLs      *UpdateStrings
	TagIDs    *UpdateIDs
	CreatedAt OptionalTime
	UpdatedAt OptionalTime
}

func NewGroupPartial() GroupPartial {
	currentTime := time.Now()
	return GroupPartial{
		UpdatedAt: NewOptionalTime(currentTime),
	}
}
