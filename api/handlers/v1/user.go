package v1

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	"github/test-ci-cd-github/api/handlers/models"
	l "github/test-ci-cd-github/pkg/logger"
)

// TEST PING
// @Summary TEST PING
// @Description Api for test ping
// @Tags test
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Pong
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/ping/{id} [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	_, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	if id != "ping" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errors.New("Not_ping"),
			})
			h.log.Error("failed to get user", l.Error(errors.New("dsfsdf")))
			return
		
	}

	c.JSON(http.StatusOK, &models.Pong{
		Pong: "pong",
	})
}


// TEST PONG
// @Summary TEST PONG
// @Description Api for test pong
// @Tags test
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Pong
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/pong/{id} [get]
func (h *handlerV1) TestPong(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	_, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	if id != "pong" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errors.New("Not_pong"),
			})
			h.log.Error("failed to get user", l.Error(errors.New("dsfsdf")))
			return
		
	}

	c.JSON(http.StatusOK, &models.Pong{
		Pong: "ping",
	})
}

// TEST GITHUB
// @Summary TEST GITHUB
// @Description Api for test github
// @Tags github
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Pong
// @Failure 400 {object} models.StandardErrorModel
// @Failure 401 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/github/{id} [get]
func (h *handlerV1) TestGithub(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	_, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	if id != "github" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errors.New("Not_github"),
			})
			h.log.Error("failed to get user", l.Error(errors.New("dsfsdf")))
			return
		
	}

	c.JSON(http.StatusOK, &models.Pong{
		Pong: "github action is working",
	})
}