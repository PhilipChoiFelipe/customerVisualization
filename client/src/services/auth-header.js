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