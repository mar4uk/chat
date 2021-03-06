import {
  FETCH_MESSAGES,
  FETCH_USER,
  POST_MESSAGE,
  LOGIN_USER,
  REGISTER_USER
} from './actionTypes';

const defaultState = {
  user: null,
  messages: []
};

export const rootReducer = (state = defaultState, action) => {
  switch (action.type) {
    case FETCH_USER:
      return { ...state, user: action.payload };
    case FETCH_MESSAGES:
      return { ...state, messages: action.payload };
    case POST_MESSAGE:
      return { ...state, messages: [ ...state.messages, action.payload ] };
    case LOGIN_USER:
      return { ...state, user: action.payload };
    case REGISTER_USER:
      return { ...state, user: action.payload };
    default:
      return state;
  }
};