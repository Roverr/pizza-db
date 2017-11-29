import React from 'react';
import ReactTable from 'react-table';
import "react-table/react-table.css";
import axios from 'axios';


export default class Table extends React.Component {
  constructor(props) {
    super(props);
    this.url = this.props.url;
    this.data = [];
    this.columns = [];
  }

  render() {
    this.setColumns();
    this.fetch();
    return (<ReactTable
      data={this.data}
      columns={this.columns}
    />);
  }

  setColumns() {
    this.columns = [];
  }

  async fetch() {
    if (this.data.length > 1) {
      return;
    }
    const resp = await axios.get(this.url);
    this.data = resp.data.data;
    console.log(resp.data);
    this.forceUpdate();
    return
  }
}
