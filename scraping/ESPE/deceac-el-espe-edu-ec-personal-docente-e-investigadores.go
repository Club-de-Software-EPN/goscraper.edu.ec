// Last execution on November 06, 2021

package scraping

import (
	"fmt"

	"github.com/Club-de-Software-EPN/goscraper.edu.ec/resources"
	"github.com/Club-de-Software-EPN/goscraper.edu.ec/sites"
	"github.com/gocolly/colly"
)


func GetResearcherAndTeachersInfo(){

	var ESPE = sites.GetAllWebsites()["Universidad de las Fuerzas Armadas"]
	domain := ESPE.Subdomains["Ciencias Económicas Administrativas y de Comercio"] + "." +
	ESPE.RootDomain + "." + ESPE.STLD + "." + ESPE.CCTLD

	fmt.Println(domain)
	// manages the network communication
	collector := colly.NewCollector(
		// the domain used by the Collector
		colly.AllowedDomains(domain),
	)

	var data [][]string
	data = append(data, [][]string{{"Professor", "Email", "UrlPhotoProfile"}}...)

	collector.OnHTML(".elementor-widget-wrap", func(htmlElement *colly.HTMLElement){
		professorName := htmlElement.ChildText("h5")
		email := htmlElement.ChildText("span[style='color: #008000; font-family: Agency FB,serif;']")
		photo := htmlElement.ChildAttrs("img", "src")
		if professorName != "" {
			if email == "" {
				email = "empty"
			}
			data = append(data, [][]string{{professorName,email, photo[0]}}...)
		}
	})

	collector.Visit(sites.GetVulnerableCollageURLs()["ESPE"][0])
	resources.SaveDataOnCSVFile(data, "data/ESPE/deceac-el-espe-edu-ec-personal-docente-e-investigadores")
}