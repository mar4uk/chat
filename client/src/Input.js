import React from 'react';
import {
  TextField
} from '@material-ui/core';

export const Input = ({ field, form, ...props }) => {
  return (
    <TextField
      id="standard-basic"
      variant="outlined"
      {...field}
      {...props}
    />
  );
}
