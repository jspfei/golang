package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

//json
func jsonHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "someJSON", "status": 200})
}

// 结构体响应

func someStructHandler(c *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Number  int
	}
	msg.Name = "root"
	msg.Message = "message"
	msg.Number = 123
	c.JSON(200, msg)
}

//xml
func someXmlHandler(c *gin.Context) {
	c.XML(200, gin.H{"message": "abc"})
}

//YAML响应
func someYamlHandler(c *gin.Context) {
	c.YAML(200, gin.H{"message": "zhangsan"})
}

//protobuf 格式
func someProtoBufHandler(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	//定义数据
	label := "label"
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(200, data)
}
