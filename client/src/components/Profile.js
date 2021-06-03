import React, { useCallback, useState } from "react";
import { Redirect } from 'react-router-dom';
import { useDispatch, useSelector } from "react-redux";
import { getAllCustomers } from "../actions/customer";
import CustomerService from "../services/user.service.customer";
import { getAllItems } from "../actions/item";
import _ from "lodash";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSeedling, faUserFriends, faThumbsUp } from '@fortawesome/free-solid-svg-icons'

const Profile = () => {
  const { user: currentUser, token: authToken } = useSelector((state) => state.auth);
  const { customers } = useSelector((state) => state.customer);
  // const { items } = useSelector((state) => state.item)

  const dispatch = useDispatch();

  const [recentWeekCus, setRecentWeekCus] = useState(null);
  const [recentMonthCus, setRecentMonthCus] = useState(null);
  
  if (!currentUser || !authToken) {
    return <Redirect to="/login" />;
  }

  const handleRecentCustomers = (days, setState) => {
    let today = new Date();
    let lastDate = new Date(today.setDate(today.getDate() - days))
    lastDate = lastDate.getFullYear() + "-" + (lastDate.getMonth() + 1) + "-" + lastDate.getDate();
    console.log(lastDate)

    CustomerService.getAllCustomers(currentUser.id, {sort: "last_visited", reverse: true, before: lastDate}).then(
      (response) => {
        setState(response);
      },
      (error) => {
        const err =
          (error.response &&
            error.response.data &&
            error.response.data.message) ||
          error.message ||
          error.toString();
          console.log(err)
      }
    );
  }

  if (customers && customers.length === 0) {
    dispatch(getAllCustomers(currentUser.id))
    handleRecentCustomers(7, setRecentWeekCus);
    handleRecentCustomers(31, setRecentMonthCus);
  }
  
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
                {recentMonthCus && recentMonthCus.length > 0 ? (recentMonthCus.length):(0)}
              </b></h1>
            <p class="lead">new visiters this month</p>
          </div>
          <div class="col-md-4">
            <div class="FAIcon"><FontAwesomeIcon icon={faUserFriends}  size="5x" color="#f7efc3"/></div>
            
              <h1 class="display-4"><b>
                {recentWeekCus && recentWeekCus.length > 0 ? (recentWeekCus.length):(0)}
              </b></h1>
            <p class="lead">customers have visited total</p>
          </div>
          <div class="col-md-4">
            <div class="FAIcon"><FontAwesomeIcon icon={faThumbsUp}  size="5x" color="#c3e9f7"/></div>

            {/* TODO: get most common fav item id in customers, SQL query?*/}
            <h1 class="display-4"><b>Hamburger</b></h1>
            <p class="lead">is currently everyone's favorite</p>
          </div>
        </div>

      </div>
    </div>
  );
};

export default Profile;