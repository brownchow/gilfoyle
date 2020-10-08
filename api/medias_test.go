package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dreamvo/gilfoyle/api/db"
	"github.com/dreamvo/gilfoyle/ent"
	"github.com/dreamvo/gilfoyle/ent/enttest"
	"github.com/dreamvo/gilfoyle/ent/schema"
	"github.com/dreamvo/gilfoyle/httputils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	assertTest "github.com/stretchr/testify/assert"
	"testing"
)

func TestMedias(t *testing.T) {
	assert := assertTest.New(t)
	r = gin.Default()
	r = RegisterRoutes(r, RouterOptions{
		ExposeSwaggerUI: false,
	})

	t.Run("GET /medias", func(t *testing.T) {
		t.Run("should return empty array", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			res, err := performRequest(r, "GET", "/medias", nil)
			assert.NoError(err, "should be equal")

			var body struct {
				Code int         `json:"code"`
				Data []ent.Media `json:"data,omitempty"`
			}
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(res.Result().StatusCode, 200, "should be equal")
			assert.Equal(200, body.Code)
			assert.Equal([]ent.Media{}, body.Data)
		})

		t.Run("should return latest medias", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			for i := 0; i < 5; i++ {
				_, _ = db.Client.Media.
					Create().
					SetTitle(fmt.Sprintf("%d", i)).
					SetStatus(schema.MediaStatusProcessing).
					Save(context.Background())
			}

			res, err := performRequest(r, "GET", "/medias", nil)
			assert.NoError(err, "should be equal")

			var body struct {
				Code int         `json:"code"`
				Data []ent.Media `json:"data,omitempty"`
			}
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(res.Result().StatusCode, 200, "should be equal")
			assert.Equal(200, body.Code)
			assert.Equal(5, len(body.Data))
		})

		t.Run("should limit results to 2", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			for i := 0; i < 3; i++ {
				_, _ = db.Client.Media.
					Create().
					SetTitle(fmt.Sprintf("%d", i)).
					SetStatus(schema.MediaStatusProcessing).
					Save(context.Background())
			}

			res, err := performRequest(r, "GET", "/medias?limit=2", nil)
			assert.NoError(err, "should be equal")

			var body struct {
				Code int         `json:"code"`
				Data []ent.Media `json:"data,omitempty"`
			}
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(res.Result().StatusCode, 200, "should be equal")
			assert.Equal(200, body.Code)
			assert.Equal(2, len(body.Data))
		})

		t.Run("should return results with offset 1", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			v, _ := db.Client.Media.
				Create().
				SetTitle("video1").
				SetStatus(schema.MediaStatusProcessing).
				Save(context.Background())

			_, _ = db.Client.Media.
				Create().
				SetTitle("video2").
				SetStatus(schema.MediaStatusProcessing).
				Save(context.Background())

			res, err := performRequest(r, "GET", "/medias?offset=1", nil)
			assert.NoError(err, "should be equal")

			var body struct {
				Code int         `json:"code"`
				Data []ent.Media `json:"data,omitempty"`
			}
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(res.Result().StatusCode, 200, "should be equal")
			assert.Equal(200, body.Code)
			assert.Equal(1, len(body.Data))
			assert.Equal(v.ID.String(), body.Data[0].ID.String())
		})
	})

	t.Run("GET /medias/{id}", func(t *testing.T) {
		t.Run("should return error for invalid UUID", func(t *testing.T) {
			res, err := performRequest(r, "GET", "/medias/uuid", nil)
			assert.NoError(err, "should be equal")

			var body httputils.ErrorResponse
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(400, res.Result().StatusCode, "should be equal")
			assert.Equal(400, body.Code)
			assert.Equal("invalid UUID provided", body.Message)
		})

		t.Run("should return media", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			v, _ := db.Client.Media.
				Create().
				SetTitle("no u").
				SetStatus(schema.MediaStatusProcessing).
				Save(context.Background())

			res, err := performRequest(r, "GET", "/medias/"+v.ID.String(), nil)
			assert.NoError(err, "should be equal")

			var body struct {
				Code int       `json:"code"`
				Data ent.Media `json:"data,omitempty"`
			}
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(200, res.Result().StatusCode, "should be equal")
			assert.Equal(200, body.Code)
			assert.Equal(v.Title, body.Data.Title)
		})
	})

	t.Run("DELETE /medias/{id}", func(t *testing.T) {
		t.Run("should delete newly created media", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			v, _ := db.Client.Media.
				Create().
				SetTitle("test").
				SetStatus(schema.MediaStatusProcessing).
				Save(context.Background())

			res, err := performRequest(r, "DELETE", "/medias/"+v.ID.String(), nil)
			assert.NoError(err, "should be equal")

			assert.Equal(res.Result().StatusCode, 200, "should be equal")

			res, err = performRequest(r, "DELETE", "/medias/"+v.ID.String(), nil)
			assert.NoError(err, "should be equal")

			var body httputils.ErrorResponse
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(404, res.Code)
			assert.Equal("resource not found", body.Message)
		})

		t.Run("should return error on invalid uid", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			res, err := performRequest(r, "DELETE", "/medias/uuid", nil)
			assert.NoError(err, "should be equal")

			var body httputils.ErrorResponse
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(res.Result().StatusCode, 400, "should be equal")
			assert.Equal(400, body.Code)
			assert.Equal("invalid UUID provided", body.Message)
		})
	})

	t.Run("POST /medias", func(t *testing.T) {
		t.Run("should create a new media", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			res, err := performRequest(r, "POST", "/medias", CreateMedia{
				Title: "test",
			})
			assert.NoError(err)

			var body httputils.DataResponse
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(200, res.Result().StatusCode)
			assert.Equal("test", body.Data.(map[string]interface{})["title"])
			assert.Equal("processing", body.Data.(map[string]interface{})["status"])
		})

		t.Run("should return validation error (1)", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			res, err := performRequest(r, "POST", "/medias", CreateMedia{
				Title: "",
			})
			assert.NoError(err, "should be equal")

			var body httputils.ValidationErrorResponse
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(400, res.Result().StatusCode, "should be equal")
			assert.Equal("Some parameters are missing or invalid", body.Message)
			assert.Equal(map[string]string{
				"title": "Key: 'CreateMedia.Title' Error:Field validation for 'Title' failed on the 'gte' tag",
			}, body.Fields)
		})

		t.Run("should return validation error (2)", func(t *testing.T) {
			db.Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer db.Client.Close()

			res, err := performRequest(r, "POST", "/medias", nil)
			assert.NoError(err, "should be equal")

			var body httputils.ValidationErrorResponse
			_ = json.NewDecoder(res.Body).Decode(&body)

			assert.Equal(400, res.Result().StatusCode, "should be equal")
			assert.Equal("Some parameters are missing or invalid", body.Message)
			assert.Equal(map[string]string{
				"title": "Key: 'CreateMedia.Title' Error:Field validation for 'Title' failed on the 'required' tag",
			}, body.Fields)
		})
	})

	t.Run("PATCH /medias/{id}", func(t *testing.T) {})
}
