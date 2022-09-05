package helmcontroller

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) UploadChart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		chartMuseumURL := "http://175.24.198.168:8080"

		data, err := ioutil.ReadFile("mychart-0.1.0.tgz")
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		client := http.Client{}
		request, err := http.NewRequest(http.MethodPost, chartMuseumURL+"/api/charts", bytes.NewReader(data))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		response, err := client.Do(request)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		b, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		log.Println(string(b))

		ctx.JSON(http.StatusOK, gin.H{})
	}
}
