import React from 'react';
import { useDispatch } from 'react-redux';
import { Formik, Field, Form } from 'formik';
import {
  Button,
  makeStyles
} from '@material-ui/core';
import {
  Input
} from './Input';

import { loginUser } from './actions';

const useStyles = makeStyles((theme) => ({
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const initialValues = {
  email: '',
  password: ''
};

function LoginForm() {
  const classes = useStyles();
  const dispatch = useDispatch()
  const onSubmit = (values) => {
    dispatch(loginUser(values.email, values.password))
  }

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={onSubmit}
    >
      <Form>
        <Field
          id="email"
          type="email"
          name="email"
          label="Email"
          fullWidth
          margin="normal"
          component={Input}
        />
        <Field
          id="password"
          type="password"
          name="password"
          label="Password"
          fullWidth
          margin="normal"
          component={Input}
        />
        <Button
          type="submit"
          fullWidth
          variant="contained"
          color="primary"
          className={classes.submit}
        >
          Sign In
        </Button>
      </Form>
    </Formik>
  );
}

export default LoginForm;
