package subdomains

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Club-de-Software-EPN/goscraper.edu.ec/resources"
)

func isActive(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println(url, response.StatusCode)
		if(response.StatusCode >= 200 && response.StatusCode <= 299){
			return true
		}
		return false
	}
}

func GetActiveSubdomains() {

	directories, _:= ioutil.ReadDir("./subdomains")

	for _, directory := range directories {

		var subdomains [][]string

		if directory != directories[len(directories)-1] {

			file, err := os.Open(fmt.Sprintf("subdomains/%s/records-sublist3r.txt", directory.Name()))

			if err != nil {
				log.Fatalln(err)
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)

			for scanner.Scan() {
				if(isActive(fmt.Sprintf("http://%s",scanner.Text()))){
					subdomains = append(subdomains, []string{scanner.Text()})
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}

			resources.SaveDataOnCSVFile(subdomains, fmt.Sprintf("./subdomains/%s/active-subdomains", directory.Name()))

		}
	}
}


