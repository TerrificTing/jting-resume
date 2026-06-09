package data

type Experience struct {
	Role         string
	Location     string
	Company      string
	StartDate    string
	EndDate      string
	KeyPointList []string
}

var ExperienceData = []Experience{
	{
		Role: 		"Crew Member",
		Location:	"Henrietta, NY",
		Company: 	"CoreLife Eatery",
		StartDate:	"September 2025",
		EndDate:  	"Present",
		KeyPointList: []string{
			"Served 100+ customers per shift while resolving service issues and maintaining efficient operations during peak hours",
			"Managed customer service, food prep, and inventory while maintaining a high level of organization",
			"Collaborated effectively in a fast-paced environment",
		},
	},
	{
		Role: 		"Technical Instructor",
		Location: 	"Garden City, NY",
		Company: 	"Lavner Education",
		StartDate:	"June 2025",
		EndDate: 	"August 2025",
		KeyPointList: []string{
			"Managed a student roster database, enhancing organization and streamlining administrative processes",
			"Taught programming concepts in Python, Java, Lua, C++, and Scratch to classes of up to 30 students through project-based learning",
			"Provided technical support for 150+ camp-issued laptops weekly, resolving both software and hardware issues",
		},
	},
	
	/*
	{
		Role:      "Camp Counselor",
		Location:  "Syosset, NY",
		Company:   "Town of Oyster Bay",
		StartDate: "July 2023",
		EndDate:   "August 2024",
		KeyPointList: []string{
			"Supervised and ensured the safety of a group of 25 seven-year-old campers",
			"Organized, prepared, and led daily camp activities",
			"Collaborated with co-counselors to create a positive, enjoyable, and safe camp experience for all children",
		},
	},
	*/
	
	/*
	{
		Role:      "Student Volunteer",
		Location:  "Rochester, NY",
		Company:   "Rochester Institute of Technology",
		StartDate: "March 2024",
		EndDate:   "May 2024",
		KeyPointList: []string{
			"Aided vendors with sales and adversiting to promote revenue growth",
			"Organized and maintained orderly event lines in order to ensure attendees' safety",
		},
	},

	{
		Role:      "Configuration Center Intern",
		Location:  "Holbrook, NY",
		Company:   "Future Tech Enterprise, Inc.",
		StartDate: "June 2022",
		EndDate:   "August 2022",
		KeyPointList: []string{
			"Increased production throughput by upwards of 20% by organizing physical products according to their orders, helping all employees work more efficiently",
			"Tagged and imaged a variety of machines, including laptops, desktops, and thin clients",
			"Optimized the processing of products by organizing orders, significantly reducing errors and operating costs",
		},
	},
	*/
}
