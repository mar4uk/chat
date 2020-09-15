import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import {
  Container,
  makeStyles
} from '@material-ui/core';

import Chat from './Chat';
import Authorization from './Authorization';
import { fetchUser } from './actions';

const useStyles = makeStyles({
  wrapper: {
    height: "100vh",
  }
});

function App() {
  const classes = useStyles()
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
    <Container maxWidth="md" className={classes.wrapper}>
      {
        user 
          ? <Chat user={user} />
          : <Authorization />
      }
      
    </Container>
  );
}

export default App;
