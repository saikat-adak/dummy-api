package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func main() {
	router := gin.Default()
	router.GET("/download/xlsx", getDummyXlsxFile)
	router.GET("/download/xlsx2", getDummyXlsxUsingExcelize)
	router.GET("/download/xlsx3", getDummyXlsxFromByteString)
	router.GET("/json", getDummyJson)

	router.Run("localhost:8080")
}

func getDummyXlsxUsingExcelize(c *gin.Context) {
	xlsx := excelize.NewFile()
	// Create a new sheet.
	index := xlsx.NewSheet("Sheet2")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)

	// var b *bytes.Buffer
	b, err := xlsx.WriteToBuffer()
	if err != nil {
		// return nil, err
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}

func getDummyXlsxFromByteString(c *gin.Context) {

	strBytes := []byte("paste your byte string here.")

	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", strBytes)
}

func getDummyXlsxFile(c *gin.Context) {
	file := xlsx.NewFile()
	for i := 0; i < 10; i++ {
		sheet, err := file.AddSheet(fmt.Sprintf("sheet-%d", i))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = "dummy cell!"
	}
	var b bytes.Buffer
	if err := file.Write(&b); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}

func getDummyJson(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []user{
	{Id: 546, Username: "John"},
	{Id: 894, Username: "Mary"},
	{Id: 326, Username: "Jane"},
}
