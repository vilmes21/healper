import {ALL_PATHO_ADJS, ADD_PATHO_ADJ} from '../action/types'

export const fillPathoAdjs = x => {
    return {
        type: ALL_PATHO_ADJS, 
        payload: x
    }
}

export const addPathoAdj = x => {
    return {
        type: ADD_PATHO_ADJ, 
        payload: x
    }
}
