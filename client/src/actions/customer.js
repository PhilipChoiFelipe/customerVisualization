import {
    CUSTOMER_GETALL_SUCCESS,
    CUSTOMER_GETALL_FAIL,
    CUSTOMER_GETSPEC_SUCCESS,
    CUSTOMER_GETSPEC_FAIL,
    SET_MESSAGE,

  } from "../actions/types";

import CustomerService from "../services/user.service.customer";

export const getAllCustomers = (user_id) => (dispatch) => {
    return CustomerService.getAllCustomers(user_id).then(
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


export const getSpecCustomer = (user_id, item_id) => (dispatch) => {
    return CustomerService.getSpecItem(user_id, item_id).then(
        (response) =>  {
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


