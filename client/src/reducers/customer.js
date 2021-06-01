import {
    CUSTOMER_GETALL_SUCCESS,
    CUSTOMER_GETALL_FAIL,
    CUSTOMER_GETSPEC_SUCCESS,
    CUSTOMER_GETSPEC_FAIL
  } from "../actions/types";
  
  //state
  const initialState = {
      customers: [],
      specCustomer: null
  }
  
  //reducer
  //action = describes type of dispatch
  const customer = (state = initialState, action) => {
    const { type, payload } = action;
  
    switch (type) {
      case CUSTOMER_GETALL_SUCCESS:
        return {
            ...state,
            customers: payload
        };
      case CUSTOMER_GETALL_FAIL:
        return {
            ...state,
            customers: null
        };
      case CUSTOMER_GETSPEC_SUCCESS:
        return {
          ...state,
          specCustomer: payload
        };
      case CUSTOMER_GETSPEC_FAIL:
        return {
          ...state,
          specCustomer: null
        };
      default:
        return state;
    }
  }

  export default customer;