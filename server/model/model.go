package model

import (
	"../config"
)

type Author struct {
	Name string
}

type Source struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	AuthorId   int    `json:"authorId"`
	LanguageId int    `json:"languageId"`
}

func (x Source) Table() string {
	return `source`
}

type Disease struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Namezh   string `json:"namezh"`
	Pinyin   string `json:"pinyin"`
	SourceId int    `json:"sourceId"`
}

func (x Disease) Table() string {
	return config.TbDisease
}
func (x Disease) FieldId() string {
	return config.IdStr
}
func (x Disease) FieldName() string {
	return config.NameStr
}
func (x Disease) FieldNamezh() string {
	return config.NamezhStr
}
func (x Disease) FieldPinyin() string {
	return config.PinyinStr
}
func (x Disease) FieldSourceId() string {
	return "source_id"
}

type Recipe struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Namezh   string `json:"namezh"`
	Pinyin   string `json:"pinyin"`
	SourceId int    `json:"sourceId"`
}

func (x Recipe) Table() string {
	return config.TbRecipe
}
func (x Recipe) FieldId() string {
	return config.IdStr
}
func (x Recipe) FieldName() string {
	return config.NameStr
}
func (x Recipe) FieldNamezh() string {
	return config.NamezhStr
}
func (x Recipe) FieldPinyin() string {
	return config.PinyinStr
}

type Herb struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Namezh string `json:"namezh"`
	Pinyin string `json:"pinyin"`
}

func (x Herb) Table() string {
	return config.TbHerb
}
func (x Herb) FieldId() string {
	return config.IdStr
}
func (x Herb) FieldName() string {
	return config.NameStr
}
func (x Herb) FieldNamezh() string {
	return config.NamezhStr
}
func (x Herb) FieldPinyin() string {
	return config.PinyinStr
}

type RecipeHerb struct {
	RecipeId int    `json:"recipeId"`
	HerbId   int    `json:"herbId"`
	Gram     int    `json:"gram"`
	Note     string `json:"note"`
}

type Policy struct {
	Id              int `json:"id"`
	TreatmentVerbId int `json:"treatmentVerbId"`
	OrganId         int `json:"organId"`
	TcmOrganId      int `json:"tcmOrganId"`
}

func (x Policy) Table() string {
	return config.TbPolicyStr
}
func (x Policy) FieldId() string {
	return config.IdStr
}
func (x Policy) FieldTreatVerbId() string {
	return config.TreatVerbIdStr
}
func (x Policy) FieldOrganId() string {
	return config.OrganIdStr
}
func (x Policy) FieldTcmOrganId() string {
	return config.TcmOrganIdStr
}

type DiseasePolicy struct {
	DiseaseId int `json:"diseaseId"`
	PolicyId  int `json:"policyId"`
	Priority  int `json:"priority"`
}

type OrganSymptom struct {
	Id        int `json:"id"`
	OrganId   int `json:"organId"`
	SymptomId int `json:"symptomId"`
}

func (x OrganSymptom) Table() string {
	return config.TbOrganSymptom
}
func (x OrganSymptom) FieldId() string {
	return config.IdStr
}
func (x OrganSymptom) FieldOrganId() string {
	return config.OrganIdStr
}
func (x OrganSymptom) FieldSymptomId() string {
	return config.SymptomIdStr
}

type DiseaseOrgansymptom struct {
	DiseaseId      int `json:"diseaseId"`
	OrgansymptomId int `json:"organsymptomId"`
}

func (x DiseaseOrgansymptom) Table() string {
	return config.TbDiseOrgSymp
}

func (x DiseaseOrgansymptom) FieldDiseId() string {
	return config.DiseIdStr
}

func (x DiseaseOrgansymptom) FieldOrgSympId() string {
	return config.OrgSympIdStr
}

type Pathology struct {
	Id            int `json:"id"`
	OrganId       int `json:"organId"`
	TcmOrganId    int `json:"tcmOrganId"`
	DescriptionId int `json:"descriptionId"`
}

func (x Pathology) Table() string {
	return config.TbPathologyStr
}
func (x Pathology) FieldId() string {
	return config.IdStr
}
func (x Pathology) FieldOrganId() string {
	return config.OrganIdStr
}
func (x Pathology) FieldTcmOrganId() string {
	return config.TcmOrganIdStr
}
func (x Pathology) FieldDescriptionId() string {
	return config.DescriptionIdStr
}

type Category struct {
	Id        int    `json:"id"`
	CatTypeId int    `json:"catTypeId"`
	Name      string `json:"name"`
	Namezh    string `json:"namezh"`
	Pinyin    string `json:"pinyin"`
}

func (x Category) Table() string {
	return config.TbCategoryStr
}
func (x Category) FieldId() string {
	return config.IdStr
}
func (x Category) FieldCatTypeId() string {
	return config.CatTypeIdStr
}
func (x Category) FieldName() string {
	return config.NameStr
}
func (x Category) FieldNamezh() string {
	return config.NamezhStr
}
func (x Category) FieldPinyin() string {
	return config.PinyinStr
}
