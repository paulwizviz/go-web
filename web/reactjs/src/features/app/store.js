// Copyright 2019 Paul Sitoh
// SPDX-License-Identifier: Apache-2.0

import { createStore, combineReducers } from 'redux';

import { usersReducer } from '../users';

const reducers = combineReducers({usersReducer});

export const store = createStore(reducers);