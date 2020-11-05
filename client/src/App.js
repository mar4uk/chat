import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import {
  CssBaseline,
} from '@material-ui/core';

import Chat from './Chat';
import Authorization from './Authorization';
import { fetchUser } from './actions';

function App() {
  const {
    user
  } = useSelector((state) => ({
    user: state.user
  }));

  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(fetchUser());
  }, [dispatch])
  
  return (
    <>
      <CssBaseline />
      {
        user 
          ? <Chat user={user} />
          : <Authorization />
      }
    </>
  );
}

export default App;
