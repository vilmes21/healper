import {ADD_DISE} from '../action/types'

const dises = (state = [], action) => {
    switch (action.type) {
        case ADD_DISE:
            return [...state, action.payload];
        default:
            return state;
    }
}

export default dises;
