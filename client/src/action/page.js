import {GOTO} from './types'

export const redirect = toUrl => {
    return {
        type: GOTO, 
        payload: toUrl
    }
}

// export const redirect2 = e => {
//     return async (dispatch) => {
//         try {
//             return dispatch({
//                 type: GOTO, 
//                 payload: `new-disease`
//             })
//         } catch (e) {
//             console.log("actions/page.js redirect e: ", e)
//         }
//     }
// }