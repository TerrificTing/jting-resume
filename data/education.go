package data

type Education struct {
	School             string
	GPA                string // optional
	DegreeTitle        string
	MajorList          []string
	MinorList          []string // optional
	RelevantCoursework []string
	StartDate          string
	EndDate            string
}

var EducationData = []Education{
	{
		School:      "Rochester Institute of Technology",
		GPA:         "",
		DegreeTitle: "Bachelor of Science",
		MajorList: []string{
			"Computer Science",
		},
		MinorList: []string{},
		RelevantCoursework: []string{
			"Software Engineering", 
			"Concepts of Computer Systems",
			"Principles of Data Management",
			"Discrete Mathematics", 
			"Probability and Statistics",
			"Linear Algebra",
			"Computer Science Theory",
			"Data Structures and Algorithms",
			"Algorithm Design and Analysis", 
			"Calculus B/C", 
			"Physics 1/2",
			"Computer Systems Programming", 
		},
		StartDate: "2023",
		EndDate:   "2027",
	},
}
