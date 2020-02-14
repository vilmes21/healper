import {ADD_DISE} from './types'

export const addDise = x => {
    return {
        type: ADD_DISE, 
        payload: x
    }
}
