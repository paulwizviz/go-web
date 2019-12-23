// Copyright 2019 Paul Sitoh
// SPDX-License-Identifier: Apache-2.0

import { createAction } from 'redux-actions';

// Actions
const ACTION__GET_USERS = 'app/users/get';
const ACTION__ADD_USER = 'app/users/add';

export const getUsersAction = createAction(ACTION__GET_USERS);
export const addUserAction = createAction(ACTION__ADD_USER);

// reducers
const initialState = {
    user: {
        name: null
    }
};

export const usersReducer = ( state = initialState, action ) => {
    switch ( action.type ) {
    case ACTION__GET_USERS:
        return state;
    case ACTION__ADD_USER:
        state = {
            ...state,
            user: action.payload
        };
        return state;
    default:
        return state;
    }
};
