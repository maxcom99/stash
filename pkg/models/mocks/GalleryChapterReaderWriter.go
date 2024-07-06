// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/stashapp/stash/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// GalleryChapterReaderWriter is an autogenerated mock type for the GalleryChapterReaderWriter type
type GalleryChapterReaderWriter struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, newGalleryChapter
func (_m *GalleryChapterReaderWriter) Create(ctx context.Context, newGalleryChapter *models.GalleryChapter) error {
	ret := _m.Called(ctx, newGalleryChapter)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GalleryChapter) error); ok {
		r0 = rf(ctx, newGalleryChapter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Destroy provides a mock function with given fields: ctx, id
func (_m *GalleryChapterReaderWriter) Destroy(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, id
func (_m *GalleryChapterReaderWriter) Find(ctx context.Context, id int) (*models.GalleryChapter, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.GalleryChapter
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.GalleryChapter); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GalleryChapter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByGalleryID provides a mock function with given fields: ctx, galleryID
func (_m *GalleryChapterReaderWriter) FindByGalleryID(ctx context.Context, galleryID int) ([]*models.GalleryChapter, error) {
	ret := _m.Called(ctx, galleryID)

	var r0 []*models.GalleryChapter
	if rf, ok := ret.Get(0).(func(context.Context, int) []*models.GalleryChapter); ok {
		r0 = rf(ctx, galleryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.GalleryChapter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, galleryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMany provides a mock function with given fields: ctx, ids
func (_m *GalleryChapterReaderWriter) FindMany(ctx context.Context, ids []int) ([]*models.GalleryChapter, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*models.GalleryChapter
	if rf, ok := ret.Get(0).(func(context.Context, []int) []*models.GalleryChapter); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.GalleryChapter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []int) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, updatedGalleryChapter
func (_m *GalleryChapterReaderWriter) Update(ctx context.Context, updatedGalleryChapter *models.GalleryChapter) error {
	ret := _m.Called(ctx, updatedGalleryChapter)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GalleryChapter) error); ok {
		r0 = rf(ctx, updatedGalleryChapter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePartial provides a mock function with given fields: ctx, id, updatedGalleryChapter
func (_m *GalleryChapterReaderWriter) UpdatePartial(ctx context.Context, id int, updatedGalleryChapter models.GalleryChapterPartial) (*models.GalleryChapter, error) {
	ret := _m.Called(ctx, id, updatedGalleryChapter)

	var r0 *models.GalleryChapter
	if rf, ok := ret.Get(0).(func(context.Context, int, models.GalleryChapterPartial) *models.GalleryChapter); ok {
		r0 = rf(ctx, id, updatedGalleryChapter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GalleryChapter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, models.GalleryChapterPartial) error); ok {
		r1 = rf(ctx, id, updatedGalleryChapter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}