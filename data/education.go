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
			"Discrete Mathematics", "Probability and Statistics", "Computer Systems Programming", "Computer Science Theory", "Data Structures and Algorithms", "Calculus B/C", "Physics 1/2",
		},
		StartDate: "2023",
		EndDate:   "2028",
	},
}
