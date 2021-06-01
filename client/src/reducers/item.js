import {
    ITEM_GETALL_SUCCESS,
    ITEM_GETALL_FAIL,
    ITEM_GETSPEC_SUCCESS,
    ITEM_GETSPEC_FAIL,
  } from "../actions/types";
  
  //state
  const initialState = {
      items: [],
      specItem: null
  }
  
  //reducer
  //action = describes type of dispatch
  const item = (state = initialState, action) => {
    const { type, payload } = action;
  
    switch (type) {
      case ITEM_GETALL_SUCCESS:
        return {
          ...state,
          items: payload
        };
      case ITEM_GETALL_FAIL:
        return {
            ...state,
          items: null
        };
      case ITEM_GETSPEC_SUCCESS:
        return {
            ...state,
          specItem: payload.item
        };
      case ITEM_GETSPEC_FAIL:
        return {
            ...state,
          specItem: null
        };
      default:
        return state;
    }
  }

  export default item;