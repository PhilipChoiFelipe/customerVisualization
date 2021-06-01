import {
    REGISTER_SUCCESS,
    REGISTER_FAIL,
    LOGIN_SUCCESS,
    LOGIN_FAIL,
    LOGOUT,
  } from "../actions/types";
  
  const user = JSON.parse(localStorage.getItem("user"));
  const token = localStorage.getItem("token");
  
  //state
  const initialState = user
    ? { isLoggedIn: true, user, token}
    : { isLoggedIn: false, user: null, token: null};
  
  //reducer
  //action = describes type of dispatch
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