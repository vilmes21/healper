import {ALL_ORG_SYMS} from '../action/types'

const orgSyms = (state = [], action) => {
    switch (action.type) {
        case ALL_ORG_SYMS:
            return action.payload;
        default:
            return state;
    }
}

export default orgSyms;
