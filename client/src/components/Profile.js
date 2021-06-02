import React, { useCallback, useState } from "react";
import { Redirect } from 'react-router-dom';
import { useDispatch, useSelector } from "react-redux";
import { getAllCustomers } from "../actions/customer";
import { getAllItems } from "../actions/item";
import _ from "lodash";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSeedling, faUserFriends, faThumbsUp } from '@fortawesome/free-solid-svg-icons'
// import ReactCSSTransitionGroup from 'react-transition-group'; // ES6
// var ReactCSSTransitionGroup = require('react-transition-group'); // ES5 with npm

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
  
  // TODO: deleted bc it causes infinite loop

  // if (items && items.length === 0) {
  //   dispatch(getAllItems(currentUser.id));
  // }

  // if (customers && customers.length === 0) {
  //   dispatch(getAllCustomers(currentUser.id))
  // }
  
  return (
    <div className="container">
      <header className="jumbotron">
        <h1 class="display-4"><strong>Hello {currentUser.userName}!</strong></h1>
        <p class="lead">
          Here's an overview of your business
        </p>
        <hr class="my-4"></hr>
      </header>

      <div id="services" class="container-fluid text-center">
        <br></br>
        <div class="row slideanim">
          <div class="col-md-4">
            <div class="FAIcon"><FontAwesomeIcon icon={faSeedling}  size="5x" color="#d9f7c3"/></div>
              <h1 class="display-4"><b>
                {/* 128 */}
                {customers && customers.length > 0 ? (customers.length):(0)}
              </b></h1>
            <h6  class="lead">new visiters this week</h6>
          </div>
          <div class="col-md-4">
            <div class="FAIcon"><FontAwesomeIcon icon={faUserFriends}  size="5x" color="#f7efc3"/></div>
            
              <h1 class="display-4"><b>
                {customers && customers.length > 0 ? (customers.length):(0)}
              </b></h1>
            <p class="lead">customers have visited</p>
          </div>
          <div class="col-md-4">
            <div class="FAIcon"><FontAwesomeIcon icon={faThumbsUp}  size="5x" color="#c3e9f7"/></div>

            {/* TODO: get most common fav item id in customers, SQL query?*/}
            <h1 class="display-4"><b>Hamburger</b></h1>
            <p class="lead">is currently everyone's favorite</p>
          </div>
        </div>

      </div>
      {/* </div> */}

      {/* {customers &&
        (<ul>{customers.map(customer => {
          return <li>{customer.firstName}</li>
        })}</ul>)} */}

      
      {/* <p>
        <strong>Token:</strong> {authToken.substring(0, 20)} ...{" "}
        {authToken.substr(authToken.length - 20)}
      </p> */}
      {/* <strong>Authorities:</strong> */}
      {/* <ul>
        {currentUser.roles &&
          currentUser.roles.map((role, index) => <li key={index}>{role}</li>)}
      </ul> */}
    </div>
  );
};

export default Profile;