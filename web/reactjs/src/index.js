import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import { applyMiddleware, createStore, combineReducers } from 'redux';
import { createLogger } from 'redux-logger';
import { createPromise } from 'redux-promise-middleware';

const promise = createPromise({ types: { fulfilled: 'success' } });

import App from './App';

const initialUserState = {
    fetching: false,
    fetched: false,
    users: [],
    error: null
};

const userReducer = (state = initialUserState, action) => {
    switch (action.type) {
    case 'FETCH_USERS_PENDING': {
        return {
            ...state,
            fetching: true
        };
    }
    case 'FETCH_USERS_FULFILLED': {
        return {
            ...state,
            fetching: false,
            fetched: true,
            users: action.payload
        };
    }
    case 'FETCH_USERS_REJECTED': {
        return {
            ...state,
            fetching: false,
            fetched: false,
            error: action.payload
        };
    }
    default:
        return state;
    }
};

const store = createStore(
    combineReducers({ userReducer }),
    {},
    applyMiddleware(promise, createLogger()));

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.getElementById('app')
);