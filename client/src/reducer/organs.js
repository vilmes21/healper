import {ALL_ORGANS} from '../action/types'

const organs = (state = [], action) => {
    switch (action.type) {
        case ALL_ORGANS:
            return action.payload;
        default:
            return state;
    }
}


export default organs;
