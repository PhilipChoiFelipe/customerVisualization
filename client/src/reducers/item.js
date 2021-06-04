import {
  ITEM_GETALL_SUCCESS,
  ITEM_GETALL_FAIL,
  ITEM_GETSPEC_SUCCESS,
  ITEM_GETSPEC_FAIL,
} from "../actions/types";

/**
* @description initialstate is redux storage for state management. 
*/
const initialState = {
  items: [],
  specItem: null
}

/**
* @description auth updates initial tate with different incoming dispatch with actions
*/
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
        specItem: payload
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