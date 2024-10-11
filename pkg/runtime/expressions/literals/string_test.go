package literals

import (
	"context"
	"testing"

	"github.com/MontFerret/ferret/pkg/runtime/values"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStringLiteral_Exec(t *testing.T) {
	Convey("Given a StringLiteral", t, func() {
		tests := []struct {
			input    StringLiteral
			expected string
		}{
			{input: "hello", expected: "hello"},
			{input: "world", expected: "world"},
		}

		for _, test := range tests {
			Convey("When Exec is called with input "+string(test.input), func() {
				result, err := test.input.Exec(context.Background(), nil)
				Convey("Then it should not return an error", func() {
					So(err, ShouldBeNil)
				})

				Convey("And the result should be the expected value", func() {
					value, ok := result.(values.String)
					So(ok, ShouldBeTrue)

					So(value.String(), ShouldEqual, test.expected)
				})
			})
		}
	})
}
