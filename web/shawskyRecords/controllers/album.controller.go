package controllers

import (
	"net/http"
	"shawskyRecords/models"
	"shawskyRecords/services"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetAlbums godoc
//
//	@Summary		Retrieve all albums
//	@Tags			Albums
//	@Produce		json
//	@Success		200	{array}		models.Album
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/albums [get]
//
// @Param X-ApplicationId header string true "Application  ID"
func (c *Controller) GetAlbums(ctx *gin.Context) {
	service := services.NewAlbumService()
	ctx.IndentedJSON(http.StatusOK, service.GetAlbums())
}

// GetAlbumById godoc
//
//	@Summary		Retrieve a single album
//	@ID Retrieve an album by id
//	@Tags			Albums
//	@Produce		json
//	@Param id path string true "album id"
//	@Success		200	{object}	models.Album
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/albums/{id} [get]
//
// @Param X-ApplicationId header string true "Application  ID"
func (c *Controller) GetAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	service := services.NewAlbumService()
	result, err := service.GetAlbumById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, result)
}

// CreateAlbum godoc
//
//	@Summary		Create a single album
//	@Tags			Albums
//	@Produce		json
//	@Accept			json
//	@Param			album	body	models.Album	true	"Add album"
//	@Success		200	{object}	models.Album
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/albums [post]
//
// @Param X-ApplicationId header string true "Application  ID"
func (c *Controller) CreateAlbum(ctx *gin.Context) {
	service := services.NewAlbumService()
	var newAlbum models.Album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid album body"})
		return
	}
	result, err := service.CreateAlbum(newAlbum)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Duplicate albums cannot be created"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, result)
}

// UpdateAlbum godoc
//
//	@Summary		Update a single album
//	@Tags			Albums
//	@Produce		json
//	@Accept			json
//	@Param id path string true "album id"
//	@Param			album	body	models.Album	true	"Update album"
//	@Success		200	{object}	models.Album
//	@Failure		400	{object}	httputil.HTTPError
//	@Router			/albums/{id} [put]
//
// @Param X-ApplicationId header string true "Application  ID"
func (c *Controller) UpdateAlbum(ctx *gin.Context) {
	id := ctx.Param("id")
	service := services.NewAlbumService()

	var updatedAlbum models.Album

	if err := ctx.BindJSON(&updatedAlbum); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid album body"})
		return
	}

	if !strings.EqualFold(id, updatedAlbum.Id) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request. Album id in route must match album in request body"})
		return
	}

	result, err := service.UpdateAlbum(updatedAlbum)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Albums could not be updated. Please verify that album exists."})
		return
	}
	ctx.IndentedJSON(http.StatusOK, result)
}

// DeleteAlbum godoc
//
//	@Summary		Delete an album
//	@Tags			Albums
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Album ID"	Format(string)
//	@Success		204
//	@Failure		404	{object}	httputil.HTTPError
//	@Router			/albums/{id} [delete]
//
// @Param X-ApplicationId header string true "Application  ID"
func (c *Controller) DeleteAlbum(ctx *gin.Context) {
	id := ctx.Param("id")
	service := services.NewAlbumService()

	result, err := service.DeleteAlbum(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Albums could not be deleted. Please verify that album exists."})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, result)
}
