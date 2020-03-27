import {ADD_DISE, ADD_SYM_TO_DISE} from '../action/types'

const dises = (state = [], action) => {
    switch (action.type) {
        case ADD_DISE:
            return [...state, action.payload];
        case ADD_SYM_TO_DISE:
            const {diseId, syms}=action.payload;
            const dise = state.find(x => x.id === diseId);
            const dise2 = {...dise, syms }
            const state2 = state.filter(x => x.id !== diseId);
            state2.push(dise2)
            return state2;
        default:
            return state;
    }
}

export default dises;
