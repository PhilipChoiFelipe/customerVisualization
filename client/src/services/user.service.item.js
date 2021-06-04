import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "https://441final-api.erinchang.me/v1";

/**
 * @description - getAllItems sends GET request to the server
 * with the current user's id and query that represents the column name to sort the entries by.
 * Returns all the items matching the provided parameters.
 */
const getAllItems = async (user_id, query) => {
  console.log(user_id, query);
  const response = await axios.get(API_URL + `/user/${user_id}/items`, { headers: authHeader() }, {params: query})
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
};

/**
 * @description - createItem sends POST request to the server
 * with the current user's id and an item object to store in the database.
 */
const createItem = async (user_id, item) => {
  axios.defaults.headers.common['Authorization'] = authHeader()["Authorization"];
  console.log("user.service.item 30", item);
  const response = await axios.post(API_URL + `/user/${user_id}/items`, 
    item,
  )
  // { 
  //   headers: authHeader(),
  //   body: JSON.stringify(item)
  // },{})
  if(response.status !== 201) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

/**
 * @description - getSpecItem sends POST request to the server
 * with the current user's id and a item ID and returns the matching user.
 */
const getSpecItem = async (user_id, item_id) => {
  const response = await axios.get(API_URL + `/user/${user_id}/items/${item_id}`, { headers: authHeader() })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

/**
 * @description - updateSpecItem sends POST request to the server
 * with the current user's id, a item ID, and an update object 
 * to update the matching item using the values inside update object.
 */
const updateSpecItem = async (user_id, item_id, item_update) => {
  axios.defaults.headers.common['Authorization'] = authHeader()["Authorization"];
  const response = await axios.patch(API_URL + `/user/${user_id}/items/${item_id}`, 
      item_update
    )
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

/**
 * @description - deleteSpecItem sends DELETE request to the server
 * with the current user's id and an item ID to delete the matching 
 * item from the database.
 */
const deleteSpecItem = async (user_id, item_id) => {
  const response = await axios.delete(API_URL + `/user/${user_id}/items/${item_id}`, { headers: authHeader() })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

export default {
  getAllItems,
  createItem,
  getSpecItem,
  updateSpecItem,
  deleteSpecItem
};