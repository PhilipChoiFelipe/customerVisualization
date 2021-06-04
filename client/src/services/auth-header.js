
/**
 * @description - authHeader add authorization token of the currently 
 * logged in user from Local Storage to HTTP header and returns the object.
 */
export default function authHeader() {
    const user = JSON.parse(localStorage.getItem('user'));
    const token = localStorage.getItem('token');
    console.log("auth-header, line:4", token);
    if (user && token) {
      return { Authorization: 'Bearer ' + token };
    } else {
      return {};
    }
  }