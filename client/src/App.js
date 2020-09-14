import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import {
  Container,
} from '@material-ui/core';

import Chat from './Chat';
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
    <Container maxWidth="md">
      {
        user 
          ? <Chat user={user} />
          : <p>Show registration form</p>
      }
      
    </Container>
  );
}

export default App;
