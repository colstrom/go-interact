package interact_test

import (
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/vito/go-interact/interact"
)

var _ = Describe("Resolving into strings", func() {
	Context("when the destination is empty", func() {
		BeforeEach(func() {
			destination = strDst("")
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when a string is entered", Example{
				Prompt: "some prompt",

				Input: "forty two\r",

				ExpectedAnswer: "forty two",
				ExpectedOutput: "some prompt (): forty two\r\n",
			}),

			Entry("when a blank line is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: "",
				ExpectedOutput: "some prompt (): \r\n",
			}),
		)

		Context("when required", func() {
			BeforeEach(func() {
				destination = interact.Required(destination)
			})

			DescribeTable("Resolve", (Example).Run,
				Entry("when a string is entered", Example{
					Prompt: "some prompt",

					Input: "forty two\r",

					ExpectedAnswer: "forty two",
					ExpectedOutput: "some prompt: forty two\r\n",
				}),

				Entry("when a blank line is entered, followed by EOF", Example{
					Prompt: "some prompt",

					Input: "\r",

					ExpectedAnswer: "",
					ExpectedErr:    io.EOF,
					ExpectedOutput: "some prompt: \r\nsome prompt: ",
				}),

				Entry("when a blank line is entered, followed by a string", Example{
					Prompt: "some prompt",

					Input: "\rforty two\r",

					ExpectedAnswer: "forty two",
					ExpectedOutput: "some prompt: \r\nsome prompt: forty two\r\n",
				}),
			)
		})
	})

	Context("when the destination is not empty", func() {
		BeforeEach(func() {
			destination = strDst("some default")
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when a string is entered", Example{
				Prompt: "some prompt",

				Input: "forty two\r",

				ExpectedAnswer: "forty two",
				ExpectedOutput: "some prompt (some default): forty two\r\n",
			}),

			Entry("when a blank line is entered", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: "some default",
				ExpectedOutput: "some prompt (some default): \r\n",
			}),
		)
	})
})

func strDst(dst string) *string {
	return &dst
}
