package comm

import (
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStockCalendar(t *testing.T) {
	configure := new(Conf)
	configure.Parse("/Users/datochan/WorkSpace/GoglandProjects/src/cquant/configure.toml")

	Convey("检测默认股票日历实例的生成", t, func() {
		// 默认加载股票日历数据
		calendarPath := fmt.Sprintf("%s%s", configure.App.DataPath, configure.Tdx.Files.Calendar)
		calendar, err := DefaultStockCalendar(calendarPath)
		So(err, ShouldEqual, nil)

		Convey("测试下一个交易日的生成", func() {
			next, err := calendar.NextDay("20170931")
			So(err, ShouldEqual, nil)
			So(next, ShouldEqual, "20171009")
		})

		Convey("测试前一个交易日的生成", func() {
			prev, err := calendar.PrevDay("20171007")
			So(err, ShouldEqual, nil)
			So(prev, ShouldEqual, "20170929")
		})

		Convey("测试股市日历的遍历情况", func() {
			calendar.Each(func (dateItem CalendarModel) error{
				if dateItem.Date == 19981231 {
					So(dateItem.YearEnd, ShouldBeTrue)
					return fmt.Errorf("正常停止")
				}
				return nil
			})
		})


	})
}
