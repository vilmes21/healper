import {ALL_PATHO_ADJS, ADD_PATHO_ADJ} from '../action/types'

const pathoAdjs = (state = [], action) => {
    switch (action.type) {
        case ALL_PATHO_ADJS:
            return action.payload;
        case ADD_PATHO_ADJ:
            return [action.payload, ...state];
        default:
            return state;
    }
}

export default pathoAdjs;
