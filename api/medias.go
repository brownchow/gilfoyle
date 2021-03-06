package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/dreamvo/gilfoyle/api/db"
	"github.com/dreamvo/gilfoyle/ent"
	_ "github.com/dreamvo/gilfoyle/ent"
	"github.com/dreamvo/gilfoyle/ent/media"
	"github.com/dreamvo/gilfoyle/ent/schema"
	"github.com/dreamvo/gilfoyle/httputils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type CreateMedia struct {
	Title string `json:"title" validate:"required,gte=1,lte=255" example:"Sheep Discovers How To Use A Trampoline"`
}

type UpdateMedia struct {
	CreateMedia
}

// @ID getAllMedias
// @Tags Medias
// @Summary Query medias
// @Description Get latest created medias
// @Produce  json
// @Success 200 {object} httputils.DataResponse{data=[]ent.Media}
// @Failure 500 {object} httputils.ErrorResponse
// @Router /medias [get]
// @Param limit query int false "Max number of results"
// @Param offset query int false "Number of results to ignore"
func getAllMedias(ctx *gin.Context) {
	limit := ctx.GetInt("limit")
	offset := ctx.GetInt("offset")

	medias, err := db.Client.Media.
		Query().
		Order(ent.Desc(media.FieldCreatedAt)).
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		httputils.NewError(ctx, http.StatusInternalServerError, errors.Unwrap(err))
		return
	}

	httputils.NewData(ctx, http.StatusOK, medias)
}

// @ID getMedia
// @Tags Medias
// @Summary Get a media
// @Description Get one media
// @Produce  json
// @Param id path string true "Media ID" validate(required)
// @Success 200 {object} httputils.DataResponse{data=ent.Media}
// @Failure 404 {object} httputils.ErrorResponse
// @Failure 500 {object} httputils.ErrorResponse
// @Router /medias/{id} [get]
func getMedia(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, fmt.Errorf(ErrInvalidUUID))
		return
	}

	v, err := db.Client.Media.Get(context.Background(), parsedUUID)
	if v == nil {
		httputils.NewError(ctx, http.StatusNotFound, errors.New(ErrResourceNotFound))
		return
	}
	if err != nil {
		httputils.NewError(ctx, http.StatusInternalServerError, errors.Unwrap(err))
		return
	}

	httputils.NewData(ctx, http.StatusOK, v)
}

// @ID deleteMedia
// @Tags Medias
// @Summary Delete a media
// @Description Delete one media
// @Produce  json
// @Param id path string true "Media ID" validate(required)
// @Success 200 {object} httputils.DataResponse
// @Failure 400 {object} httputils.ErrorResponse
// @Failure 404 {object} httputils.ErrorResponse
// @Failure 500 {object} httputils.ErrorResponse
// @Router /medias/{id} [delete]
func deleteMedia(ctx *gin.Context) {
	id := ctx.Param("id")

	parsedUUID, err := httputils.ValidateUUID(id)
	if err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, fmt.Errorf(ErrInvalidUUID))
		return
	}

	v, _ := db.Client.Media.Get(context.Background(), parsedUUID)
	if v == nil {
		httputils.NewError(ctx, http.StatusNotFound, errors.New(ErrResourceNotFound))
		return
	}

	err = db.Client.Media.DeleteOneID(parsedUUID).Exec(context.Background())
	if err != nil {
		httputils.NewError(ctx, http.StatusInternalServerError, errors.Unwrap(err))
		return
	}

	httputils.NewData(ctx, http.StatusOK, nil)
}

// @ID createMedia
// @Tags Medias
// @Summary Create a media
// @Description Create a new media
// @Accept  json
// @Produce  json
// @Success 200 {object} httputils.DataResponse{data=ent.Media}
// @Failure 400 {object} httputils.ErrorResponse
// @Failure 500 {object} httputils.ErrorResponse
// @Router /medias [post]
// @Param media body CreateMedia true "Media data" validate(required)
func createMedia(ctx *gin.Context) {
	var body CreateMedia

	if err := httputils.ValidateBody(ctx, &body); err != nil {
		httputils.NewValidationError(ctx, err)
		return
	}

	v, err := db.Client.Media.
		Create().
		SetTitle(body.Title).
		SetStatus(schema.MediaStatusProcessing).
		Save(context.Background())
	if ent.IsValidationError(err) {
		httputils.NewError(ctx, http.StatusBadRequest, errors.Unwrap(err))
		return
	}
	if err != nil {
		httputils.NewError(ctx, http.StatusInternalServerError, errors.Unwrap(err))
		return
	}

	httputils.NewData(ctx, http.StatusOK, v)
}

// @ID updateMedia
// @Tags Medias
// @Summary Update a media
// @Description Update an existing media
// @Accept  json
// @Produce  json
// @Success 200 {object} httputils.DataResponse{data=ent.Media}
// @Failure 400 {object} httputils.ErrorResponse
// @Failure 404 {object} httputils.ErrorResponse
// @Failure 500 {object} httputils.ErrorResponse
// @Router /medias/{id} [patch]
// @Param id path string true "Media ID" validate(required)
// @Param media body UpdateMedia true "Media data" validate(required)
func updateMedia(ctx *gin.Context) {
	var body CreateMedia

	if err := httputils.ValidateBody(ctx, &body); err != nil {
		httputils.NewValidationError(ctx, err)
		return
	}

	id := ctx.Param("id")

	parsedUUID, err := httputils.ValidateUUID(id)
	if err != nil {
		httputils.NewError(ctx, http.StatusBadRequest, fmt.Errorf(ErrInvalidUUID))
		return
	}

	v, _ := db.Client.Media.Get(context.Background(), parsedUUID)
	if v == nil {
		httputils.NewError(ctx, http.StatusNotFound, errors.New(ErrResourceNotFound))
		return
	}

	v, err = db.Client.Media.
		UpdateOneID(parsedUUID).
		SetTitle(body.Title).
		Save(context.Background())
	if ent.IsValidationError(err) {
		httputils.NewError(ctx, http.StatusBadRequest, errors.Unwrap(err))
		return
	}
	if err != nil {
		httputils.NewError(ctx, http.StatusInternalServerError, errors.Unwrap(err))
		return
	}

	httputils.NewData(ctx, http.StatusOK, v)
}
