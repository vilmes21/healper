import {ALL_TREAT_VERBS, ADD_TREAT_VERB} from '../action/types'

export const fillTreatVerbs = x => {
    return {
        type: ALL_TREAT_VERBS, 
        payload: x
    }
}

export const addTreatVerb = x => {
    return {
        type: ADD_TREAT_VERB, 
        payload: x
    }
}
