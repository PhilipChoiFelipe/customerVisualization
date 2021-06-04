import axios from "axios";
import authHeader from "./auth-header";
const API_URL = "https://441final-api.erinchang.me/v1";

/**
 * @description - register sends POST request to the server with an object containing the 
 * user's information to add the user to the database
 */
const register = (userName, email, password, passwordConf, firstName, lastName, storeName) => {
  // console.log("auth.service, line: 7", {userName, email, password, passwordConf, firstName, lastName, storeName});
  return axios.post(API_URL + "/user", {
    email,
    password,
    passwordConf,
    userName,
    firstName,
    lastName,
    storeName
  });
};

/**
 * @description - login sends POST request to the server with an object 
 * containing the user's login credentials and store the response data in the local storage
 */
const login = async (email, password) => {
  const response = await axios
    .post(API_URL + "/sessions", {
      email,
      password,
    });
  if (response.headers['authorization']) {
    localStorage.setItem("user", JSON.stringify(response.data));
    localStorage.setItem("token", response.headers['authorization']);
  }
  // let user = localStorage.getItem("user");
  // let token = localStorage.getItem("token");
  // console.log("auth.service, line:28", user);
  // console.log("auth.service, line:29", token);
  let result = {
    user: response.data,
    token: response.headers['authorization']
  };
  return result;
};

/**
 * @description - logout sends DELETE request to the server with the user's 
 * token and remove it from the local storage
 */
const logout = async () => {
    console.log("logout");
    localStorage.removeItem("user");
    localStorage.removeItem("token");
    const res = await axios.delete(API_URL + "/sessions/mine", { headers: authHeader() });
    if(res.status === 200) {
        console.log("success")
    }
    
};

export default {
  register,
  login,
  logout
};