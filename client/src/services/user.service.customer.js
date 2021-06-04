import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "https://441final-api.erinchang.me/v1";

/**
 * @description - getAllCustomers sends GET request to the server
 * with the current user's id and query that represents the column name to sort the entries by.
 * Returns all the customers matching the provided parameters.
 */
const getAllCustomers = async (user_id, query = {}) => {
  const response = await axios.get(API_URL + `/user/${user_id}/customers`, { headers: authHeader(), params: query })
  if (response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
};

/**
 * @description - createCustomer sends POST request to the server
 * with the current user's id and a customer object to store in the database.
 */
const createCustomer = async (user_id, customer) => {
  axios.defaults.headers.common['Authorization'] = authHeader()["Authorization"];
  console.log(customer);
  const response = await axios.post(API_URL + `/user/${user_id}/customers`,
    customer
  )
  if (response.status !== 201) {
    console.log(response.statusText);
  } else {
    console.log("user.service.customer 44", response)
    return response.data;
  }
}

/**
 * @description - getSpecCustomer sends POST request to the server
 * with the current user's id and a customer ID and returns the matching user.
 */
const getSpecCustomer = async (user_id, customer_id) => {
  const response = await axios.post(API_URL + `/user/${user_id}/customers/${customer_id}`, { headers: authHeader() })
  if (response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

/**
 * @description - updateSpecCustomer sends POST request to the server
 * with the current user's id, a customer ID, and an update object 
 * to update the matching customer using the values inside update object.
 */
const updateSpecCustomer = async (user_id, customer_id, update) => {
  console.log(user_id, customer_id, update)
  axios.defaults.headers.common['Authorization'] = authHeader()["Authorization"];
  const response = await axios.patch(API_URL + `/user/${user_id}/customers/${customer_id}`,
    update
  )
  if (response.status !== 200) {
    console.log(response.statusText);
  } else {
    return response.data;
  }
}

/**
 * @description - deleteSpecCustomer sends DELETE request to the server
 * with the current user's id and a customer ID to delete the matching 
 * customer from the database.
 */
const deleteSpecCustomer = async (user_id, customer_id) => {
  const response = await axios.delete(API_URL + `/user/${user_id}/customers/${customer_id}`, { headers: authHeader() })
  if (response.status !== 200) {
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