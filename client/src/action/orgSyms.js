import {ALL_ORG_SYMS} from './types'

export const fillOrgSyms = x => {
    return {
        type: ALL_ORG_SYMS, 
        payload: x
    }
}
