const config = {
    modalKid: {
        organForm: 1,
        symAdjForm: 2,
        pathoAdjForm: 3,
    },
    frontendType: {
        organ: 1,
        symptomAdj: 2,
        organSymptom: 3,
        patho: 4,
        policy: 5,
        recipe: 6,
        herb: 7,
        pathoAdj: 8,
        treatVerb: 9
    },
    catnum: {
        noOrganId: 117
    },
    categoryType: {
        organ: 1,
        sympAdj: 7,
        pathoAdj: 5,
        treatVerb: 6
    },
    urls: {
        newDisease: `new-disease`,
        viewDise: `disease/`
    },
    apiUrl: {
        createDisease: `/disease/create`,
        createSymps: `/organ-symptom/create`,
        getSelectables: `/selectables`,
        createCategory: `/category/create`,
        getDise: `/disease/`
    },
    host: window
        .location
        .href
        .indexOf("1234") > -1
        ? "http://127.0.0.1:3000"
        : ""
}

export default config;