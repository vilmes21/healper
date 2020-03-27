import {ALL_TREAT_VERBS, ADD_TREAT_VERB} from '../action/types'

const treatVerbs = (state = [], action) => {
    switch (action.type) {
        case ALL_TREAT_VERBS:
            return action.payload;
        case ADD_TREAT_VERB:
            return [action.payload, ...state];
        default:
            return state;
    }
}

export default treatVerbs;
