import {ALL_SYMP_ADJS, ADD_SYM_ADJ} from '../action/types'

const sympAdjs = (state = [], action) => {
    switch (action.type) {
        case ALL_SYMP_ADJS:
            return action.payload;
        case ADD_SYM_ADJ:
            return [action.payload, ...state];
        default:
            return state;
    }
}

export default sympAdjs;
