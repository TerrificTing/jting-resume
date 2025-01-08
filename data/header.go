package data

type Link struct {
	Title string
	URL   string
}

type Header struct {
	Name        string
	Email       string
	Location    string
	PhoneNumber string
	LinkList    []Link
}

var HeaderData = Header{
	Name:        "Jeffrey Ting",
	Email:       "jting@bu.edu",
	Location:    "New York, NY",
	PhoneNumber: "(917)757-5886",
	LinkList: []Link{
		{
			"LinkedIn",
			"https://www.linkedin.com/in/jeffrey-ting-08b606253/",
		},
		{
			"GitHub",
			"https://github.com/KingTingTheGreat",
		},
	},
}
