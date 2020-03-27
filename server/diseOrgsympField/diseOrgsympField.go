package diseOrgsympField

type Field int

const (
    DiseaseId Field = iota
    OrganSymptomId
)

func (d Field) Str() string {
        switch d {
    case DiseaseId:
        return `disease_id`
    case OrganSymptomId:
                return `organ_symptom_id`
    default:
        return `BAD`
    }
}