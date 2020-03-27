import {ADD_DISE, ADD_SYM_TO_DISE} from '../action/types'

export const addDise = x => {
    return {
        type: ADD_DISE, 
        payload: x
    }
}

export const addSym2Dise = x => {
    return {
        type: ADD_SYM_TO_DISE, 
        payload: x
    }
}

