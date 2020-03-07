// Copyright 2020 Paul Sitoh
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

const initialUserState = {
    fetching: false,
    fetched: false,
    users: [],
    error: null
};

const usersReducer = (state = initialUserState, action) => {
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

export default usersReducer;