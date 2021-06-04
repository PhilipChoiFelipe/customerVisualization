import { combineReducers } from "redux";
import auth from "./auth";
import item from './item';
import customer from './customer';
import message from "./message";

/**
* @description combineReducer combines all decalred reducers. 
*/
export default combineReducers({
  auth,
  message,
  customer,
  item
});