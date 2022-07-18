package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"fmt"
)

func main(){
	os.Setenv("VERSION", "go1.18")
	r := gin.Default()


	r.Use(SetVersionInResponse,ReturnRequestHeaderWithinResponse)

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[ACCESS] %v |%v %s| %s | %s %s| %s\n",
        param.TimeStamp.Format("2006/01/02 - 15:04:05"),
        param.StatusCode, 
        param.Latency,
        param.ClientIP,
        param.Method,
        param.Path,
        param.ErrorMessage,
    )
	}))

	
	r.GET("/healthz", healthz)  
	r.Run(":9090") 
}

func healthz(c *gin.Context){

	c.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK, 
		"msg":"get health",  
		"data":"welcome", 
	})
}

//get version and return in response
func SetVersionInResponse(c *gin.Context) {
	c.Header("VERSION", os.Getenv("VERSION"))
	c.Next()
}

//get request headers and return within response
func ReturnRequestHeaderWithinResponse(c *gin.Context) {
	for k, v := range c.Request.Header {
		c.Header(k, strings.Join(v, " "))
	}
	c.Next()
}



