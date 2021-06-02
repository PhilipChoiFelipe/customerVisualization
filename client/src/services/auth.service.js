import axios from "axios";
import authHeader from "./auth-header";
const API_URL = "http://localhost:80/v1";


const register = (userName, email, password, passwordConf, firstName, lastName, storeName) => {
  console.log("auth.service, line: 7", {userName, email, password, passwordConf, firstName, lastName, storeName});
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