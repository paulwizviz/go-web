// Copyright 2019 Paul Sitoh
// SPDX-License-Identifier: Apache-2.0

import React from 'react';
import { Router } from 'react-router-dom';

import { Provider } from 'react-redux';

import { ThemeProvider } from '@material-ui/styles';

import { createBrowserHistory } from 'history';
const browserHistory = createBrowserHistory();

import theme from './theme';
import Routes from './Routes';
import { store } from './store';

export default class App extends React.Component {
    render() {
        return (
            <Provider store={store}>
                <ThemeProvider theme={theme}>
                    <Router history={browserHistory}>
                        <Routes />
                    </Router>
                </ThemeProvider>
            </Provider>
        );
    }
}
