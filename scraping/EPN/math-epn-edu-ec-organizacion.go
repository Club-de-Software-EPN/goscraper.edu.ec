package scrapmath

import (
	"fmt"

	"github.com/Club-de-Software-EPN/goscraper.edu.ec/resources"
	"github.com/Club-de-Software-EPN/goscraper.edu.ec/sites"
	"github.com/gocolly/colly"
)

func GetPersonalInfo(){

	var EPN = sites.GetAllWebsites()["Escuela Politécnica Nacional"]
	domain := EPN.Subdomains["Departamento de Matemática"] + "." +
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
		email := htmlElement.ChildText("span[id='cloak322ac1b843daf7e99f383a737ed04a75']")
		photo := htmlElement.ChildAttrs("img", "src")
		data = append(data, [][]string{{personalName,email, photo[0]}}...)
	})

	collector.Visit(sites.GetVulnerableCollageURLs()["EPN"][1])
	resources.SaveDataOnCSVFile(data, "data/EPN/math-epn-edu-ec-organizacion")
}
