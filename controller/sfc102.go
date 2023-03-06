package controller

import (
	"erp-web/modles"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Sfc102Controller struct {
	*gorm.DB
}

// func (ctl *Sfc102Controller) GetOrdnoList(gctx *gin.Context) {
func (c *Sfc102Controller) GetOrdnoList(ctx *gin.Context) {

	var orders []modles.Order
	var count int64
	c.Limit(10).Find(&orders).Limit(-1).Count(&count)
	zap.S().Info(fmt.Sprintf("查询数据：%d 条，共：%d 条", len(orders), count))
	ctx.JSON(200, orders)

	/*	//stream
		conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		c := proto.NewOrderClient(conn)
		//go语言推荐的是返回一个error和一个正常的信息
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
		resp, err := c.GetOrderList(ctx, &proto.OrderRequest{CurrentPage: 1, PageSize: 20})
		if err != nil {
			st, ok := status.FromError(err)
			if !ok {
				// Error was not a status error
				panic("解析error失败")
			}
			fmt.Println(st.Message())
			fmt.Println(st.Code())
		}

		//c.Limit(10).Find(&orders).Limit(-1).Count(&count)
		zap.S().Info(fmt.Sprintf("查询数据：%d 条，共：%d 条", len(resp.Data), resp.Total))
		gctx.JSON(200, resp)*/

}

func NewSfc102Controller(db *gorm.DB) *Sfc102Controller {
	return &Sfc102Controller{db}
}
