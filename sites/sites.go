package sites

type WebSite struct{
	scheme string
	subdomains []string
	rootDomain string
	sTLD string // Sponsored top-level domains
	ccTLD string // Country-code top-level domains 
}

func GetAllSites() (sites map[string]WebSite){
	sites = map[string]WebSite{
		"Escuela Politécnica Nacional" : {
			scheme: "https",
			subdomains: []string{
				"",
				"fis",
			},
			rootDomain: "epn",
			sTLD: "edu",
			ccTLD: "ec",
		},
		"Universidad de las Fuerzas Armadas" : {
			scheme: "https",
			subdomains: []string{
				"",
			},
			rootDomain: "espe",
			sTLD: "edu",
			ccTLD: "ec",
		},
		"Universidad San Francisco de Quito" : {
			scheme: "https",
			subdomains: []string{
				"",
			},
			rootDomain: "usfq",
			sTLD: "edu",
			ccTLD: "ec",
		},
		"Universidad Central del Ecuador" : {
			scheme: "https",
			subdomains: []string{
				"",
			},
			rootDomain: "uce",
			sTLD: "edu",
			ccTLD: "ec",
		},
		"Escuela Politécnica del Litoral" : {
			scheme: "https",
			subdomains: []string{
				"",
			},
			rootDomain: "espol",
			sTLD: "edu",
			ccTLD: "ec",
		},
	}
	return
}