import {ALL_ORGANS, ADD_ORGAN} from './types'

export const fillOrgans = x => {
    return {
        type: ALL_ORGANS, 
        payload: x
    }
}

export const addOrgan = x => {
    return {
        type: ADD_ORGAN, 
        payload: x
    }
}