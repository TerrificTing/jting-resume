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
	Name:        "Jason Ting",
	Email:       "jt8227@rit.edu",
	Location:    "Rochester, NY",
	PhoneNumber: "(516)-580-2303",
	LinkList: []Link{
		{
			"LinkedIn",
			"https://www.linkedin.com/in/jason-ting-829b66316/",
		},
		{
			"GitHub",
			"https://github.com/TerrificTing",
		},
	},
}
