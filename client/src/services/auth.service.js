import axios from "axios";
import authHeader from "./auth-header";
const API_URL = "http://localhost:80/v1";


const register = (userName, email, password, passwordConf, firstName, lastName) => {
  return axios.post(API_URL + "/user", {
    userName,
    email,
    password,
    passwordConf,
    firstName,
    lastName
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
    const res = await axios.delete(API_URL + "/sessions/mine", { headers: authHeader() });
    if(res.status === 200) {
        console.log("success")
    }
    localStorage.removeItem("user");
    localStorage.removeItem("token");
};

export default {
  register,
  login,
  logout
};