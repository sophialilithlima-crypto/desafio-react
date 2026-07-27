package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/sophialilithlima-crypto/desafio-react-backend/config"
)

func setupTestRouter() *gin.Engine {

	config.ConnectDatabase()

	router := gin.Default()

	router.POST("/categorias", CreateCategoria)

	return router
}


func TestCreateCategoriaEndpoint(t *testing.T) {

	router := setupTestRouter()

	body := map[string]string{
		"nome": "Categoria Teste Integração",
	}

	jsonBody, _ := json.Marshal(body)

	request := httptest.NewRequest(
		http.MethodPost,
		"/categorias",
		bytes.NewBuffer(jsonBody),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)


	assert.Equal(
		t,
		http.StatusCreated,
		response.Code,
	)


	var resultado map[string]interface{}

	err := json.Unmarshal(
		response.Body.Bytes(),
		&resultado,
	)

	assert.Nil(t, err)

	assert.NotNil(
		t,
		resultado["data"],
	)
}