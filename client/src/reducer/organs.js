import {ALL_ORGANS, ADD_ORGAN} from '../action/types'

const organs = (state = [], action) => {
    switch (action.type) {
        case ALL_ORGANS:
            return action.payload;
        case ADD_ORGAN:
            return [action.payload, ...state];
        default:
            return state;
    }
}


export default organs;
