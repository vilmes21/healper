package model

type Author struct {
	Name string
}

type Source struct {
	Id int `json:"id"`
	Name string `json:"name"`
	AuthorId int `json:"authorId"`
	LanguageId int `json:"languageId"`
}

type Disease struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Namezh string `json:"namezh"`
	Pinyin string `json:"pinyin"`
	SourceId int `json:"sourceId"`
}

type Recipe struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Namezh string `json:"namezh"`
	Pinyin string `json:"pinyin"`
	SourceId int `json:"sourceId"`
}

type Herb struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Namezh string `json:"namezh"`
	Pinyin string `json:"pinyin"`
}

type RecipeHerb struct {
	RecipeId int `json:"recipeId"`
	HerbId int `json:"herbId"`
	Gram int `json:"gram"`
	Note string `json:"note"`
}

type Policy struct {
	TreatmentVerbId int `json:"treatmentVerbId"`
	OrganId int `json:"organId"`
	TcmOrganId int `json:"tcmOrganId"`
}

type DiseasePolicy struct {
	DiseaseId int `json:"diseaseId"`
	PolicyId int `json:"policyId"`
	Priority int `json:"priority"`
}

type OrganSymptom struct {
	Id int `json:"id"`
	OrganId int `json:"organId"`
	SymptomId int `json:"symptomId"`
}

type DiseaseOrgansymptom struct {
	DiseaseId int `json:"diseaseId"`
	OrgansymptomId int `json:"organsymptomId"`
}



//   treatment_verb_id INTEGER REFERENCES category(id) NOT NULL,
//   organ_id INTEGER REFERENCES category(id),
//   tcm_organ_id INTEGER REFERENCES category(id),
  