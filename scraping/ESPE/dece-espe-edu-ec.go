package scraping

import (
	"fmt"

	"github.com/Club-de-Software-EPN/goscraper.edu.ec/resources"
	"github.com/Club-de-Software-EPN/goscraper.edu.ec/sites"
	"github.com/gocolly/colly"
)


func GetResearcherAndTeachersInfo(){

	var ESPE = sites.GetAllWebsites()["Universidad de las Fuerzas Armadas"]
	domain := ESPE.Subdomains["Ciencias Exactas"] + "." +
	ESPE.RootDomain + "." + ESPE.STLD + "." + ESPE.CCTLD

	fmt.Println(domain)
	// manages the network communication
	collector := colly.NewCollector(
		// the domain used by the Collector
		colly.AllowedDomains(domain),
	)

	var data [][]string
	data = append(data, [][]string{{"Professor", "Email"}}...)

	collector.OnHTML(".elementor-widget-wrap", func(htmlElement *colly.HTMLElement){
		professorName := htmlElement.ChildText("h4 class='elementor-heading-title elementor-size-medium'")
		email := htmlElement.ChildText("span[style='font-family: helvetica, arial, sans-serif; font-size: 8pt;'']")
		if professorName != "" {
			if email == "" {
				email = "empty"
			}
			data = append(data, [][]string{{professorName,email, photo[0]}}...)
		}
	})

	collector.Visit(sites.GetVulnerableCollageURLs()["ESPE"][0])
	resources.SaveDataOnCSVFile(data, "data/ESPE/dece-espe-edu-ec-personal-docente-e-investigadores")
}
