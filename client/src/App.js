import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import {
  Container,
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
    <Container maxWidth="md">
      {
        user 
          ? <Chat user={user} />
          : <Authorization />
      }
      
    </Container>
  );
}

export default App;
