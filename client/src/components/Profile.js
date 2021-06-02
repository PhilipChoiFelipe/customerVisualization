import React, { useCallback, useState } from "react";
import { Redirect } from 'react-router-dom';
import { useDispatch, useSelector } from "react-redux";
import { getAllCustomers } from "../actions/customer";
import { getAllItems } from "../actions/item";
import _ from "lodash";

const Profile = () => {
  const { user: currentUser, token: authToken } = useSelector((state) => state.auth);
  const { customers } = useSelector((state) => state.customer);
  const { items } = useSelector((state) => state.item)

  const dispatch = useDispatch();
  // const maxItem = useCallback(() => {
  //   let item_temp = _.countBy(customers, 'favItem');
  //   console.log(customers)
  //   console.log(item_temp)
  //   let maxItem = {
  //     item: null,
  //     count: 0
  //   };
  //   for (let key in item_temp) {
  //     if (item_temp[key] > maxItem['count']) {
  //       maxItem['count'] = item_temp[key];
  //       maxItem['item'] = key;
  //     }
  //   }
  //   return maxItem;
  // }, [customers]);
  
  if (!items) {
    return <Redirect to="/login" />;
  }

  if (items && items.length === 0) {
    dispatch(getAllItems(currentUser.id));
  }

  if (customers && customers.length === 0) {
    dispatch(getAllCustomers(currentUser.id))
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