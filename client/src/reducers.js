import {
  FETCH_MESSAGES,
  FETCH_USER,
  POST_MESSAGE
} from './actionTypes';

const defaultState = {
  user: null,
  messages: []
};

export const rootReducer = (state = defaultState, action) => {
  switch (action.type) {
    case FETCH_USER:
      return { ...state, user: action.payload }
    case FETCH_MESSAGES:
      return { ...state, messages: action.payload };
    case POST_MESSAGE:
      return { ...state, messages: [ ...state.messages, action.payload ] };
    default:
      return state;
  }
};