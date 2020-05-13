package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/charts"
)

func StatisticsOfRT(responseTimeStatistics []int64) {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.TitleOpts{Title: "压测响应时间时序图"},
		// charts.XAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
		// charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
	)
	scatter.SetSeriesOptions(charts.LabelTextOpts{Show: true, Position: "right"})

	nameItems := []string{}
	responseTime := []int64{}
	if len(responseTimeStatistics) > 100 {
		for i, v := range responseTimeStatistics {
			if i%3 == 0 { // 图像数据展示 颗粒度 ，每N个点取1个
				nameItems = append(nameItems, strconv.Itoa(i))
				responseTime = append(responseTime, v)
			}

		}
	} else {
		for i, v := range responseTimeStatistics {
			nameItems = append(nameItems, strconv.Itoa(i))
			responseTime = append(responseTime, v)
		}

	}

	scatter.AddXAxis(nameItems).
		AddYAxis("响应时间(单位ms）", responseTime)

	f, err := os.Create("StressResponseTime.html")

	if err != nil {
		log.Println(err)
	}
	scatter.Render(f)
}
