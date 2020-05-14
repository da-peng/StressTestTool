package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/charts"
)

// StatisticsOfRT 响应时间统计图 生成RT统计文件
func StatisticsOfRT(transactionRTs []int64) {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.TitleOpts{Title: "响应时间&时间-图表"},
		// charts.XAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
		// charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
	)
	scatter.SetSeriesOptions(charts.LabelTextOpts{Show: true, Position: "right"})

	items := []string{}
	RT := []int64{}
	if len(transactionRTs) > 100 { //事务响应统计数据 大于100 则显示数据粒度减小
		for index, v := range transactionRTs {
			if index%3 == 0 { // 图像数据展示 颗粒度  取3的倍数
				items = append(items, strconv.Itoa(index))
				RT = append(RT, v)
			}
		}
	} else {
		for index, v := range transactionRTs {
			items = append(items, strconv.Itoa(index))
			RT = append(RT, v)
		}

	}

	scatter.AddXAxis(items).
		AddYAxis("响应时间(单位ms）", RT)

	f, err := os.Create("StressResponseTime.html")

	if err != nil {
		log.Println(err)
	}
	scatter.Render(f)
}
