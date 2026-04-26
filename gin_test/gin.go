package gintest

import "github.com/gin-gonic/gin"

type Status struct {
	Code int            `json:"code"`
	Data string         `json:"data"`
	Body map[string]any `json:"body"`
}

func IndexF(c *gin.Context) {
	c.JSON(200, Status{
		Code: 0,
		Data: "success",
		Body: make(map[string]any),
	})
}

func MyGin() {
	// 1. 初始化MyGin
	r := gin.Default()
	// 2. 挂载路由
	r.GET("/index")
	// 3. 绑定端口，运行
	r.Run(":8081")
}
