import React from 'react';
import ReactTable from 'react-table';
import "react-table/react-table.css";
import Table from './Table'


export default class Order extends Table {
  constructor(props) {
    super(props);
    this.url = props.url + '/api/orders';
    this.original = {};
  }

  setColumns() {
    this.columns = [{
      Header: 'ID',
      accessor: 'id',
    }, {
      Header: 'Address',
      accessor: 'address',
    }, {
      Header: 'Price',
      accessor: 'price',
    }, {
      Header: 'Customer name',
      accessor: 'customer.name',
    }, {
      Header: 'Started',
      accessor: 'startedAt',
    }, {
      Header: 'Finished',
      accessor: 'completedAt',
    }];
  }

  setPizzaColumns() {
    this.pizzaColumns = [{
      Header: 'Name',
      accessor: 'name',
    }, {
      Header: '# of pizzas in order',
      accessor: 'number',
    }];
  }

  getTable() {
    return (
    <ReactTable
      className='-striped -highlight'
     SubComponent={row => {
       this.setPizzaColumns();
       return (
          <div style={{ padding: "20px" }}>
            <em>
              Pizzas for order <b>{row.original.id}</b>
            </em>
            <br />
            <br />
            <ReactTable
              data={row.original.pizzas}
              columns={this.pizzaColumns}
              defaultPageSize={3}
              showPagination={false}
            />
          </div>
        );
          }}
      data={this.state.items}
      columns={this.columns}
    />);
  }
}
