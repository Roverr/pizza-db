import React from 'react';
import ReactTable from 'react-table';
import "react-table/react-table.css";
import axios from 'axios';


export default class Table extends React.Component {
  constructor(props) {
    super(props);
    this.url = this.props.url;
    this.state = { items: [] };
    this.columns = [];
  }

  getTable() {
    return (
    <ReactTable
      className='-striped'
      data={this.state.items}
      columns={this.columns}
    />);
  }

  componentDidMount() {
    this.fetch();
  }

  render() {
    this.setColumns();
    return this.getTable();
  }

  setColumns() {
    this.columns = [];
  }

  final(data) { return data }

  async fetch() {
    if (this.state.items > 1) {
      return;
    }
    const resp = await axios.get(this.url);
    this.setState({ items: this.final(resp.data.data) })
    return
  }
}
