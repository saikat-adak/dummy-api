package main

import (
    "fmt"
	"bytes"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/tealeg/xlsx"
)

func main(){
	router := gin.Default()
	router.GET("/download/xlsx", getDummyXlsxFile)
	router.GET("/json", getDummyJson)

    router.Run("localhost:8080")
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
    Id        int     `json:"id"`
    Username  string  `json:"username"`
}
var users = []user{
    {Id: 546, Username: "John"},
    {Id: 894, Username: "Mary"},
    {Id: 326, Username: "Jane"},
}