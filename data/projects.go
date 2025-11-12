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
			"Collaborated with a small group to make a website for a non-profit organization",
			"Embedded multiple different types of accounts to give users a unique experience",
			"Built using Angular, TypeScript, HTML/CSS, Spring, HTTP, and Java",
		},
	},
	
	{
		Title:	"Music Database Project",
		Role:	"Full-Stack Software Engineer",
		URL: 	"https://github.com/hl9082/csci320-group20",
		KeyPointList: []string{
			"Collaborated with a team of 4 to design and develop a full-stack music recommendation website",
			"Implemented the backend using Python/Flask to handle data processing and the recommendation system",
			"Performed data analysis on user preference data to generate personalized song recommendations using similarity-based filtering techniques",
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
