package config

const (
	OrganCatType            = 1
	HerbPeerRelationCatType = 2
	PathoAdjCatType         = 5
	TreatmentVerbCatType    = 6
	SymptomAdjCatType       = 7

	NoSourceId       = 1
	BlankOrganId     = 117
	EmptyStr         = ``
	SpaceStr         = ` `
	TbDiseOrgSymp    = `disease_organsymptom`
	DiseIdStr        = `disease_id`
	OrgSympIdStr     = `organ_symptom_id`
	TbOrganSymptom   = `organ_symptom`
	OrganIdStr       = `organ_id`
	TcmOrganIdStr    = `tcm_organ_id`
	DescriptionIdStr = `description_id`
	SymptomIdStr     = `symptom_id`
	TbCategoryStr    = `category`
	IdStr            = `id`
	NameStr          = `name`
	NamezhStr        = `name_zh`
	PinyinStr        = `pinyin`
	CatTypeIdStr     = `category_type_id`
	TbPathologyStr   = `pathology`
	TbPolicyStr      = `policy`
	TreatVerbIdStr   = `treatment_verb_id`
	TbRecipe         = `recipe`
	TbHerb           = `herb`
	NotExist         = `Not exist`
	TbDisease        = `disease`
)

type SqlAction int

const (
	NoType SqlAction = iota
	SelectType
	InsertType
)

type FrontendType int

const (
	Organ        FrontendType = 1
	SymptomAdj                = 2
	OrganSymptom              = 3
	Patho                     = 4
	Policy                    = 5
	Recipe                    = 6
	Herb                      = 7
	PathoAdj                  = 8
	TreatVerb                 = 9
)
