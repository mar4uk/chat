import React from 'react';
import { useDispatch } from 'react-redux';
import { Formik, Field, Form } from 'formik';

import { registerUser } from './actions';

const initialValues = {
  name: '',
  email: '',
  password: ''
};

function RegisterForm() {
  const dispatch = useDispatch()
  const onSubmit = (values) => {
    dispatch(registerUser(values))
  }

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={onSubmit}
    >
      <Form>
        <Field id="name" name="name" placeholder="Name" />
        <Field id="email" type="email" name="email" placeholder="Email" />
        <Field id="password" type="password" name="password" placeholder="Password" />
        <button type="submit">Submit</button>
      </Form>
    </Formik>
  );
}

export default RegisterForm;
