import {
  CUSTOMER_GETALL_SUCCESS,
  CUSTOMER_GETALL_FAIL,
  CUSTOMER_GETSPEC_SUCCESS,
  CUSTOMER_GETSPEC_FAIL,
  SET_MESSAGE,

} from "../actions/types";

import CustomerService from "../services/user.service.customer";

/**
* @description getAllCustomers query customers with userid and query and dispatch result message and customer data
*/
export const getAllCustomers = (user_id, query) => (dispatch) => {
  return CustomerService.getAllCustomers(user_id, query).then(
    (result) => {
      console.log("actions/item.js 16", result);
      dispatch({
        type: CUSTOMER_GETALL_SUCCESS,
        payload: result
      })

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
        type: CUSTOMER_GETALL_FAIL,
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
* @description getSpecCustomer query one customer with userid and customerid and dispatch result message and one customer data
*/
export const getSpecCustomer = (user_id, customer_id) => (dispatch) => {
  return CustomerService.getSpecCustomer(user_id, customer_id).then(
    (response) => {
      console.log("actions/item.js 28", response)
      dispatch({
        type: CUSTOMER_GETSPEC_SUCCESS
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
        type: CUSTOMER_GETSPEC_FAIL,
      });

      dispatch({
        type: SET_MESSAGE,
        payload: message,
      });

      return Promise.reject();
    }
  );
};


