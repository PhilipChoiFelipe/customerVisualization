import {
  REGISTER_SUCCESS,
  REGISTER_FAIL,
  LOGIN_SUCCESS,
  LOGIN_FAIL,
  LOGOUT,
  SET_MESSAGE,
} from "./types";

import AuthService from "../services/auth.service";

/**
* @description register registers user data and dispatch result message 
*/
export const register = (userName, email, password, passwordConf, firstName, lastName, storeName) => (dispatch) => {
  return AuthService.register(userName, email, password, passwordConf, firstName, lastName, storeName).then(
    (response) => {
      dispatch({
        type: REGISTER_SUCCESS,
      });

      dispatch({
        type: SET_MESSAGE,
        payload: response.data.message,
      });

      return Promise.resolve();
    },
    (error) => {
      const message =
        (error.response &&
          error.response.data &&
          error.response.data.message) ||
        error.message ||
        error.toString();

      dispatch({
        type: REGISTER_FAIL,
      });

      dispatch({
        type: SET_MESSAGE,
        payload: message,
      });

      return Promise.reject();
    }
  );
};

/**
* @description login sends validate user sigin data and dispatch result message, user token, and user data
*/
export const login = (email, password) => (dispatch) => {
  return AuthService.login(email, password).then(
    (result) => {
      console.log("actions/auth.js 51", result);
      dispatch({
        type: LOGIN_SUCCESS,
        payload: result,
      });

      return Promise.resolve();
    },
    (error) => {
      const message =
        (error.response &&
          error.response.data &&
          error.response.data.message) ||
        error.message ||
        error.toString();

      dispatch({
        type: LOGIN_FAIL,
      });

      dispatch({
        type: SET_MESSAGE,
        payload: message,
      });

      return Promise.reject();
    }
  );
};

/**
* @description logout deletes current user's token and dispatch result message 
*/
export const logout = () => (dispatch) => {
  AuthService.logout();
  dispatch({
    type: LOGOUT,
  });
};