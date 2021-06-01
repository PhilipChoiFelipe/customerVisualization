//TODO
//IMPLEMENT LATER

import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "http://localhost:80/v1";

//get all items by user
const getAllStores = async (user_id, store_id) => {
  const response = await axios.get(API_URL + `/user/${user_id}/stores`, { headers: authHeader() })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
};

//user create store
/*
type Store struct {
	ID        int64  `json:"id"`
	UserID		int64 `json:"userId"`
	StoreName     string `json:"storeName"`
	StoreLocation string `json:"storeLocation"`
}
*/ 
const createStore = async (user_id, store) => {
  const response = await axios.post(API_URL + `/user/${user_id}/stores`, 
  { 
    headers: authHeader(),
    body: JSON.stringify(store)
  })
  if(response.status !== 201) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

//user get specific store
const getSpecStore = async (user_id, store_id) => {
  const response = await axios.get(API_URL + `/user/${user_id}/stores/${store_id}`, { headers: authHeader() })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

//user patch specific store
/*
type StoreUpdate struct {
	StoreName     string `json:"storeName"`
	StoreLocation string `json:"storeLocation"`
}
*/
const updateSpecStore = async (user_id, store_id, store_update) => {
  const response = await axios.patch(API_URL + `/user/${user_id}/stores/${store_id}`, 
    { 
      body:JSON.stringify(store_update), 
      headers: authHeader() 
    })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

//user delete specific store
const deleteSpecStore = async (user_id, store_id) => {
  const response = await axios.delete(API_URL + `/user/${user_id}/stores/${store_id}`, { headers: authHeader() })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}




export default {
  getAllStores,
  getSpecStore,
  updateSpecStore,
  deleteSpecStore,
  createStore
};