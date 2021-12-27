package bdd

import (
	"testing"
	//在covey命名空间，不用convey.xxx
	. "github.com/smartystreets/goconvey/convey"
)

//Behavior Driven Development
//Tests are also good documents
//BDD in go : github.com/smartystreets/goconvey

func TestSpec(t *testing.T) {
	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})

}
