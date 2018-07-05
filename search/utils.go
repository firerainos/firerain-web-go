package search

import "github.com/PuerkitoBio/goquery"

type Option struct {
	Value string
	Text string
}

func GetSelectText(selector string) []Option {
	result := make([]Option,0)

	doc,err:=goquery.NewDocument("https://www.archlinux.org/packages")
	if err != nil {
		return result
	}

	doc.Find(selector+" > option").Each(func(i int, selection *goquery.Selection) {
		value,exists:=selection.Attr("value")
		if !exists {
			value = ""
		}
		option := Option{value,selection.Text()}
		result = append(result, option)
	})

	return result
}