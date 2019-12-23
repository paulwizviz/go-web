// Copyright 2019 Paul Sitoh
// SPDX-License-Identifier: Apache-2.0

import { createMuiTheme } from '@material-ui/core';

import { blue } from '@material-ui/core/colors';

const theme = createMuiTheme({
    palette: {
        primary: blue,
    },
    typography: {
        useNextVariants: true,
    },
});

export default theme;