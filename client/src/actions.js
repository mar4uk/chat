import axios from 'axios';
import {
  FETCH_USER,
  FETCH_MESSAGES,
  POST_MESSAGE,
  LOGIN_USER
} from './actionTypes';

export const fetchUser = () => {
  return (dispatch) => {
    const cookies = document.cookie.split('; ').reduce((memo, cookie) => {
      const [key, value] = cookie.split('=');
      if (key && value) {
        memo[key] = value;
      }
      return memo;
    }, {});
    
    const { jwt } = cookies;

    if (!jwt) {
      dispatch({
        type: FETCH_USER,
        payload: null
      });
      
      return;
    }

    return axios({
      url: 'http://localhost:8080/user',
      method: 'GET',
      headers: {
        Authorization: `Bearer ${jwt}`
      }
    }).then((user) => {
      dispatch({
        type: FETCH_USER,
        payload: {
          ...user.data,
          jwt
        }
      });
    }).catch(() => {
      dispatch({
        type: FETCH_USER,
        payload: null
      });
    });
  }
}

export const fetchMessages = () => {
  return (dispatch, getState) => {
    const state = getState();
    const jwt = state.user && state.user.jwt;

    if (!jwt) {
      dispatch({
        type: FETCH_MESSAGES,
        payload: []
      });
    }

    return axios({
      url: 'http://localhost:8080/chat/1/messages',
      method: 'GET',
      headers: {
        Authorization: `Bearer ${jwt}`
      }
    }).then((messages) => {
      dispatch({
        type: FETCH_MESSAGES,
        payload: messages.data
      });
    }).catch(() => {
      dispatch({
        type: FETCH_MESSAGES,
        payload: []
      });
    })
  }
}

export const sendMessage = (message) => {
  return (dispatch) => {
    dispatch({
      type: POST_MESSAGE,
      payload: message
    });
  }
}

export const loginUser = (email, password) => {
  return (dispatch) => {
    return axios({
      url: 'http://localhost:8080/login',
      method: 'POST',
      data: {
        email,
        password
      }
    }).then(({ data }) => {
      dispatch({
        type: LOGIN_USER,
        payload: {
          ...data.user,
          jwt: data.token
        }
      });
      document.cookie = `jwt=${data.token}`
    }).catch(() => {
      dispatch({
        type: LOGIN_USER,
        payload: null
      });
    });
  }
}