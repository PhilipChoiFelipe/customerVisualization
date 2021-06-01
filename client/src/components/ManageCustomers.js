import React, { useState, useEffect, useRef } from "react";
import { useDispatch, useSelector } from "react-redux";

import Form from "react-validation/build/form";
import Input from "react-validation/build/input";
import CheckButton from "react-validation/build/button";

import { getAllItems, getSpecItem } from "../actions/customer";

import CustomerService from "../services/user.service.customer";


const required = (value) => {
  if (!value) {
    return (
      <div className="alert alert-danger" role="alert">
        This field is required!
      </div>
    );
  }
};

const ManageCustomers = () => {
  
  const { user: currentUser } = useSelector((state) => state.auth);
  const { customers } = useSelector((state) => state.customers);

  const form = useRef();
  const checkBtn = useRef();

  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [ethnicity, setEthnicity] = useState("");
  const [gender, setGender] = useState("");
  const [birthday, setBirthday] = useState("");
  const [postalCode, setPostalCode] = useState(0);
  const [lastVisited, setLastVisited] = useState("");
  const [disChannel, setDisChannel] = useState("");
  const [favItem, setFavItem] = useState("");

  const [successful, setSuccessful] = useState(false);

  const { message } = useSelector(state => state.message);
  const dispatch = useDispatch();

  if (customers && customers.length === 0) {
    dispatch(getAllItems(currentUser.id));
  }
  console.log("manageCustomers: 49", customers)

  const onChangeFirstName = (e) => {
    const firstName = e.target.value;
    setFirstName(firstName);
  };
  const onChangeLastName = (e) => {
    const lastName = e.target.value;
    setLastName(lastName);
  };
  const onChangeEthnicity = (e) => {
    const ethnicity = e.target.value;
    setEthnicity(ethnicity);
  };
  const onChangeGender = (e) => {
    const gender = e.target.value;
    setGender(gender);
  };
  const onChangeBirthday = (e) => {
    const birthday = e.target.value;
    setBirthday(birthday);
  };
  const onChangePostalCode = (e) => {
    const postalCode = e.target.value;
    setPostalCode(postalCode);
  };
  const onChangeLastVisted = (e) => {
    const lastVisited = e.target.value;
    setLastVisited(lastVisited);
  };
  const onChangeDisChannel = (e) => {
    const disChannel = e.target.value;
    setDisChannel(disChannel);
  };
  const onChangeFavItem = (e) => {
    const favItem = e.target.value;
    setFavItem(favItem);
  };

  const handleAddItem = (e) => {
    e.preventDefault();

    setSuccessful(false);

    // form.current.validateAll();

    if (checkBtn.current.context._errors.length === 0) {
      let customerObj = {
        userId: currentUser.id,
        firstName,
        lastName,
        ethnicity,
        gender,
        birthday,
        postalCode,
        lastVisited,
        disChannel,
        favItem
      }
      CustomerService.createCustomer(currentUser.id, customerObj).then(
        (response) => {
          console.log(response.data)
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
  };

  return (
    <div>
      <div className="col-md-12">
        <div className="card card-container">

          <Form onSubmit={handleAddItem} ref={form}>
            {!successful && (
              <div>
                <div className="form-group">
                  <label htmlFor="firstName">First Name</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="firstName"
                    value={firstName}
                    onChange={onChangeFirstName}
                    // validations={required}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="lirstName">Last Name</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="lastName"
                    value={lastName}
                    onChange={onChangeLastName}
                    // validations={required}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="price">Price</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="price"
                    value={price}
                    onChange={onChangePrice}
                    // validations={required}
                  />
                </div>

                <div className="form-group">
                  <button className="btn btn-primary btn-block">Add Item</button>
                </div>
              </div>
            )}

            {message && (
              <div className="form-group">
                <div className={successful ? "alert alert-success" : "alert alert-danger"} role="alert">
                  {message}
                </div>
              </div>
            )}
            <CheckButton style={{ display: "none" }} ref={checkBtn} />
          </Form>
        </div>
      </div>

      <div className="container">
        <header className="jumbotron">
          {/* For testing */}
          {/* <div>{JSON.stringify(items)}</div> */}

          {/* TODO: maps individual item, add storeId */}
          {items.length > 0 ? (
            items.map(item => {
              return (
                <div>
                  <h3>{item.itemName}</h3>
                  <p>{item.price}</p>
                </div>
              )
            })
          ) :
            (<h3>No item added yet</h3>)
          }
        </header>
      </div>

    </div>
  );
};

export default ManageCustomers;