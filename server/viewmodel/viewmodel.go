package viewmodel

import (
	"../config"
	"../model"
)

type JsonRes struct {
	Ok  bool   `json:"ok"`
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

type Src struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type DiseSymptom struct {
	DiseId     int                  `json:"diseId"`
	OrgsympIds []int                `json:"orgsympIds"`
	NewSymps   []model.OrganSymptom `json:"newSymps"`
}

type DiseSym struct {
	DiseId    int    `json:"diseId"`
	OrgsympId int    `json:"orgsympId"`
	OrganId   int    `json:"organId"`
	SymptomId int    `json:"symptomId"`
	Txt       string `json:"txt"`
}

type Patho struct {
	Id            int    `json:"id"`
	FrontendType  int    `json:"frontendType"`
	OrganId       int    `json:"organId"`
	TcmOrganId    int    `json:"tcmOrganId"`
	DescriptionId int    `json:"descriptionId"`
	Txt           string `json:"txt"`
	Searchable    string `json:"searchable"`
}

type Policy struct {
	Id           int    `json:"id"`
	FrontendType int    `json:"frontendType"`
	TreatVerbId  int    `json:"treatVerbId"`
	OrganId      int    `json:"organId"`
	TcmOrganId   int    `json:"tcmOrganId"`
	Txt          string `json:"txt"`
	Searchable   string `json:"searchable"`
}

type OrganSymp struct {
	Id           int    `json:"id"`
	FrontendType int    `json:"frontendType"`
	OrganId      int    `json:"organId"`
	SymptomId    int    `json:"symptomId"`
	Txt          string `json:"txt"`
	Searchable   string `json:"searchable"`
}

type Cat struct {
	Id           int                 `json:"id"`
	FrontendType config.FrontendType `json:"frontendType"`
	CatTypeId    int                 `json:"catTypeId"`
	Name         string              `json:"name"`
	Namezh       string              `json:"namezh"`
	Pinyin       string              `json:"pinyin"`
	Searchable   string              `json:"searchable"`
	Txt          string              `json:"txt"`
}

type Recipe struct {
	Id           int    `json:"id"`
	FrontendType int    `json:"frontendType"`
	Name         string `json:"name"`
	Namezh       string `json:"namezh"`
	Pinyin       string `json:"pinyin"`
	SourceId     int    `json:"sourceId"`
	Searchable   string `json:"searchable"`
	Txt          string `json:"txt"`
}

type Herb struct {
	Id           int    `json:"id"`
	FrontendType int    `json:"frontendType"`
	Name         string `json:"name"`
	Namezh       string `json:"namezh"`
	Pinyin       string `json:"pinyin"`
	Searchable   string `json:"searchable"`
	Txt          string `json:"txt"`
}

type Selectables struct {
	OrganSymptoms []OrganSymp `json:"organSymptoms"`
	Organs        []Cat       `json:"organs"`
	SymptomAdjs   []Cat       `json:"symptomAdjs"`
	Pathos        []Patho     `json:"pathos"`
	PathoAdjs     []Cat       `json:"pathoAdjs"`
	Policies      []Policy    `json:"policies"`
	TreatVerbs    []Cat       `json:"treatVerbs"`
	Recipes       []Recipe    `json:"recipes"`
	Herbs         []Herb      `json:"herbs"`
}

type SaveOrgSymptomsRes struct {
	Ok       bool      `json:"ok"`
	Msg      string    `json:"msg"`
	DiseSyms []DiseSym `json:"diseSyms"`
}

type DisePatho struct {
	DiseId     int    `json:"diseId"`
	PathoId    int    `json:"pathoId"`
	OrganId    int    `json:"organId"`
	TcmOrganId int    `json:"tcmOrganId"`
	AdjId      int    `json:"adjId"`
	Txt        string `json:"txt"`
}

type DisePolicy struct {
	Id          int    `json:"id"`
	TreatVerbId int    `json:"treatVerbId"`
	OrganId     int    `json:"organId"`
	TcmOrganId  int    `json:"tcmOrganId"`
	Txt         string `json:"txt"`
}

type DiseRecipe struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Namezh string `json:"namezh"`
	Txt    string `json:"txt"`
}

type Dise struct {
	Id       int          `json:"id"`
	Name     string       `json:"name"`
	Source   string       `json:"source"`
	Syms     []DiseSym    `json:"syms"`
	Pathos   []DisePatho  `json:"pathos"`
	Policies []DisePolicy `json:"policies"`
	Recipes  []DiseRecipe `json:"recipes"`
	Msg      string       `json:"msg"`
}
