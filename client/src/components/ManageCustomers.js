import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import Form from "react-validation/build/form";
import Input from "react-validation/build/input";

import { getAllCustomers } from "../actions/customer";

import CustomerService from "../services/user.service.customer";

import { Modal, Button} from 'react-bootstrap';


/**
 *@description Component ManageCustomers returns table of customers that user can add, update, and delete
*/
const ManageCustomers = () => {
  
  const { user: currentUser } = useSelector((state) => state.auth);
  const { customers } = useSelector((state) => state.customer);

  const form = useRef();

  //Inserting new Item
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

  //Sorting by column 
  const [reverse, setReverse] = useState(false);
  const [currentSort, setCurrentSort] = useState(null);

  //checks
  const [successful, setSuccessful] = useState(false);
  const { message } = useSelector(state => state.message);

  //Modal
  const [show, setShow] = useState(false);
  const [modalTitle, setModalTitle] = useState(null);

  const dispatch = useDispatch();

  if (customers && customers.length === 0) {
    dispatch(getAllCustomers(currentUser.id));
  }

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

  };

  const handleUpdateCustomer = (e) => {
    setShow(false);
    setSuccessful(false);
      let cusUpdateObj = {
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

  const handleSortCustomers = (query) => {
    //if clicked column second time, switch reverse
    if (currentSort === query['sort']) {
      query['reverse'] = !reverse;
      setReverse(!reverse);
    } else {
    //if clicked column first time, reverse = false
      setCurrentSort(query['sort']);
      setReverse(false);
      query['reverse'] = false;
    }
    dispatch(getAllCustomers(currentUser.id, query));
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
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="ethnicity">Ethnicity</label>
                  <select class="form-select" aria-label="Default select example" onChange={onChangeEthnicity}>
                    <option selected>Ethnicity</option>
                    <option value="Black">Black</option>
                    <option value="White">White</option>
                    <option value="Native American">Native American</option>
                    <option value="Asian">Asian</option>
                    <option value="Latin">Latin</option>
                    <option value="Pacific Islander">Pacific Islander</option>
                  </select>
                </div>

                <div className="form-group">
                  <label htmlFor="gender">Gender</label>
                  <select class="form-select" aria-label="Default select example" onChange={onChangeGender}>
                    <option selected>Gender</option>
                    <option value="Male">Male</option>
                    <option value="Female">Female</option>
                    <option value="Others">Don't want to answer</option>
                  </select>
                </div>

                <div className="form-group">
                  <label htmlFor="birthday">Birthday</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="birthday"
                    value={birthday}
                    onChange={onChangeBirthday}
                    placeholder={"YYYY-MM-DD"}
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
                    placeholder={"YYYY-MM-DD"}
                  />
                </div>

                <div className="form-group">
                  <label htmlFor="disChannel">Found Us Through</label>
                  <select class="form-select" aria-label="Default select example" onChange={onChangeDisChannel}>
                    <option selected>Channel</option>
                    <option value="Instagram">Instagram</option>
                    <option value="Facebook">Facebook</option>
                    <option value="Google">Google</option>
                    <option value="News">News</option>
                    <option value="Poster">Poster</option>
                    <option value="Youtube">Youtube</option>
                    <option value="Friends">Friends</option>
                  </select>
                </div>

                <div className="form-group">
                  <label htmlFor="favItem">Favorite Item ID</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="favItem"
                    value={favItem}
                    onChange={onChangeFavItem}
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
        <thead >
          <tr>
            <th scope="col"><Button variant="light" size="sm" onClick={() => {handleSortCustomers({sort: "id", reverse: reverse})}}>Customer ID</Button></th>
            <th scope="col"><Button variant="light" size="sm" onClick={() => {handleSortCustomers({sort: "first_name", reverse: reverse})}}>First Name</Button></th>
            <th scope="col"><Button variant="light" size="sm" onClick={() => {handleSortCustomers({sort: "last_name", reverse: reverse})}}>Last Name</Button></th>
            <th scope="col"><Button variant="light" size="sm" onClick={() => {handleSortCustomers({sort: "ethnicity", reverse: reverse})}}>Ethnicity</Button></th>
            <th scope="col"><Button variant="light" size="sm" onClick={() => {handleSortCustomers({sort: "gender", reverse: reverse})}}>Gender</Button></th>
            <th scope="col"><Button variant="light" size="sm" onClick={() => {handleSortCustomers({sort: "birthday", reverse: reverse})}}>Birthday</Button></th>
            <th scope="col"><Button variant="light" size="sm" >Postal Code</Button></th>
            <th scope="col"><Button variant="light" size="sm" onClick={() => {handleSortCustomers({sort: "last_visited", reverse: reverse})}}>Last Visited</Button></th>
            <th scope="col"><Button variant="light" size="sm" onClick={() => {handleSortCustomers({sort: "dis_channel", reverse: reverse})}}>Found Us Through</Button></th>
            <th scope="col"><Button variant="light" size="sm" onClick={() => {handleSortCustomers({sort: "fav_item", reverse: reverse})}}>Favorite Item</Button></th>
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
            <h6>No customer added yet! Try inserting some customers and remember to fill out every field!</h6>
          </div>
        )}
      
      </div>
    </div>
  );
};

export default ManageCustomers;