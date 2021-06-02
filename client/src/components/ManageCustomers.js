import React, { useState, useEffect, useRef } from "react";
import { useDispatch, useSelector } from "react-redux";

import Form from "react-validation/build/form";
import Input from "react-validation/build/input";
import CheckButton from "react-validation/build/button";

import { getAllCustomers, getSpecCustomers } from "../actions/customer";

import CustomerService from "../services/user.service.customer";

import { Modal, Button} from 'react-bootstrap';

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
  const { customers } = useSelector((state) => state.customer);

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
  const [favItem, setFavItem] = useState(0);
  const [customerId, setCustomerId] = useState(0);

  const [successful, setSuccessful] = useState(false);

  //Modal
  const [show, setShow] = useState(false);
  const [modalTitle, setModalTitle] = useState(null);

  const { message } = useSelector(state => state.message);
  const dispatch = useDispatch();

  if (customers && customers.length === 0) {
    dispatch(getAllCustomers(currentUser.id));
  }
  // console.log("manageCustomers: 49", customers)

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
  const onChangeLastVisited = (e) => {
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

  const handleAddCustomer= (e) => {
    e.preventDefault();

    setSuccessful(false);
    setShow(false);
    // form.current.validateAll();

    // if (checkBtn.current.context._errors.length === 0) {
      let customerObj = {
        userId: currentUser.id,
        firstName,
        lastName,
        ethnicity,
        gender,
        birthday,
        postalCode: parseInt(postalCode),
        lastVisited,
        disChannel,
        favItem: parseInt(favItem)
      }
      CustomerService.createCustomer(currentUser.id, customerObj).then(
        (response) => {
          console.log(response.data)
          dispatch(getAllCustomers(currentUser.id))
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

    // }
  };

  const handleUpdateCustomer = (e) => {
    setShow(false);
    setSuccessful(false);
    // form.current.validateAll();
      let cusUpdateObj = {
        firstName,
        lastName
      }
      console.log("manageItems: 108", cusUpdateObj)
      CustomerService.updateSpecCustomer(currentUser.id, customerId ,cusUpdateObj).then(
        (response) => {
          console.log(response.data)
          dispatch(getAllCustomers(currentUser.id))
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
  };

  const handleDeleteCustomer = () => {
    CustomerService.deleteSpecCustomer(currentUser.id, customerId).then(
      (response) => {
        console.log(response.data)
        dispatch(getAllCustomers(currentUser.id))
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

  const createRows = (customers) => {
    return customers.map( customer => {
      return (
          <tr className={customerId && customer.id === customerId?"table-active":""} key={customer.id} onClick={()=>setCustomerId(customer.id)}>
            <th scope="row">{customer.id}</th>
            <td>{customer.firstName}</td>
            <td>{customer.lastName}</td>
            <td>{customer.ethnicity}</td>
            <td>{customer.gender}</td>
            <td>{customer.birthday}</td>
            <td>{customer.postalCode}</td>
            <td>{customer.lastVisited}</td>
            <td>{customer.disChannel}</td>
            <td>{customer.favItem}</td>
          </tr>
      );
    })
    
  }

  const handleClose = () => setShow(false);
  const handleShow = (title) => {
    setShow(true);
    setModalTitle(title);
  }

   // Ethnicity  string `json:"ethnicity"`
  // Gender     string `json:"gender"`
  // Birthday   string `json:"birthday"` // TOASK: better datatype?
  // PostalCode int64  `json:"postalCode"`
  // LastVisted string `json:"lastVisited"`
  // DisChannel string `json:"disChannel"`
  // FavItem 

  return (
    <div>
      <div className="container">  
        <Modal show={show} onHide={handleClose}>
          <Modal.Header >
            <Modal.Title>{modalTitle}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
          <Form onSubmit={handleUpdateCustomer} ref={form}>
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
                  <label htmlFor="lastName">Last Name</label>
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
                  <label htmlFor="ethnicity">Ethnicity</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="ethnicity"
                    value={ethnicity}
                    onChange={onChangeEthnicity}
                    // validations={required}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="gender">Gender</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="gender"
                    value={gender}
                    onChange={onChangeGender}
                    // validations={required}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="birthday">Birthday</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="birthday"
                    value={birthday}
                    onChange={onChangeBirthday}
                    // validations={required}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="postalCode">Postal Code</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="postalCode"
                    value={postalCode}
                    onChange={onChangePostalCode}
                    // validations={required}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="lastVisited">Last Visited Date</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="lastVisited"
                    value={lastVisited}
                    onChange={onChangeLastVisited}
                    // validations={required}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="disChannel">Found Us Through</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="disChannel"
                    value={disChannel}
                    onChange={onChangeDisChannel}
                    // validations={required}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="favItem">Favorite Item ID</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="favItem"
                    value={favItem}
                    onChange={onChangeFavItem}
                    // validations={required}
                  />
                </div>

              </div>
            </Form>
          </Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={handleClose}>
              Close
            </Button>
            {modalTitle === "Add Customer" ? <Button variant="primary" onClick={handleAddCustomer}>
              Add Customer
            </Button> : <Button variant="primary" onClick={handleUpdateCustomer}>
              Update Customer
            </Button>}
          </Modal.Footer>
        </Modal>
      <Button onClick={() => handleShow("Add Customer")}>Insert</Button>
      {customerId ? <><Button onClick={() => handleShow("Update Customer")}>Edit</Button>
      <Button onClick={handleDeleteCustomer}>Delete</Button> </>: <></>}
      <table class="table table-hover">
        <thead>
          <tr>
            <th scope="col">Customer ID</th>
            <th scope="col">First Name</th>
            <th scope="col">Last Name</th>
            <th scope="col">Ethnicity</th>
            <th scope="col">Gender</th>
            <th scope="col">Birthday</th>
            <th scope="col">Postal Code</th>
            <th scope="col">Last Visited</th>
            <th scope="col">Found Us Through</th>
            <th scope="col">Favorite Item ID</th>
          </tr>
        </thead>
        <tbody>
          {customers && customers.length > 0 &&
            createRows(customers)
          }
        </tbody>
      </table>
        {!customers && (
          <div class="alert alert-primary">
            <h6>No customer added yet! Try inserting some customers.</h6>
          </div>
        )}
      
      </div>
    </div>
  );
};

export default ManageCustomers;

// FirstName  string `json:"firstName"`
//     LastName   string `json:"lastName"`
//     Ethnicity  string `json:"ethnicity"`
//     Gender     string `json:"gender"`
//     Birthday   string `json:"birthday"` // TOASK: better datatype?
//     PostalCode int64  `json:"postalCode"`
//     LastVisted string `json:"lastVisited"`
//     DisChannel string `json:"disChannel"`
//     FavItem    int64  `json:"favItem"`

// <div>
//       <div className="col-md-12">
//         <div className="card card-container">

//           <Form onSubmit={handleAddCustomer} ref={form}>
//             {!successful && (
//               <div>
//                 <div className="form-group">
//                   <label htmlFor="firstName">First Name</label>
//                   <Input
//                     type="text"
//                     className="form-control"
//                     name="firstName"
//                     value={firstName}
//                     onChange={onChangeFirstName}
//                     // validations={required}
//                   />
//                 </div>

//                 <div className="form-group">
//                   <label htmlFor="lastName">Last Name</label>
//                   <Input
//                     type="text"
//                     className="form-control"
//                     name="lastName"
//                     value={lastName}
//                     onChange={onChangeLastName}
//                     // validations={required}
//                   />
//                 </div>

//                 <div className="form-group">
//                   <label htmlFor="ethnicity">Ethnicity</label>
//                   <Input
//                     type="text"
//                     className="form-control"
//                     name="ethnicity"
//                     value={ethnicity}
//                     onChange={onChangeEthnicity}
//                     // validations={required}
//                   />
//                 </div>

//                 <div className="form-group">
//                   <label htmlFor="gender">Gender</label>
//                   <Input
//                     type="text"
//                     className="form-control"
//                     name="gender"
//                     value={gender}
//                     onChange={onChangeGender}
//                     // validations={required}
//                   />
//                 </div>

//                 <div className="form-group">
//                   <label htmlFor="birthday">Birthday</label>
//                   <Input
//                     type="text"
//                     className="form-control"
//                     name="birthday"
//                     value={birthday}
//                     onChange={onChangeBirthday}
//                     // validations={required}
//                   />
//                 </div>

//                 <div className="form-group">
//                   <label htmlFor="postalCode">Postal Code</label>
//                   <Input
//                     type="text"
//                     className="form-control"
//                     name="postalCode"
//                     value={postalCode}
//                     onChange={onChangePostalCode}
//                     // validations={required}
//                   />
//                 </div>

//                 <div className="form-group">
//                   <label htmlFor="lastVisited">Last Visited Date</label>
//                   <Input
//                     type="text"
//                     className="form-control"
//                     name="lastVisited"
//                     value={lastVisited}
//                     onChange={onChangeLastVisited}
//                     // validations={required}
//                   />
//                 </div>

//                 <div className="form-group">
//                   <label htmlFor="disChannel">Found Us Through</label>
//                   <Input
//                     type="text"
//                     className="form-control"
//                     name="disChannel"
//                     value={disChannel}
//                     onChange={onChangeDisChannel}
//                     // validations={required}
//                   />
//                 </div>

//                 <div className="form-group">
//                   <label htmlFor="favItem">Favorite Item</label>
//                   <Input
//                     type="text"
//                     className="form-control"
//                     name="favItem"
//                     value={favItem}
//                     onChange={onChangeFavItem}
//                     // validations={required}
//                   />
//                 </div>

//                 <div className="form-group">
//                   <button className="btn btn-primary btn-block">Add Customer</button>
//                 </div>
//               </div>
//             )}

//             {message && (
//               <div className="form-group">
//                 <div className={successful ? "alert alert-success" : "alert alert-danger"} role="alert">
//                   {message}
//                 </div>
//               </div>
//             )}
//             <CheckButton style={{ display: "none" }} ref={checkBtn} />
//           </Form>
//         </div>
//       </div>

//       <div className="container">
//         <header className="jumbotron">
//           {customers && customers.length > 0 ? (
//             customers.map(customer => {
//               return (
//                 <div>
//                   <h3>Name: {customer.firstName} {customer.lastName}</h3>
//                   <p>Ethnicity: {customer.ethnicity}</p>
//                   <p>Gender: {customer.gender}</p>
//                   <p>Birthday: {customer.birthday}</p>
//                   <p>Postal Code: {customer.postalCode}</p>
//                   <p>Last Visited Date: {customer.lastVisited}</p>
//                   <p>Found Us Through: {customer.disChannel}</p>
//                   <p>Favorite Item: {customer.favItem}</p>
//                 </div>
//               )
//             })
//           ) :
//             (<h3>No customer added yet</h3>)
//           }
//         </header>
//       </div>

//     </div>