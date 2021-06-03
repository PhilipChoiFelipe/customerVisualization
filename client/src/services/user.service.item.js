//TODO
//IMPLEMENT LATER

import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "http://localhost:80/v1";

//get all items by user
const getAllItems = async (user_id, query) => {
  console.log(user_id, query);
  const response = await axios.get(API_URL + `/user/${user_id}/items`, { headers: authHeader() }, {params: query})
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
};

//user create item
/*
type Item struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"userId"`
	ItemName string `json:"itemName"`
	Price    int64  `json:"price"`
}
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

//user get specific item
const getSpecItem = async (user_id, item_id) => {
  const response = await axios.get(API_URL + `/user/${user_id}/items/${item_id}`, { headers: authHeader() })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

//user patch specific item
/*
type ItemUpdate struct {
	ItemName string `json:"itemName"`
	Price    int64  `json:"price"`
}
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

//user delete specific item
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