import React from 'react';
import ReactTable from 'react-table';
import "react-table/react-table.css";
import axios from 'axios';
import Table from './Table'


export default class Pizza extends Table {
  constructor(props) {
    super(props);
    this.url = "https://b694992c.ngrok.io/api/pizzas";
  }

  setColumns() {
    this.columns = [{
      Header: 'ID',
      accessor: 'id'
    }, {
      Header: 'Name',
      accessor: 'name'
    }, {
      Header: 'Price',
      accessor: 'price',
    }];
  }
}
