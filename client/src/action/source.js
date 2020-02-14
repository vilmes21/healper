import {ALL_SOURCES} from './types'

export const getAll = x => {
    return {
        type: ALL_SOURCES, 
        payload: x
    }
}
