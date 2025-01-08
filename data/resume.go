package data

type Resume struct {
	Header         Header
	ExperienceList []Experience
	ProjectList    []Project
	EducationList  []Education
	Skills         Skills
}

var ResumeData = Resume{
	Header:         HeaderData,
	ExperienceList: ExperienceData,
	ProjectList:    ProjectData,
	EducationList:  EducationData,
	Skills:         SkillsData,
}
