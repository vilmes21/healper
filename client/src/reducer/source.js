import {ALL_SOURCES} from '../action/types'

const source = (state = [], action) => {
    switch (action.type) {
        case ALL_SOURCES:
            return action.payload;
        default:
            return state;
    }
}


export default source;
