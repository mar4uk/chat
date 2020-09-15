import React, { useState } from 'react';

import {
  Paper,
  Button,
  makeStyles
} from '@material-ui/core';
import LoginForm from './LoginForm';
import RegisterForm from './RegisterForm';

const useStyles = makeStyles({
  container: {
    height: "100%",
    width: "100%",
  }
});

function Authorization() {
  const classes = useStyles();
  const [ view, setView ] = useState('login')

  const onToggleView = () => {
    setView(view === 'login' ? 'register' : 'login')
  }
  
  return (
    <Paper className={classes.container}>
      {
        view === 'login'
          ? <LoginForm />
          : <RegisterForm />
      }
      <Button onClick={onToggleView}>
        {
          view === 'login'
            ? 'Register'
            : 'Login'
        }
      </Button>
    </Paper>
  )
}

export default Authorization;
