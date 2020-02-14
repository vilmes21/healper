import page from './page'
import source from './source'
import dises from './dises'
import sympAdjs from './sympAdjs'
import organs from './organs'
import {combineReducers} from 'redux';

export default combineReducers({page, source, dises, sympAdjs, organs});

/*
store-shape == {
    page: `string`,
    soruce: [{}, {}],
    dises: [{}],
    sympAdjs: [{id: 703, catTypeId: 7, name: "painful", namezh: "痛", pinyin: "t"}],
    organs: [{id: 101, catTypeId: 1, name: "face", namezh: "脸", pinyin: "l"}],
}
*/


 