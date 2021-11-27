package sites

type WebSite struct {
	Scheme         string
	Subdomains     map[string]string
	RootDomain     string
	STLD           string // Sponsored top-level domains
	CCTLD          string // Country-code top-level domains
	VulnerableURLs map[string]string
}

func GetAllWebsites() (sites map[string]WebSite) {
	sites = map[string]WebSite{
		"Escuela Politécnica Nacional": {
			Scheme:     "https",
			Subdomains: map[string]string{
				"ModeMat": "modemat",
				"Departamento de Matemática": "math",
				"Fisica": "fisica",
			},
			RootDomain: "epn",
			STLD:       "edu",
			CCTLD:      "ec",
			VulnerableURLs: map[string]string{
				"ModeMat": "/personal",
				"Organización": "/index.php/organizacion",
				"Personal académico": "/index.php/departamento/personal-academico",
				"Directorio Telefónico": "/wp-content/uploads/2019/09/DirectorioTelefonicoV5.pdf",
			},
		},
		"Universidad de las Fuerzas Armadas": {
			Scheme: "https",
			Subdomains: map[string]string{
				"Ciencias Exactas": "dece",
				"Ciencias Económicas Administrativas y de Comercio": "deceac-el",
			},
			RootDomain: "espe",
			STLD:       "edu",
			CCTLD:      "ec",
			VulnerableURLs: map[string]string{
				"Personal-Docentes-Investigadores": "/personal-docente-e-investigadores/",
			},
		},
		"Universidad San Francisco de Quito": {
			Scheme:     "https",
			Subdomains: map[string]string{},
			RootDomain: "usfq",
			STLD:       "edu",
			CCTLD:      "ec",
		},
		"Universidad Central del Ecuador": {
			Scheme:     "https",
			Subdomains: map[string]string{},
			RootDomain: "uce",
			STLD:       "edu",
			CCTLD:      "ec",
		},
		"Escuela Politécnica del Litoral": {
			Scheme:     "https",
			Subdomains: map[string]string{},
			RootDomain: "espol",
			STLD:       "edu",
			CCTLD:      "ec",
		},
		"Universidad Técnica de Cotopaxi": {
			Scheme:     "https",
			Subdomains: map[string]string{},
			RootDomain: "utc",
			STLD:       "edu",
			CCTLD:      "ec",
		},
	}
	return
}

func GetVulnerableCollageURLs() map[string][]string {

	var ESPE = GetAllWebsites()["Universidad de las Fuerzas Armadas"]

	return map[string][]string{
		"ESPE": {
			ESPE.Scheme + "://" +
				ESPE.Subdomains["Ciencias Económicas Administrativas y de Comercio"] + "." +
				ESPE.RootDomain + "." + ESPE.STLD + "." + ESPE.CCTLD + // espe.edu.ec
				ESPE.VulnerableURLs["Personal-Docentes-Investigadores"],
		},
	}
}
