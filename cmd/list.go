// Copyright © 2016 davey mcglade  <davey.mcglade@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/daveym/lint/lintapi"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Short description",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {

		var ItemRequest Lint.ItemRequest
		ItemRequest.ConsumerKey = ConsumerKey
		ItemRequest.AccessToken = AccessToken
		ItemRequest.State = Lint.ItemStateAll
		ItemRequest.ContentType = string(Lint.ContentTypeArticle)
		ItemRequest.DetailType = string(Lint.DetailTypeSimple)
		ItemRequest.Count = 10

		items, err := Lint.GetItems(ItemRequest)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(items.GivenTitle)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
