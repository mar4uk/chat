import React, { useEffect, useRef } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import moment from "moment";
import {
  Typography,
  Paper,
  List,
  ListItem,
  ListItemText,
  Box,
  Grid,
  makeStyles
} from '@material-ui/core';

import SendMessageForm from './SendMessageForm';
import { fetchMessages, sendMessage } from './actions';

const useStyles = makeStyles({
  wrapper: {
    height: "100vh",
  },
  chatWrapper: {
    display: "flex",
    flex: "1 0 80%",
    minHeight: 0,
  },
  formWrapper: {
    display: "flex",
    flex: "0 0 auto",
    minHeight: 0,
  },
  chatContainer: {
    overflowY: "scroll",
    height: "100%",
    width: "100%",
    display: "flex",
    flexDirection: "column-reverse",
  },
  formContainer: {
    overflowY: "scroll",
    height: "100%",
    width: "100%",
  },
});

function Chat({ user }) {
  const classes = useStyles();
  const ws = useRef(null);
  const {
    messages
  } = useSelector((state) => ({
    messages: state.messages,
  }));

  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(fetchMessages());
  }, [dispatch]);

  useEffect(() => {
    ws.current = new WebSocket(`ws://localhost:8080/socket?jwt=${user.jwt}`);
    ws.current.onopen = () => console.log("connected");
    ws.current.onclose = () => console.log("disconnected");

    return () => {
      ws.current.close();
    };
  }, [user.jwt])

  useEffect(() => {
    if (!ws.current) return;

    ws.current.onmessage = evt => {
      const message = JSON.parse(evt.data)
      dispatch(sendMessage(message));
    }
  }, [ws, dispatch]);

  function onSubmit(values, { resetForm }) {
    if (!values.text) {
      return;
    }
    const message = {
      chatId: 1,
      text: values.text,
      createdAt: moment().toISOString()
    };

    ws.current.send(JSON.stringify(message));
    resetForm();
  }

  return (
    <>
      <Grid item className={classes.chatWrapper}>
        <Paper className={classes.chatContainer}>
          <List>
            {messages.map((message) => (
              <ListItem key={message.id} style={{justifyContent: message.user.id === user.id ? "flex-end" : "flex-start"}}>
                <Box color="primary.text" p={2} boxShadow={1}>
                  <ListItemText
                    primary={
                      <Typography component="span" variant="subtitle1">
                        {message.user.name}
                      </Typography>
                    }
                    secondary={
                      <React.Fragment>
                        <Typography
                          component="span"
                          color="textPrimary"
                        >
                          {message.text}
                        </Typography>
                        <Typography component="span" display="block" variant="body2">
                          {moment(message.createdAt).format('LLL')}
                        </Typography>
                      </React.Fragment>
                    }
                  />
                </Box>
              </ListItem>
            ))}
          </List>
        </Paper>
      </Grid>
      <Grid item className={classes.formWrapper}>
        <Paper className={classes.formContainer}>
          <Box p={2}>
            <SendMessageForm onSubmit={onSubmit} />
          </Box>
        </Paper>
      </Grid>
    </>
  );
}

export default Chat;
