package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/multimedia_ms/domain/model"
	"github.com/multimedia_ms/usecase/repository"
	"net/http"
	"strconv"
)

type entryHandler struct {
	entryRepository repository.EntryRepository
}

type EntryHandler interface {
	CreateEntry(c *gin.Context)
	GetEntries(c *gin.Context)
	GetEntry(c *gin.Context)
	UpdateEntry(c *gin.Context)
	DeleteEntry(c *gin.Context)
}

func NewEntryHandler(eR repository.EntryRepository) EntryHandler {
	return &entryHandler{entryRepository: eR}
}

func (eH *entryHandler) GetEntry(c *gin.Context) {
	idString := c.Param("userId")
	userId, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	e, err := eH.entryRepository.FindByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func (eH *entryHandler) GetEntries(c *gin.Context) {

	e, err := eH.entryRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func (eH *entryHandler) CreateEntry(c *gin.Context) {
	e := &model.Files{}

	if err := c.Bind(e); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := eH.entryRepository.Store(e); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func (eH *entryHandler) UpdateEntry(c *gin.Context) {
	idString := c.Param("userId")
	userId, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	e := &model.Files{UserId: userId}

	if err := c.Bind(e); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := eH.entryRepository.Update(e); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func (eH *entryHandler) DeleteEntry(c *gin.Context) {
	idString := c.Param("userId")
	userId, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := eH.entryRepository.Delete(&model.Files{UserId: userId}); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "ok"})
}
