import React from 'react';
import axios from 'axios';

export default class Button extends React.Component {
  constructor(props) {
    super(props);
    this.renderPizza = this.props.renderPizza;
  }
  fetchPizzas() {
    this.renderPizza();
  }
  render() {
    return (
      <button
        className="btn btn-default"
        onClick={this.fetchPizzas.bind(this)}>
      Pizzas
      </button>
    );
  }
}
