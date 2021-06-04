import React, { useCallback } from "react";
import { Redirect } from 'react-router-dom';
import { useDispatch, useSelector } from "react-redux";
import { getAllCustomers } from "../actions/customer";
import { getSpecItem } from "../actions/item";
import _ from "lodash";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSeedling, faChartLine, faUserFriends, faThumbsUp } from '@fortawesome/free-solid-svg-icons'


/**
 *@description Component Profile returns customer data in different ways  
*/
const Profile = () => {
  const { user: currentUser, token: authToken } = useSelector((state) => state.auth);
  const { customers } = useSelector((state) => state.customer);
  const { specItem } = useSelector((state) => state.item);

  const dispatch = useDispatch();

  const handleRecentCustomers = useCallback((days) => {
    let today = new Date();
    let lastDate = new Date(today.setDate(today.getDate() - days))
    let recentCustomers = customers.filter(customer => {
      let tempDate = new Date(customer['lastVisited'])
      return tempDate > lastDate
    })
    console.log("recent customers", recentCustomers)
    return recentCustomers;
  }, [customers])

  const handleFavoriteItem = useCallback(() => {
    let cusToItem = _.countBy(customers, 'favItem');
    let max = 0;
    let favItem = null;
    for (let key in cusToItem) {
      if (cusToItem[key] > max) {
        max = cusToItem[key]
        favItem = key
      }
    }
    console.log(favItem);
    dispatch(getSpecItem(currentUser.id, favItem))
  }, [customers, currentUser, dispatch])

  if (!currentUser || !authToken) {
    return <Redirect to="/login" />;
  }

  if (customers && customers.length === 0) {
    dispatch(getAllCustomers(currentUser.id))
  }

  if ((customers && customers.length > 0) && !specItem) {
    console.log("ITEM SPEC FUNCTION")
    handleFavoriteItem();
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
          <div class="col-md-3">
            <div class="FAIcon"><FontAwesomeIcon icon={faSeedling} size="5x" color="#d9f7c3" /></div>
            <h1 class="display-4"><b>
              {customers && customers.length > 0 && handleRecentCustomers(7).length > 0 ? (handleRecentCustomers(7).length) : (0)}
            </b></h1>
            <h6 class="lead">new visiters this week</h6>
          </div>
          <div class="col-md-3">
            <div class="FAIcon"><FontAwesomeIcon icon={faChartLine} size="5x" color="#faeed2" /></div>

            <h1 class="display-4"><b>
              {customers && customers.length > 0 && handleRecentCustomers(31).length > 0 ? (handleRecentCustomers(31).length) : (0)}
            </b></h1>
            <p class="lead">new visiters this month</p>
          </div>
          <div class="col-md-3">
            <div class="FAIcon"><FontAwesomeIcon icon={faUserFriends} size="5x" color="#c3e9f7" /></div>

            <h1 class="display-4"><b>
              {customers && customers.length > 0 ? (customers.length) : (0)}
            </b></h1>
            <p class="lead">customers have visited total</p>
          </div>
          <div class="col-md-3">
            <div class="FAIcon"><FontAwesomeIcon icon={faThumbsUp} size="5x" color="#ffdeea" /></div>

            <h1 class="display-4">
              <b>
                {specItem ? specItem['itemName'] : "no favorite item"}
              </b></h1>
            <p class="lead">is currently everyone's favorite</p>
          </div>
        </div>

      </div>
    </div>
  );
};

export default Profile;