import {
  ITEM_GETALL_SUCCESS,
  ITEM_GETALL_FAIL,
  ITEM_GETSPEC_SUCCESS,
  ITEM_GETSPEC_FAIL,

  SET_MESSAGE,
} from "./types";

import ItemService from "../services/user.service.item";

/**
* @description get all itmes with userid and query and dispatch result message and returning items. 
*/
export const getAllItems = (user_id, query) => (dispatch) => {
  return ItemService.getAllItems(user_id, query).then(
    (result) => {
      console.log("actions/item.js 16", result);
      dispatch({
        type: ITEM_GETALL_SUCCESS,
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
        type: ITEM_GETALL_FAIL,
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
* @description get one item with userid and query and dispatch result message and returning item. 
*/
export const getSpecItem = (user_id, item_id) => (dispatch) => {
  return ItemService.getSpecItem(user_id, item_id).then(
    (response) => {
      console.log("actions/item.js 28", response)
      dispatch({
        type: ITEM_GETSPEC_SUCCESS,
        payload: response
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
        type: ITEM_GETSPEC_FAIL,
      });

      dispatch({
        type: SET_MESSAGE,
        payload: message,
      });

      return Promise.reject();
    }
  );
};


