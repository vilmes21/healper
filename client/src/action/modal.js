import {SET_MODAL} from './types'

export const setModal = x => {
    return {
        type: SET_MODAL, 
        payload: x
    }
}
