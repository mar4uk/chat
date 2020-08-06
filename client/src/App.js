import React, { useState, useEffect } from 'react';
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
  makeStyles
} from '@material-ui/core';

const useStyles = makeStyles({
  wrapper: {
    height: "100vh",
    overflowY: "scroll"
  },
});

function App() {
  const classes = useStyles();
  const [data, setData] = useState({ hits: [] });
  const currentUserId = 1;

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
  });

  return (
    <Container maxWidth="md">
      <Paper className={classes.wrapper}>
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
    </Container>
  );
}

export default App;
