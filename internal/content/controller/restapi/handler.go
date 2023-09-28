package restapi

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/something-to-start-with/api-server-go/internal/content"
)

type Handler struct {
	s content.Service
}

func New(s content.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) GetContents(ctx *gin.Context) {
	contents, err := h.s.GetAll()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, newContentsResponses(contents))
}

func (h *Handler) Create(ctx *gin.Context) {
	cr := new(contentRequest)
	if err := ctx.BindJSON(cr); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	created, err := h.s.Create(cr.toModel())
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, newContentResponse(created))
}

func (h *Handler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	cr := new(contentRequest)
	if err := ctx.BindJSON(cr); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	updated, err := h.s.Update(id, cr.toModel())
	if err != nil {
		if errors.Is(err, content.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	ctx.JSON(http.StatusOK, newContentResponse(updated))
}

func (h *Handler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	found, err := h.s.GetByID(id)
	if err != nil {
		if errors.Is(err, content.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	ctx.JSON(http.StatusOK, newContentResponse(found))
}

func (h *Handler) DeleteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.s.DeleteByID(id)
	if err != nil {
		if errors.Is(err, content.ErrNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	ctx.Status(http.StatusOK)
}
