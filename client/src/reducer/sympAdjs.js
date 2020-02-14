import {ALL_SYMP_ADJS} from '../action/types'

const sympAdjs = (state = [], action) => {
    switch (action.type) {
        case ALL_SYMP_ADJS:
            return action.payload;
        default:
            return state;
    }
}


export default sympAdjs;
