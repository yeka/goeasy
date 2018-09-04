package idata

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIData(t *testing.T) {
	json := `{"name":"hello","skills":[{"title":"swimming"},{"title":"running"}],"addr":{"country":"indonesia"}}`
	i := FromJSON([]byte(json))

	Convey("Testing IData", t, func() {
		Convey("Traversing through content", func() {
			Convey("Check simple key", func() {
				So(i.KeyExists("name"), ShouldBeTrue)
				So(i.KeyExists("unknown"), ShouldBeFalse)
				So(i.KeyExists("skills"), ShouldBeTrue)
				So(i.KeyExists("addr"), ShouldBeTrue)
			})

			Convey("Check object key", func() {
				So(i.KeyExists("addr.country"), ShouldBeTrue)
				So(i.KeyExists("addr.city"), ShouldBeFalse)
			})
			Convey("Check array key", func() {
				So(i.KeyExists("skills.0"), ShouldBeTrue)
				So(i.KeyExists("skills.1"), ShouldBeTrue)
				So(i.KeyExists("skills.2"), ShouldBeFalse)
			})

			Convey("Check object in array", func() {
				So(i.KeyExists("skills.0.age"), ShouldBeFalse)
				So(i.KeyExists("skills.0.title"), ShouldBeTrue)
				So(i.KeyExists("skills.1.title"), ShouldBeTrue)
			})
		})

		Convey("Array checking", func() {
			So(i.IsArray("name"), ShouldBeFalse)
			So(i.IsArray("unknown"), ShouldBeFalse)
			So(i.IsArray("skills"), ShouldBeTrue)
			So(i.IsArray("addr"), ShouldBeFalse)
		})

		Convey("Object checking", func() {
			So(i.IsObject("name"), ShouldBeFalse)
			So(i.IsObject("unknown"), ShouldBeFalse)
			So(i.IsObject("skills"), ShouldBeFalse)
			So(i.IsObject("addr"), ShouldBeTrue)
		})

		Convey("Counting", func() {
			So(i.Count("name"), ShouldEqual, 0)
			So(i.Count("unknown"), ShouldEqual, 0)
			So(i.Count("skills"), ShouldEqual, 2)
			So(i.Count("addr"), ShouldEqual, 1)
			So(i.Count("skills.0"), ShouldEqual, 1)
			So(i.Count("skills.1"), ShouldEqual, 1)
			So(i.Count("skills.2"), ShouldEqual, 0)
			So(i.Count("addr.country"), ShouldEqual, 0)
			So(i.Count("addr.city"), ShouldEqual, 0)
		})

		Convey("Get string", func() {
			So(i.GetString("name"), ShouldEqual, "hello")
			So(i.GetString("unknown"), ShouldEqual, "")
			So(i.GetString("skills"), ShouldEqual, "")
			So(i.GetString("skills.0.title"), ShouldEqual, "swimming")
			So(i.GetString("skills.1.title"), ShouldEqual, "running")
			So(i.GetString("skills.2.title"), ShouldEqual, "")
			So(i.GetString("addr"), ShouldEqual, "")
			So(i.GetString("addr.country"), ShouldEqual, "indonesia")
			So(i.GetString("addr.city"), ShouldEqual, "")
		})
	})
}
