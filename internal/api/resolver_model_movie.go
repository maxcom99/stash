package api

import (
	"context"
	"time"

	"github.com/stashapp/stash/internal/api/loaders"
	"github.com/stashapp/stash/internal/api/urlbuilders"
	"github.com/stashapp/stash/pkg/models"
	"github.com/stashapp/stash/pkg/utils"
)

func (r *movieResolver) Name(ctx context.Context, obj *models.Movie) (string, error) {
	if obj.Name.Valid {
		return obj.Name.String, nil
	}
	return "", nil
}

func (r *movieResolver) URL(ctx context.Context, obj *models.Movie) (*string, error) {
	if obj.URL.Valid {
		return &obj.URL.String, nil
	}
	return nil, nil
}

func (r *movieResolver) Aliases(ctx context.Context, obj *models.Movie) (*string, error) {
	if obj.Aliases.Valid {
		return &obj.Aliases.String, nil
	}
	return nil, nil
}

func (r *movieResolver) Duration(ctx context.Context, obj *models.Movie) (*int, error) {
	if obj.Duration.Valid {
		rating := int(obj.Duration.Int64)
		return &rating, nil
	}
	return nil, nil
}

func (r *movieResolver) Date(ctx context.Context, obj *models.Movie) (*string, error) {
	if obj.Date.Valid {
		result := utils.GetYMDFromDatabaseDate(obj.Date.String)
		return &result, nil
	}
	return nil, nil
}

func (r *movieResolver) Rating(ctx context.Context, obj *models.Movie) (*int, error) {
	if obj.Rating.Valid {
		rating := models.Rating100To5(int(obj.Rating.Int64))
		return &rating, nil
	}
	return nil, nil
}

func (r *movieResolver) Rating100(ctx context.Context, obj *models.Movie) (*int, error) {
	if obj.Rating.Valid {
		rating := int(obj.Rating.Int64)
		return &rating, nil
	}
	return nil, nil
}

func (r *movieResolver) Studio(ctx context.Context, obj *models.Movie) (ret *models.Studio, err error) {
	if obj.StudioID.Valid {
		return loaders.From(ctx).StudioByID.Load(int(obj.StudioID.Int64))
	}

	return nil, nil
}

func (r *movieResolver) Director(ctx context.Context, obj *models.Movie) (*string, error) {
	if obj.Director.Valid {
		return &obj.Director.String, nil
	}
	return nil, nil
}

func (r *movieResolver) Synopsis(ctx context.Context, obj *models.Movie) (*string, error) {
	if obj.Synopsis.Valid {
		return &obj.Synopsis.String, nil
	}
	return nil, nil
}

func (r *movieResolver) FrontImagePath(ctx context.Context, obj *models.Movie) (*string, error) {
	var hasImage bool
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		var err error
		hasImage, err = r.repository.Movie.HasFrontImage(ctx, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	baseURL, _ := ctx.Value(BaseURLCtxKey).(string)
	imagePath := urlbuilders.NewMovieURLBuilder(baseURL, obj).GetMovieFrontImageURL(hasImage)
	return &imagePath, nil
}

func (r *movieResolver) BackImagePath(ctx context.Context, obj *models.Movie) (*string, error) {
	var hasImage bool
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		var err error
		hasImage, err = r.repository.Movie.HasBackImage(ctx, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	// don't return anything if there is no back image
	if !hasImage {
		return nil, nil
	}

	baseURL, _ := ctx.Value(BaseURLCtxKey).(string)
	imagePath := urlbuilders.NewMovieURLBuilder(baseURL, obj).GetMovieBackImageURL()
	return &imagePath, nil
}

func (r *movieResolver) SceneCount(ctx context.Context, obj *models.Movie) (ret *int, err error) {
	var res int
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		res, err = r.repository.Scene.CountByMovieID(ctx, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return &res, err
}

func (r *movieResolver) Scenes(ctx context.Context, obj *models.Movie) (ret []*models.Scene, err error) {
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		var err error
		ret, err = r.repository.Scene.FindByMovieID(ctx, obj.ID)
		return err
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func (r *movieResolver) CreatedAt(ctx context.Context, obj *models.Movie) (*time.Time, error) {
	return &obj.CreatedAt.Timestamp, nil
}

func (r *movieResolver) UpdatedAt(ctx context.Context, obj *models.Movie) (*time.Time, error) {
	return &obj.UpdatedAt.Timestamp, nil
}
