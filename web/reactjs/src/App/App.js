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

import React, {useState} from 'react';

import { 
    makeStyles,
    Button
} from '@material-ui/core';

import axios from 'axios';

const useStyles = makeStyles((theme) => ({
    root: {
        '& > *': {
            margin: theme.spacing(1),
            width: '25ch',
        },
    },
}));

const App = () =>{
    const classes = useStyles();

    const [result, setResult] = useState('Waiting for results ...');

    const sendHandler = async () => {
        try{
            const resp = await axios.post('/api/auth',{'id':'id', 'secrets':'secrets'}, {timeout: 1000});
            setResult(resp.data);
        } catch (err){
            setResult(err);
        }
    };

    return (
        <div classes={classes}>
            <Button onClick={sendHandler}>Click me to get result from backend</Button>
            <h1>{ 
                `${JSON.stringify(result)}`
            }</h1>
        </div>
    );
};

export default App;