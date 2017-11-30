import React from 'react';
import "react-table/react-table.css";
import Table from './Table'


export default class Customer extends Table {
  constructor(props) {
    super(props);
    this.url = props.url + '/api/customers';
    this.renderEditable = this.renderEditable.bind(this);
  }

  setColumns() {
    this.columns = [{
      Header: 'ID',
      accessor: 'id',
    }, {
      Header: 'Name',
      accessor: 'name',
      Cell: this.renderEditable,
    }, {
      Header: 'Email',
      accessor: 'email',
      Cell: this.renderEditable,
    }];
  }

  renderEditable(cellInfo) {
    return (
      <div
        style={{ backgroundColor: "#fafafa" }}
        contentEditable
        suppressContentEditableWarning
        onBlur={e => {
          const data = [...this.state.items];
          data[cellInfo.index][cellInfo.column.id] = e.target.innerHTML;
          this.setState({ data });
        }}
        dangerouslySetInnerHTML={{
          __html: this.state.items[cellInfo.index][cellInfo.column.id]
        }}
      />
    );
  }
}
