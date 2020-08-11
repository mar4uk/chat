import React from 'react';
import { Formik, Field, Form } from 'formik';
import axios from 'axios';
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

function SendMessageForm() {
  async function onSubmit (values, { resetForm }) {
    if (!values.text) {
      return;
    }
    try {
      await axios.post(
        'http://localhost:8080/chat/1/messages',
        {
          userId: 1,
          text: values.text,
          createdAt: "2020-01-02T15:04:07-0700"
        }
      );
      resetForm(initialValues);
    } catch (err) {
      alert('ERROR!')
    }
  }

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
