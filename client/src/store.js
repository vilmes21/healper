import { createStore, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'
import rootReducer from './reducer/index'

const store = createStore(rootReducer, applyMiddleware(thunk))

store.subscribe(() => {
    console.log("subscribe fn. store getState >>> ", store.getState())
})

export default store

/*
store structure:

{
    page: ``,
    isLoading: true,
    appointments: {},
    currentUser: {},
    doctors: [{}, {}],
    doctorUrls: {}
}

*/