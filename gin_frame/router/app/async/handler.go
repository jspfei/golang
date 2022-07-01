package async

import (
	"github.com/gin-gonic/gin"
)

func asyncHandler(c *gin.Context) {
	// 需要搞一个副本
	copyContext := c.Copy()
	// 异步处理
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行：" + copyContext.Request.URL.Path)
	}()

}
func syncHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行：" + c.Request.URL.Path)
}
