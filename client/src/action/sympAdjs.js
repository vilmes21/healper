import {ALL_SYMP_ADJS, ADD_SYM_ADJ} from './types'

export const fillSympAdjs = x => {
    return {
        type: ALL_SYMP_ADJS, 
        payload: x
    }
}

export const addSymAdj = x => {
    return {
        type: ADD_SYM_ADJ, 
        payload: x
    }
}
