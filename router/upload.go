package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func files(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		if err := c.SaveUploadedFile(file, "/works/cache/images/"+file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	c.String(http.StatusOK, "文件上传成功")
}

func uploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "image.html", gin.H{
		"title": "图片上传",
	})
}

func uploadView(c *gin.Context) {
	files, err := os.Open("/works/cache/images")
	if err != nil {
		c.String(http.StatusBadGateway, "获取路径错误")
	}
	defer files.Close()
	lists, err := files.Readdir(-1)
	if err != nil {
		c.String(http.StatusBadGateway, "获取图片资源失败")
	}
	pics := []map[string]string{}
	for _, file := range lists {
		s1 := "http://hrm.allkic.xyz:8080/img/"
		str := s1 + file.Name()
		pics = append(pics[:], map[string]string{
			"label": file.Name(),
			"url":   str,
		})
	}
	c.HTML(http.StatusOK, "view.html", gin.H{
		"title": "商品预览图",
		"lists": pics,
	})
}

func upload(route *gin.Engine) {
	route.LoadHTMLGlob("html/upload/*")
	route.Static("/static", "html/static")
	route.Static("/img", "/works/cache/images")
	route.POST("/upload/files", files)
	route.GET("/upload/page", uploadPage)
	route.GET("/upload/view", uploadView)
}
