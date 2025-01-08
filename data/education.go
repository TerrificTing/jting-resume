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
		School:      "Boston University",
		GPA:         "3.86",
		DegreeTitle: "Bachelor's Degree",
		MajorList: []string{
			"Computer Science",
			"Philosophy",
		},
		MinorList: []string{
			"Women's Studies",
		},
		RelevantCoursework: []string{
			"Distributed Systems", "Web Application Development", "Computer Networks", "Databases", "Software Engineering", "Information Security", "Algorithm Design and Analysis", "Concepts of Programming Languages", "Computer Systems Architecture", "Discrete Mathematics", "Computational Linear Algebra", "Probabilistic Computing",
		},
		StartDate: "2021",
		EndDate:   "2025",
	},
}
