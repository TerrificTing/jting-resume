package data

type Project struct {
	Title        string
	Role         string
	URL          string
	KeyPointList []string
}

var ProjectData = []Project{
	{
		Title: "2048 Game",
		Role:  "Full-Stack Software Engineer",
		URL:   "https://github.com/CSX-20235/lab-7-tiles-2048-gui-TerrificTing",
		KeyPointList: []string{
			"Collaborated with a partner to recreate the game 2048 from scratch in Java as a comprehensive frontend and backend project",
			"Utilized the Java keyboard API to allow users to intuitively access game controls from the keyboard, significantly improving the user experience",
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
}
