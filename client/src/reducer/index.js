import page from './page'
import source from './source'
import dises from './dises'
import sympAdjs from './sympAdjs'
import pathoAdjs from './pathoAdjs'
import pathos from './pathos'
import treatVerbs from './treatVerbs'
import organs from './organs'
import modal from './modal'
import orgSyms from './orgSyms'
import {combineReducers} from 'redux';

export default combineReducers({
    pathos,
    page,
    treatVerbs,
    modal,
    pathoAdjs,
    source,
    dises,
    sympAdjs,
    organs,
    orgSyms
});

/*
store-shape == {
    page: `string`,
    modal: {
        open: false,
        content: -1
    },
    soruce: [{}, {}],
    dises: [{}],
    symptoms: [{id: 703, catTypeId: 7, name: "painful", namezh: "痛", pinyin: "t"},
{id: 101, catTypeId: 1, name: "face", namezh: "脸", pinyin: "l"}
]

}
*/
