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

import React from 'react';

import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';

import Grid from '@material-ui/core/Grid';
import Button from '@material-ui/core/Button';

const useStyles = makeStyles(theme => ({
    root: {
        padding: theme.spacing(4)
    }
}));

const Users = ({fetchUser, users}) => {
    const classes = useStyles();
    
    return (
        <div className={classes.root}>
            <Grid container spacing={1}>
                <Grid item lg={3} sm={6} xl={3} xs={12}>
                    <Button onClick={
                        async () => {
                            await fetchUser();
                        }
                    }
                    >Fetch User</Button>
                </Grid>
                <Grid item xs={12}>
                    {
                        JSON.stringify(users)
                    }
                </Grid>
            </Grid>
        </div>
    );

};

Users.propTypes = {
    users: PropTypes.object,
    fetchUser: PropTypes.func
};

export default Users;