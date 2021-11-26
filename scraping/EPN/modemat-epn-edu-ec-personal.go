package scrapmodemat

import (
	"fmt"

	"github.com/Club-de-Software-EPN/goscraper.edu.ec/resources"
	"github.com/Club-de-Software-EPN/goscraper.edu.ec/sites"
	"github.com/gocolly/colly"
)

func GetPersonalInfo(){

	var EPN = sites.GetAllWebsites()["Escuela Politécnica Nacional"]
	domain := EPN.Subdomains["ModeMat"] + "." +
	EPN.RootDomain + "." + EPN.STLD + "." + EPN.CCTLD

	fmt.Println(domain)
	// manages the network communication
	collector := colly.NewCollector(
		// the domain used by the Collector
		colly.AllowedDomains(domain),
	)

	var data [][]string
	data = append(data, [][]string{{"Personal", "Email", "UrlPhotoProfile"}}...)

	collector.OnHTML(".elementor-widget-wrap", func(htmlElement *colly.HTMLElement){
		personalName := htmlElement.ChildText("h4")
		email := htmlElement.ChildText("a", "href")
		photo := htmlElement.ChildAttrs("img", "src")
		data = append(data, [][]string{{personalName,email, photo[0]}}...)
	})

	collector.Visit(sites.GetVulnerableCollageURLs()["EPN"][0])
	resources.SaveDataOnCSVFile(data, "data/EPN/modemat-epn-edu-ec-personal")
}
