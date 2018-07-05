package packages

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
)

type Package struct {
	Arch string
	Repository string
	Name string
	Version string
	Description string
	LastUpdated string
	FlagDate string
}

func GetPackages(arch,repo,query,maintainer,flagged string) ([]Package,string,string) {
	packages := make([]Package,0)

	url := fmt.Sprintf("https://www.archlinux.org/packages/?q=%s&maintainer=%s&flagged=%s",query,maintainer,flagged)
	if arch != "" {
		url += "&arch="
		url += arch
	}

	if repo != "" {
		url += "&repo="
		url += repo
	}

	doc,err:=goquery.NewDocument(url)
	if err != nil {
		return packages,"0","0"
	}

	doc.Find(".results > tbody > tr").Each(func(i int, selection *goquery.Selection) {
		pkg := Package{}
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			switch i {
			case 0:
				pkg.Arch = selection.Text()
			case 1:
				pkg.Repository = selection.Text()
			case 2:
				pkg.Name = selection.Text()
			case 3:
				pkg.Version = selection.Text()
			case 4:
				pkg.Description = selection.Text()
			case 5:
				pkg.LastUpdated = selection.Text()
			case 6:
				pkg.FlagDate = selection.Text()
			}
		})
		packages = append(packages, pkg)
	})

	num := "0"
	pages := "1"
	doc.Find("#pkglist-results > div > p").Each(func(i int, selection *goquery.Selection) {
		tmp := strings.Split(strings.Replace(selection.Text(),".","",-1)," ")
		num = tmp[0]
		if len(tmp) > 10 {
			pages = tmp[len(tmp) -1]
		}
	})

	return packages,num,pages
}