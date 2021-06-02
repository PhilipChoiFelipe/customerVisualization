import React, { useState, useEffect, useRef } from "react";
import { useDispatch, useSelector } from "react-redux";

import Form from "react-validation/build/form";
import Input from "react-validation/build/input";
import CheckButton from "react-validation/build/button";

import { getAllItems, getSpecItem, updateSpecItem } from "../actions/item";
import ItemService from "../services/user.service.item";

//Style
import { Modal, Button, Table} from 'react-bootstrap';

const required = (value) => {
  if (!value) {
    return (
      <div className="alert alert-danger" role="alert">
        This field is required!
      </div>
    );
  }
};

const ManageItems = () => {
  
  
  const { user: currentUser } = useSelector((state) => state.auth);
  const { items } = useSelector((state) => state.item);

  const form = useRef();
  const checkBtn = useRef();

  const [itemName, setItemName] = useState("");
  const [price, setPrice] = useState(0);

  const [itemId, setItemId] = useState(0);

  const [show, setShow] = useState(false);


  const [successful, setSuccessful] = useState(false);

  const { message } = useSelector(state => state.message);
  const dispatch = useDispatch();


  if (items && items.length === 0) {
    dispatch(getAllItems(currentUser.id));
  }

  // useEffect(() => {
  //   dispatch(getAllItems(currentUser.id));
  // }, [successful, dispatch, currentUser])
  // // console.log("manageItems: 41",items)

  const onChangeItemName = (e) => {
    const itemName = e.target.value;
    setItemName(itemName);
  };

  const onChangePrice = (e) => {
    const price = e.target.value;
    setPrice(price);
  };

  // const setItemId = (e) => {
  //   const itemId = e.target.value;
  //   setItemId(itemId);
  // };

  const handleAddItem = (e) => {
    e.preventDefault();

    setSuccessful(false);
    // form.current.validateAll();
    if (checkBtn.current.context._errors.length === 0) {
      let itemObj = {
        userId: currentUser.id,
        itemName: itemName,
        price: parseInt(price)
      }
      console.log("manageItems: 66", itemObj)
      ItemService.createItem(currentUser.id, itemObj).then(
        (response) => {
          console.log(response.data)
          dispatch(getAllItems(currentUser.id))
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

  // TODO: change to update
  const handleUpdateItem = (e) => {
    e.preventDefault();

    setSuccessful(false);
    // form.current.validateAll();
    if (checkBtn.current.context._errors.length === 0) {
      let itemUpdateObj = {
        itemName: itemName,
        price: parseInt(price)
      }
      console.log("manageItems: 108", itemUpdateObj)
      ItemService.updateSpecItem(currentUser.id, itemId ,itemUpdateObj).then(
        (response) => {
          console.log(response.data)
          dispatch(getAllItems(currentUser.id))
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

  const createRows = (items) => {
    return items.map( item => {
      return (
          <tr key={item.id} onClick={()=>setItemId(item.id)}>
            <th scope="row">{item.id}</th>
            <td>{item.itemName}</td>
            <td>{item.price}</td>
          </tr>
      );
    })
    
  }

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  return (
    <div>
      <div className="col-md-12">
        <div className="card card-container">

          <Form onSubmit={handleAddItem} ref={form}>
            {!successful && (
              <div>
                <div className="form-group">
                  <label htmlFor="itemName">Item Name</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="itemName"
                    value={itemName}
                    onChange={onChangeItemName}
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
                  <button className="btn btn-dark btn-block">Add Item</button>
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
        <Modal show={show} onHide={handleClose}>
          <Modal.Header closeButton>
            <Modal.Title>Modal heading</Modal.Title>
          </Modal.Header>
          <Modal.Body>Woohoo, you're reading this text in a modal!</Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={handleClose}>
              Close
            </Button>
            <Button variant="primary" onClick={handleClose}>
              Save Changes
            </Button>
          </Modal.Footer>
        </Modal>
      <button onClick={handleShow}>Edit</button>
      
      <button>Delete</button>
      <table class="table table-hover">
        <thead>
          <tr>
            <th scope="col">Item ID</th>
            <th scope="col">Item Name</th>
            <th scope="col">Price</th>
          </tr>
        </thead>
        <tbody>
          {items && items.length > 0 ?
            createRows(items)
            
          :(
            <p>No item added yet</p>
          )
          }
          
        </tbody>
      </table>
      </div>
    </div>
  );
};

export default ManageItems;

                          {/* <div className="form-group">
                            <label htmlFor="itemName">Item Name</label>
                            <Input
                              type="text"
                              className="form-control"
                              name="itemName"
                              value={itemName}
                              onChange={onChangeItemName}
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
                          </div> */}


{/* {message && (
                        <div className="form-group">
                          <div className={successful ? "alert alert-success" : "alert alert-danger"} role="alert">
                            {message}
                          </div>
                        </div>
                      )}
<CheckButton style={{ display: "none" }} ref={checkBtn} /> */}