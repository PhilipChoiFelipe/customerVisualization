import React, { useState, useRef } from "react";
import { useDispatch, useSelector } from "react-redux";

import Form from "react-validation/build/form";
import Input from "react-validation/build/input";

import { getAllItems } from "../actions/item";
import ItemService from "../services/user.service.item";

import { Modal, Button } from 'react-bootstrap';


/**
 *@description Component ManageItems returns table of items that user can add, update, and delete
*/
const ManageItems = () => {


  const { user: currentUser } = useSelector((state) => state.auth);
  const { items } = useSelector((state) => state.item);

  const form = useRef();

  //Inserting
  const [itemName, setItemName] = useState("");
  const [price, setPrice] = useState(0);
  const [itemId, setItemId] = useState(0);


  //Sorting by column name
  const [reverse, setReverse] = useState(false);
  const [currentSort, setCurrentSort] = useState(null);

  //Modal
  const [show, setShow] = useState(false);
  const [modalTitle, setModalTitle] = useState(null);

  //checks
  const [successful, setSuccessful] = useState(false);
  const { message } = useSelector(state => state.message);


  const dispatch = useDispatch();


  if (items && items.length === 0) {
    dispatch(getAllItems(currentUser.id));
  }

  const onChangeItemName = (e) => {
    const itemName = e.target.value;
    setItemName(itemName);
  };

  const onChangePrice = (e) => {
    const price = e.target.value;
    setPrice(price);
  };

  const handleAddItem = () => {
    setSuccessful(false);
    setShow(false);
    let itemObj = {
      userId: currentUser.id,
      itemName: itemName,
      price: parseInt(price)
    }
    console.log(itemObj);
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
  };

  const handleUpdateItem = (e) => {
    setShow(false);
    setSuccessful(false);
    let itemUpdateObj = {
      itemName: itemName,
      price: parseInt(price)
    }
    console.log("manageItems: 108", itemUpdateObj)
    ItemService.updateSpecItem(currentUser.id, itemId, itemUpdateObj).then(
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
  };

  const handleDeleteItem = () => {
    ItemService.deleteSpecItem(currentUser.id, itemId).then(
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

  const handleSortItems = (query) => {
    if (currentSort === query['sort']) {
      query['reverse'] = !reverse;
      setReverse(!reverse);
    } else {
      //if clicked column first time, reverse = false
      setCurrentSort(query['sort']);
      setReverse(false);
      query['reverse'] = false;
    }
    dispatch(getAllItems(currentUser.id, query));
  }

  //if clicked column second time, switch reverse

  const createRows = (items) => {
    return items.map(item => {
      return (
        <tr className={itemId && item.id === itemId ? "table-active" : ""} key={item.id} onClick={() => setItemId(item.id)}>
          <th scope="row">{item.id}</th>
          <td>{item.itemName}</td>
          <td>{item.price}</td>
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
            <Form onSubmit={handleUpdateItem} ref={form}>
              <div>
                <div className="form-group">
                  <label htmlFor="itemName">Item Name</label>
                  <Input
                    type="text"
                    className="form-control"
                    name="itemName"
                    value={itemName}
                    onChange={onChangeItemName}
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
                  />
                </div>
              </div>
            </Form>
          </Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={handleClose}>
              Close
            </Button>
            {modalTitle === "Add Item" ? <Button variant="primary" onClick={handleAddItem}>
              Add Item
            </Button> : <Button variant="primary" onClick={handleUpdateItem}>
              Update Item
            </Button>}
          </Modal.Footer>
        </Modal>
        <Button onClick={() => handleShow("Add Item")}>Insert</Button>
        {itemId ? <><Button onClick={() => handleShow("Update Item")}>Edit</Button>
          <Button onClick={handleDeleteItem}>Delete</Button> </> : <></>}
        <table class="table table-hover">
          <thead>
            <tr>
              <th scope="col"><Button variant="light" size="sm" onClick={() => { handleSortItems({ sort: "id", reverse }) }}>Item ID</Button></th>
              <th scope="col"><Button variant="light" size="sm" onClick={() => { handleSortItems({ sort: "item_name", reverse }) }}>Item Name</Button></th>
              <th scope="col"><Button variant="light" size="sm" onClick={() => { handleSortItems({ sort: "price", reverse }) }}>Price</Button></th>
            </tr>
          </thead>
          <tbody>
            {items && items.length > 0 &&
              createRows(items)
            }
          </tbody>
        </table>
        {!items && (
          <div class="alert alert-primary">
            <h6>No item added yet! Try inserting some items and remember to fill out every field!</h6>
          </div>
        )}
      </div>
    </div>
  );
};

export default ManageItems;
