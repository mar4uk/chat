import React, { useState, useEffect, useRef } from 'react';
import axios from 'axios';
import moment from "moment";
import {
  Container,
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

function App() {
  const classes = useStyles();
  const [data, setData] = useState({ hits: [] });
  const currentUserId = 1;
  const ws = useRef(null);

  useEffect(() => {
    async function fetchData() {
      const result = await axios(
        'http://localhost:8080/chat/1/messages',
      );

      setData({
        hits: result.data
      });
    }
    fetchData()

    ws.current = new WebSocket('ws://localhost:8080/socket');
    ws.current.onopen = () => console.log("connected");
    ws.current.onclose = () => console.log("disconnected");

    return () => {
      ws.current.close();
    };
  }, []);

  useEffect(() => {
    if (!ws.current) return;

    ws.current.onmessage = evt => {
      const message = JSON.parse(evt.data)
      setData({
        hits: data.hits.concat(message),
      })
    }
  }, [ws, data.hits]);

  function onSubmit(values, { resetForm }) {
    if (!values.text) {
      return;
    }
    const message = {
      userId: 1,
      chatId: 1,
      text: values.text,
      createdAt: moment().toISOString()
    };

    ws.current.send(JSON.stringify(message));
    resetForm();
  }

  return (
    <Container maxWidth="md">
      <Grid container direction="column" className={classes.wrapper} spacing={2} wrap="nowrap">
        <Grid item className={classes.chatWrapper}>
          <Paper className={classes.chatContainer}>
            <List>
              {data.hits.map((hit) => (
                <ListItem key={hit.id} style={{justifyContent: hit.userId === currentUserId ? "flex-end" : "flex-start"}}>
                  <Box color="primary.text" p={2} boxShadow={1}>
                    <ListItemText
                      primary={
                        <Typography component="span" variant="subtitle1">
                          {hit.userId}
                        </Typography>
                      }
                      secondary={
                        <React.Fragment>
                          <Typography
                            component="span"
                            color="textPrimary"
                          >
                            {hit.text}
                          </Typography>
                          <Typography component="span" display="block" variant="body2">
                            {moment(hit.createdAt).format('LLL')}
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
      </Grid>
    </Container>
  );
}

export default App;
