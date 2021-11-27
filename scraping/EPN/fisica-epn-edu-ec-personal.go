package scrapfisica

import (
	"fmt"

	"github.com/Club-de-Software-EPN/goscraper.edu.ec/resources"
	"github.com/Club-de-Software-EPN/goscraper.edu.ec/sites"
	"github.com/gocolly/colly"
)

func GetPersonalInfo(){

	var EPN = sites.GetAllWebsites()["Escuela Politécnica Nacional"]
	domain := EPN.Subdomains["Fisica"] + "." +
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
		personalName := htmlElement.ChildText("h3", "a")
		email := htmlElement.ChildText("span[id='cloakfef1a37aaaf528bf710437d4de938b6d']")
		photo := htmlElement.ChildAttrs("img", "src")
		data = append(data, [][]string{{personalName,email, photo[0]}}...)
	})

	collector.Visit(sites.GetVulnerableCollageURLs()["EPN"][2])
	resources.SaveDataOnCSVFile(data, "data/EPN/fisica-epn-edu-ec-personal-academico")
}
