package assertion

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJSONShouldHave(t *testing.T) {
	Convey("Testing JSONShouldHave", t, func() {
		d := []byte(`{"name":"hello","skills":[{"title":"swimming"},{"title":"running"}],"addr":{"country":"indonesia"}}`)

		Convey("Positive comparison", func() {
			So(ShouldBeJSONAndHave(d, "name", "hello"), ShouldBeEmpty)
			So(ShouldBeJSONAndHave(d, "skills.0.title", "swimming"), ShouldBeEmpty)
			So(ShouldBeJSONAndHave(d, "addr.country", "indonesia"), ShouldBeEmpty)
		})

		Convey("Wrong value comparison", func() {
			So(ShouldBeJSONAndHave(d, "name", "hella"), ShouldNotBeEmpty)
			So(ShouldBeJSONAndHave(d, "skills.0.title", "flying"), ShouldNotBeEmpty)
			So(ShouldBeJSONAndHave(d, "addr.country", "desert"), ShouldNotBeEmpty)
		})
	})
}

func TestJSONShouldContain(t *testing.T) {
	Convey("Testing JSONShouldContain", t, func() {
		d := []byte(`{"name":"hello","skills":[{"title":"swimming"},{"title":"running"}],"addr":{"country":"indonesia"}}`)

		Convey("Positive comparison", func() {
			So(ShouldBeJSONAndContain(d, `{"name":"hello"}`), ShouldBeEmpty)
			So(ShouldBeJSONAndContain(d, `{"skills":[{"title":"swimming"}]}`), ShouldBeEmpty)
			So(ShouldBeJSONAndContain(d, `{"skills":[{"title":"running"}]}`), ShouldBeEmpty)
			So(ShouldBeJSONAndContain(d, `{"skills":[{"title":"running"},{"title":"swimming"}]}`), ShouldBeEmpty)
			So(ShouldBeJSONAndContain(d, `{"addr":{"country":"indonesia"}}`), ShouldBeEmpty)
		})

		Convey("Inexisting key comparison", func() {
			So(ShouldBeJSONAndContain(d, `{"nama":"hello"}`), ShouldNotBeEmpty)
			So(ShouldBeJSONAndContain(d, `{"skills":[{"titel":"swimming"}]}`), ShouldNotBeEmpty)
			So(ShouldBeJSONAndContain(d, `{"addr":{"contry":"indonesia"}}`), ShouldNotBeEmpty)
		})

		Convey("Wrong value comparison", func() {
			So(ShouldBeJSONAndContain(d, `{"name":"hella"}`), ShouldNotBeEmpty)
			So(ShouldBeJSONAndContain(d, `{"skills":[{"title":"flying"}]}`), ShouldNotBeEmpty)
			So(ShouldBeJSONAndContain(d, `{"addr":{"country":"desert"}}`), ShouldNotBeEmpty)
		})
	})
}

func TestJSONShouldCount(t *testing.T) {
	Convey("Testing JSONShouldCount", t, func() {
		d := []byte(`{"name":"hello","skills":[{"title":"swimming"},{"title":"running"}],"addr":{"country":"indonesia"}}`)

		So(ShouldBeJSONAndCount(d, `name`, 0), ShouldBeEmpty)
		So(ShouldBeJSONAndCount(d, `name`, 1), ShouldNotBeEmpty)
		So(ShouldBeJSONAndCount(d, `addr`, 1), ShouldBeEmpty)
		So(ShouldBeJSONAndCount(d, `addr`, 0), ShouldNotBeEmpty)
		So(ShouldBeJSONAndCount(d, `skills`, 2), ShouldBeEmpty)
		So(ShouldBeJSONAndCount(d, `skills`, 0), ShouldNotBeEmpty)
	})
}
