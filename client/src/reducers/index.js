import { combineReducers } from "redux";
import auth from "./auth";
import item from './item';
import customer from './customer';
import message from "./message";

export default combineReducers({
  auth,
  message,
  customer,
  item
  // smt
});