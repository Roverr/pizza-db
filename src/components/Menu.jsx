import React from 'react';
import { Button } from 'reactstrap';

export default class Menu extends React.Component {
  constructor(props) {
    super(props);
    this.renderPizza = this.props.renderPizza;
  }
  fetchPizzas() {
    this.renderPizza();
  }
  render() {
    return (
      <Button
        color="danger"
        onClick={this.fetchPizzas.bind(this)}>
      Pizzas
    </Button>
    );
  }
}
