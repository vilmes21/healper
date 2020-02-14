const config = {
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
        createDisease: `/disease/create`
    },
    host: window
         .location
         .href
         .indexOf("1234") > -1
         ? "http://127.0.0.1:3000"
         : "", 
}

export default config;