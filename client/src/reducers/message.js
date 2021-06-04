import { SET_MESSAGE, CLEAR_MESSAGE } from "../actions/types";

const initialState = {};

/**
* @description message stores any updated message during redux dispatch
*/
const message = (state = initialState, action) => {
  const { type, payload } = action;

  switch (type) {
    case SET_MESSAGE:
      return { message: payload };

    case CLEAR_MESSAGE:
      return { message: "" };

    default:
      return state;
  }
}

export default message;