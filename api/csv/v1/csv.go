package v1

import (
	"bufio"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func Create(cx *gin.Context) {
	file, err := cx.FormFile("file")
	if err != nil {
		pp.Println(err)
		return
	}
	f, _ := file.Open()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pp.Println(scanner.Text())
	}

	cx.String(http.StatusOK, "Uploaded...")
	return
}
