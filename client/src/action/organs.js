import {ALL_ORGANS} from './types'

export const fillOrgans = x => {
    return {
        type: ALL_ORGANS, 
        payload: x
    }
}
