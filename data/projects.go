package data

type Project struct {
	Title        string
	Role         string
	URL          string
	KeyPointList []string
}

var ProjectData = []Project{
	{
		Title: "AccessiScan",
		Role:  "Lead Software Engineer",
		URL:   "https://accessiscan.vercel.app",
		KeyPointList: []string{
			"Led a technical team in the creation of a tool that enables developers to assess the accessibility of their web projects",
			"Learned how to create, use, and publish a Chrome extension, enabling users to scan web projects without a live deployment",
			"Designed and outlined the high-level system architecture and played a key role in its implementation",
			"Built using NextJS + TypeScript, Go, MongoDB, Vercel, Python + Flask, and GitHub OAuth",
		},
	},
	{
		Title: "Toni Tunes",
		Role:  "Full-Stack Software Engineer",
		URL:   "https://toni-tunes.vercel.app",
		KeyPointList: []string{
			"Currently building a service that measures the uniqueness of users’ Spotify listening history",
			"Uses Spotify OAuth to securely sign users in and manage their data",
			"Simplified scalability by designing stateless and serverless system architecture",
			"Built using Spotify’s Web API and OAuth, NextJS + TypeScript and MongoDB +  AWS",
		},
	},
	{
		Title: "LEGO Mosaic Generator",
		Role:  "Full-Stack Software Engineer",
		URL:   "https://image-to-lego.vercel.app",
		KeyPointList: []string{
			"Developed a full-stack web application that converts an input image to a Lego mosaic and a spreadsheet containing a breakdown of the specific LEGO pieces required to build the mosaic in real life",
			"Built using Flask, PIL (Pillow), NumPy, Pygame, and OpenPyXl",
		},
	},
	{
		Title: "Trivia App",
		Role:  "Full-Stack Software Engineer",
		URL:   "https://github.com/KingTingTheGreat/trivia-app",
		KeyPointList: []string{
			"Designed and built a trivia game to be hosted on a local network, meeting the needs of my family’s annual holiday games",
			"Utilizes goroutines,websockets, and a centralized player store to easily support hundreds of concurrent users",
			"Clearly and concisely documented the code and automated its start up, allowing it to be easily shared and used by others, including those with minimal technical experience or knowledge",
			"Built using NextJS + TypeScript and Golang",
		},
	},
}
