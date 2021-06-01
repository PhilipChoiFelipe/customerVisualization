//TODO
//IMPLEMENT LATER

import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "http://localhost:80/v1";

//get all customers by user
const getAllCustomers = async (user_id) => {
  const response = await axios.get(API_URL + `/user/${user_id}/customers`, { headers: authHeader() })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
};

//user create customer
/*
type Customer struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"userId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Ethnicity  string `json:"ethnicity"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	PostalCode int64  `json:"postalCode"`
	LastVisted string `json:"lastVisited"`
	DisChannel string `json:"disChannel"`
	FavItem    int64  `json:"favItem"`
}
*/ 
const createCustomer = async (user_id, customer) => {
  const response = await axios.post(API_URL + `/user/${user_id}/customers`, 
  { 
    headers: authHeader(),
    body: JSON.stringify(customer)
  })
  if(response.status !== 201) {
    console.log(response.statusText);
  } else {
    console.log("user.service.customer 44", response)
    return response.data;
  }
}

//user get specific customer
// user/{user_id}/customers/{customer_id}
const getSpecCustomer = async (user_id, customer_id) => {
  const response = await axios.post(API_URL + `/user/${user_id}/customers/${customer_id}`, { headers: authHeader() })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

//user patch specific customer
// /user/{user_id}/customers/{customer_id}
/*
type NameUpdates struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
*/
const updateSpecCustomer = async (user_id, customer_id, nameUpdate) => {
  const response = await axios.patch(API_URL + `/user/${user_id}/customers/${customer_id}`, 
    { 
      body:JSON.stringify(nameUpdate), 
      headers: authHeader() 
    })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

//user delete specific item
const deleteSpecCustomer = async (user_id, customer_id) => {
  const response = await axios.delete(API_URL + `/user/${user_id}/customers/${customer_id}`, { headers: authHeader() })
  if(response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}




export default {
  getAllCustomers,
  createCustomer,
  getSpecCustomer,
  updateSpecCustomer,
  deleteSpecCustomer
};