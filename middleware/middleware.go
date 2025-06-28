package middleware
import(
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
)

func LoggingMiddleware()gin.HandlerFunc{
return func(ctx *gin.Context) {
	startTime:=time.Now()
	ctx.Next()
	duration:=time.Since(startTime)
	fmt.Println(duration.Seconds())
}
}