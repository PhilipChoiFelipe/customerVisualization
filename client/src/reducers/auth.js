import {
  REGISTER_SUCCESS,
  REGISTER_FAIL,
  LOGIN_SUCCESS,
  LOGIN_FAIL,
  LOGOUT,
} from "../actions/types";

/**
* @description Get user and token value from local storage
*/
const user = JSON.parse(localStorage.getItem("user"));
const token = localStorage.getItem("token");

/**
* @description initialstate is redux storage for state management. 
*/
const initialState = user
  ? { isLoggedIn: true, user, token }
  : { isLoggedIn: false, user: null, token: null };


/**
* @description auth updates initial state with different incoming dispatch with actions
*/
const auth = (state = initialState, action) => {
  const { type, payload } = action;

  switch (type) {
    case REGISTER_SUCCESS:
      return {
        ...state,
        isLoggedIn: false,
      };
    case REGISTER_FAIL:
      return {
        ...state,
        isLoggedIn: false,
      };
    case LOGIN_SUCCESS:
      return {
        ...state,
        isLoggedIn: true,
        user: payload.user,
        token: payload.token
      };
    case LOGIN_FAIL:
      return {
        ...state,
        isLoggedIn: false,
        user: null,
        token: null,
      };
    case LOGOUT:
      return {
        ...state,
        isLoggedIn: false,
        user: null,
        token: null,
      };
    default:
      return state;
  }
}

export default auth;