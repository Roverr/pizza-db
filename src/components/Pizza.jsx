import React from 'react';
import ReactTable from 'react-table';
import "react-table/react-table.css";
import axios from 'axios';
import Table from './Table'

function cellDecision(key) {
  return (row) => {
    if (row.original[key]) {
      return 'yes';
    }
    return 'no';
  }
}


export default class Pizza extends Table {
  constructor(props) {
    super(props);
    this.url = props.url + '/api/pizzas';
    this.ingredientURL = props.url + '/api/ingredients';
    this.ingredients = [];
    this.columns = [];
    this.ingredientsColumns = [];
    this.original = {};
  }

  setColumns() {
    this.columns = [{
      Header: 'ID',
      accessor: 'id',
    }, {
      Header: 'Name',
      accessor: 'name',
    }, {
      Header: 'Price',
      accessor: 'price',
    }, {
      Header: 'Ingredients available',
      accessor: 'ingredientsAvailable',
    }];
  }

  setIngredientsColumns() {
    if (this.ingredientsColumns.length > 0 ) {
      return;
    }
    this.ingredientsColumns = [{
      Header: 'ID',
      accessor: 'id',
    }, {
      Header: 'Name',
      accessor: 'name',
    }, {
      Header: 'Available',
      accessor: 'available',
      Cell: cellDecision('available'),
    }, {
      Header: 'Gluten free',
      accessor: 'glutenFree',
      Cell: cellDecision('glutenFree'),
    }];
  }

  getTable() {
    return (
    <ReactTable
      className='-striped -highlight'
     SubComponent={row => {
       this.setIngredientsColumns();
       return (
          <div style={{ padding: "20px" }}>
            <em>
              Ingredients to <b>{row.original.name}</b>
            </em>
            <br />
            <br />
            <ReactTable
              data={row.original.ingredients}
              columns={this.ingredientsColumns}
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

  final(data) {
    return data.map((item) => {
      const isAvailable = item.ingredients.every((ingredient) => ingredient.available);
      let shortcut = item;
      shortcut.ingredientsAvailable = "no";
      if (isAvailable) {
        shortcut.ingredientsAvailable = "yes";
      }
      return shortcut;
    });
  }


  async getIngredients(pizzaID, update) {
    const resp = await axios.get(`${this.ingredientURL}?pizzaId=${pizzaID}`);
  }
}
