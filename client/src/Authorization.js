import React from 'react';
import { useDispatch } from 'react-redux';
import { Formik, Field, Form } from 'formik';

import { loginUser } from './actions';

const initialValues = {
  email: '',
  password: ''
};

function Authorization() {
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
        <Field id="email" type="email" name="email" placeholder="Email" />
        <Field id="password" type="password" name="password" placeholder="Password" />
        <button type="submit">Submit</button>
      </Form>
    </Formik>
  );
}

export default Authorization;
