import React, { useState } from "react";
import { Redirect } from 'react-router-dom';
import { useDispatch, useSelector } from "react-redux";
import { getAllCustomers } from "../actions/customer";

const Profile = () => {
  const { user: currentUser, token: authToken } = useSelector((state) => state.auth);
  const { customers } = useSelector((state) => state.customer);
  const dispatch = useDispatch();
  if (!currentUser) {
    return <Redirect to="/login" />;
  }
  if (customers && customers.length === 0) {
    dispatch(getAllCustomers(currentUser.id))
  } else {
  }

  return (
    <div className="container">
      <header className="jumbotron">
        <h1>
          Hello {currentUser.userName}
        </h1>
        <h3>
          Here's an overview of your business
        </h3>
      </header>
      {customers &&
        (<ul>{customers.map(customer => {
          return <li>{customer.firstName}</li>
        })}</ul>)}
      <p>
        <strong>Token:</strong> {authToken.substring(0, 20)} ...{" "}
        {authToken.substr(authToken.length - 20)}
      </p>
      {/* <strong>Authorities:</strong> */}
      {/* <ul>
        {currentUser.roles &&
          currentUser.roles.map((role, index) => <li key={index}>{role}</li>)}
      </ul> */}
    </div>
  );
};

export default Profile;