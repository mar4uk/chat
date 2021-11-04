import React, { useState } from 'react';

import {
  Avatar,
  Typography,
  Container,
  Grid,
  Link,
  makeStyles
} from '@material-ui/core';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import LoginForm from './LoginForm';
import RegisterForm from './RegisterForm';

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  formContainer: {
    marginTop: theme.spacing(3),
  },
  form: {
    marginTop: theme.spacing(1)
  }
}));

function Authorization() {
  const classes = useStyles();
  const [ view, setView ] = useState('login')

  const onToggleView = () => {
    setView(view === 'login' ? 'register' : 'login')
  }
  
  return (
    <Container maxWidth="xs" component="main" className={classes.formContainer}>
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          {
            view === 'login'
              ? 'Sign in'
              : 'Sign up'
          }
        </Typography>
        <div className={classes.form}>
          {
            view === 'login'
              ? <LoginForm />
              : <RegisterForm />
          }
          <Grid container justifyContent="flex-end">
            <Grid item>
              <Link component="button" onClick={onToggleView} variant="body2">
                {
                  view === 'login'
                    ? 'Don\'t have an account? Sign Up'
                    : 'Already have an account? Sign in'
                }
              </Link>
            </Grid>
          </Grid>
        </div>
      </div>
    </Container>
  )
}

export default Authorization;
