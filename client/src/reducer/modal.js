import {SET_MODAL} from '../action/types'

const _init = {open: false, content: -1}

const setModal = (state = _init, action) => {
    switch (action.type) {
        case SET_MODAL:
            const {content,open}=action.payload;
            const _state = {open}
            if (content){
                _state.content = content
            } else {
                _state.content = -1
            }
            return _state;
        default:
            return state;
    }
}

export default setModal;