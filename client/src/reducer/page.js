import {GOTO} from '../action/types'

const page = (state = false, action) => {
    switch (action.type) {
        case GOTO:
            window.history.pushState('page2', 'Title', `/${action.payload}`);
            return action.payload;
        default:
            return state;
    }
}


export default page
