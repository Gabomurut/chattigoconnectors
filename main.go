package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type provinciasResponse struct {
	Status   string      `json:"status,omitempty"`
	Response []provincia `json:"response,omitempty"`
}
type departamentosResponse struct {
	Status   string         `json:"status,omitempty"`
	Response []departamento `json:"response,omitempty"`
}
type distritosResponse struct {
	Status   string     `json:"status,omitempty"`
	Response []distrito `json:"response,omitempty"`
}
type provincia struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type departamento struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type distrito struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// albums slice to seed record album data.
var provincias = []provincia{
	{ID: "01", Name: "HUANUCO"},
	{ID: "02", Name: "AMBO"},
	{ID: "03", Name: "DOS DE MAYO"},
	{ID: "04", Name: "HUACAYBAMBA"},
	{ID: "05", Name: "HUAMALIES"},
	{ID: "06", Name: "LEONCIO PRADO"},
	{ID: "07", Name: "MARAÑON"},
	{ID: "08", Name: "PACHITEA"},
	{ID: "09", Name: "PUERTO INCA"},
	{ID: "10", Name: "LAURICOCHA"},
	{ID: "11", Name: "YAROWILCA"},
}

var departamentos = []departamento{
	{ID: "01", Name: "SANTIAGO"},
	{ID: "02", Name: "VIÑA"},
	{ID: "03", Name: "HUINCA"},
}
var distritos = []distrito{
	{ID: "01", Name: "AMBO"},
	{ID: "02", Name: "CAYNA"},
	{ID: "03", Name: "COLPAS"},
	{ID: "04", Name: "CONCHAMARCA"},
	{ID: "05", Name: "HUACAR"},
	{ID: "06", Name: "SAN FRANCISCO"},
	{ID: "07", Name: "SAN RAFAEL"},
	{ID: "08", Name: "TOMAY KICHWA"},
}
var respuestaProvincia = provinciasResponse{
	Status:   "success",
	Response: provincias,
}
var respuestaDepartamento = departamentosResponse{
	Status:   "success",
	Response: departamentos,
}
var respuestaDistrito = distritosResponse{
	Status:   "success",
	Response: distritos,
}

func main() {
	router := gin.Default()
	router.GET("/provincias", getProvincias)
	router.GET("/departamentos", getDepartamentos)
	router.GET("/distritos", getDistritos)

	router.Run("localhost:8080")
}

func getDepartamentos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, respuestaDepartamento)
}

func getDistritos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, respuestaDistrito)
}

// getAlbums responds with the list of all albums as JSON.
func getProvincias(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, respuestaProvincia)
}
