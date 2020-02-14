import {ALL_SYMP_ADJS} from './types'

export const fillSympAdjs = x => {
    return {
        type: ALL_SYMP_ADJS, 
        payload: x
    }
}
