import {ALL_PATHOS} from '../action/types'

const pathos = (state = [], action) => {
    switch (action.type) {
        case ALL_PATHOS:
            return action.payload;
        default:
            return state;
    }
}

export default pathos;
