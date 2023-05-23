// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/stashapp/stash/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ScrapedItemReaderWriter is an autogenerated mock type for the ScrapedItemReaderWriter type
type ScrapedItemReaderWriter struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *ScrapedItemReaderWriter) All(ctx context.Context) ([]*models.ScrapedItem, error) {
	ret := _m.Called(ctx)

	var r0 []*models.ScrapedItem
	if rf, ok := ret.Get(0).(func(context.Context) []*models.ScrapedItem); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.ScrapedItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, newObject
func (_m *ScrapedItemReaderWriter) Create(ctx context.Context, newObject models.ScrapedItem) (*models.ScrapedItem, error) {
	ret := _m.Called(ctx, newObject)

	var r0 *models.ScrapedItem
	if rf, ok := ret.Get(0).(func(context.Context, models.ScrapedItem) *models.ScrapedItem); ok {
		r0 = rf(ctx, newObject)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ScrapedItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ScrapedItem) error); ok {
		r1 = rf(ctx, newObject)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
