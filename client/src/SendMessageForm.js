import React from 'react';
import { Formik, Field, Form } from 'formik';
import {
  TextField
} from '@material-ui/core';

const MessageInput = ({ field, form, ...props }) => {
  return (
    <TextField
      style={{ width: "100%" }}
      placeholder="Type a message..."
      multiline
      rowsMax={4}
      variant="outlined"
      onKeyPress={(e) => {
        if (e.which === 13 && !e.shiftKey) {
          e.preventDefault();
          form.submitForm();
        }
      }}
      {...field}
      {...props}
    />
  );
}

const initialValues = {
  text: ''
};

function SendMessageForm({ onSubmit }) {
  return (
    <Formik
      initialValues={initialValues}
      onSubmit={onSubmit}
    >
      <Form>
        <Field id="text" name="text" component={MessageInput} />
      </Form>
    </Formik>
  );
}

export default SendMessageForm;
