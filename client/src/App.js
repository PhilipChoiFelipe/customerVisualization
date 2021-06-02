import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Router, Switch, Route, Link } from "react-router-dom";

import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";

import Login from "./components/Login";
import Register from "./components/Register";
import Home from "./components/Home";
import Profile from "./components/Profile";
import ManageItems from "./components/manageItems";
import ManageCustomers from "./components/ManageCustomers";
import Visualization from "./components/Visualization";

// import BoardUser from "./components/BoardUser";
// import BoardModerator from "./components/BoardModerator";
// import BoardAdmin from "./components/BoardAdmin";

import { logout } from "./actions/auth";
import { clearMessage } from "./actions/message";

import { history } from "./helpers/history";

const App = () => {
  // const [showModeratorBoard, setShowModeratorBoard] = useState(false);
  // const [showAdminBoard, setShowAdminBoard] = useState(false);

  const { user: currentUser } = useSelector((state) => state.auth);
  const dispatch = useDispatch();

  useEffect(() => {
    history.listen((location) => {
      dispatch(clearMessage()); // clear message when changing location
    });
  }, [dispatch]);

  // useEffect(() => {
  //   if (currentUser) {
  //     setShowModeratorBoard(currentUser.roles.includes("ROLE_MODERATOR"));
  //     setShowAdminBoard(currentUser.roles.includes("ROLE_ADMIN"));
  //   }
  // }, [currentUser]);

  const logOut = () => {
    dispatch(logout());
  };

  return (
    <Router history={history}>
      <div>
        <nav className="navbar navbar-expand navbar-light bg-white shadow-sm p-3 mb-5 bg-white rounded">
          
          <Link to={"/"}  class="pull-left">
            <a class="navbar-brand" href="#">
              <img src="img/emlogo.png" height="60" alt="logo"></img>
            </a>
          </Link>
           
          {currentUser ? (
            <div className="navbar-nav ml-auto">
              <li className="nav-item">
                <Link to={"/profile"} className="nav-link">
                  <strong>{currentUser.userName}</strong>
                </Link>
              </li>
              <li className="nav-item">
                <Link to={"/manageItems"} className="nav-link">
                  Product / Service
                </Link>
              </li>
              <li className="nav-item">
                <Link to={"/manageCustomers"} className="nav-link">
                  Customers
                </Link>
              </li>
              <li className="nav-item">
                <Link to={"/visualization"} className="nav-link">
                  Visualization
                </Link>
              </li>
              
              <li className="nav-item">
                <a href="/login" className="nav-link" onClick={logOut}>
                  LogOut
                </a>
              </li>
            </div>
          ) : (
            <div className="nav navbar-nav ml-auto navbar-right">
              <li className="nav-item">
                <Link to={"/login"} className="nav-link">
                  Login
                </Link>
              </li>

              <li className="nav-item">
                <Link to={"/register"} className="nav-link">
                  Sign Up
                </Link>
              </li>
            </div>
          )}
        </nav>

        <div className="container mt-3">
          <Switch>
            <Route exact path={["/", "/home"]} component={Home} />
            <Route exact path="/login" component={Login} />
            <Route exact path="/register" component={Register} />
            <Route exact path="/profile" component={Profile} />
            <Route exact path="/manageItems" component={ManageItems} />
            <Route exact path="/manageCustomers" component={ManageCustomers} />
            <Route exact path="/visualization" component={Visualization} />
          </Switch>
        </div>
      </div>
    </Router>
  );
};

export default App;