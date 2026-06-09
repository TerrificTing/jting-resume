package data

type Project struct {
	Title        string
	Role         string
	URL          string
	KeyPointList []string
}

var ProjectData = []Project{
	{
		Title:	"U-Fund Project",
		Role:	"Full-Stack Software Engineer",
		URL:	"https://github.com/RIT-SWEN-261-05/team-project-team-1a",
		KeyPointList: []string{
			"Developed a full-stack web application for a nonprofit organization as part of a 4-person Agile software engineering team",
			"Designed and implemented role-based authentication and authorization features supporting multiple user account types",
			"Built frontend components in Angular/TypeScript and backend services in Java Spring, integrating REST APIs and relational database operations",
		},
	},
	
	{
		Title:	"Music Database Project",
		Role:	"Full-Stack Software Engineer",
		URL: 	"https://github.com/hl9082/csci320-group20",
		KeyPointList: []string{
			"Developed a full-stack music recommendation platform using Flask, Python, SQL, HTML/CSS, and JavaScript",
			"Implemented similarity-based recommendation algorithms to generate personalized song suggestions from user preference data",
			"Designed backend APIs for user management, recommendation generation, and database interaction",
			"Collaborated with a team of four using Git/GitHub for version control and feature integration",
		},
	},

	/*
	{
		Title: "2048 Game",
		Role:  "Full-Stack Software Engineer",
		URL:   "https://github.com/TerrificTing/2048-game",
		KeyPointList: []string{
			"Collaborated with a partner to recreate the game 2048 from scratch in Java as a comprehensive frontend and backend project",
			"Utilized the Java keyboard API to allow users to intuitively access game controls from the keyboard, significantly improving the user experience",
			"Built using Java, PTUI, and GUI",
		},
	},

	{
		Title: "Do Not Find The Fox",
		Role:  "Full-Stack Software Engineer",
		URL:   "https://github.com/TerrificTing/DoNotFindTheFox",
		KeyPointList: []string{
			"Recreated a simple game where users choose tiles and avoid spelling the word \"fox\"",
			"Enabled users to choose to use the mouse or keyboard as an input method, improving the user experience and the program's accessibility",
			"Created multiple classes to ensure code efficiency, modularity, and maintainability.",
		},
	},
	*/
}
