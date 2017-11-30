import React from 'react';
import { Button, Container, Row, Col } from 'reactstrap';


export default class Menu extends React.Component {
  constructor(props) {
    super(props);
    this.viewHandler = this.props.viewHandler;
    this.toggle = this.toggle.bind(this);
    this.state = { isOpen: false };
  }
  toggle() {
   this.setState({
     isOpen: !this.state.isOpen
   });
 }
  fetchPizzas() {
    this.viewHandler.pizzas();
  }
  fetchCustomers() {
    this.viewHandler.customers();
  }
  fetchOrders() {
    this.viewHandler.orders();
  }

  render() {
    return (
      <Container>
        <Row>
          <Col xs="6" sm="4">
            <Button
              color="danger"
              onClick={this.fetchPizzas.bind(this)}>
              Fetch Pizzas
            </Button>
          </Col>
          <Col xs="6" sm="4">
            <Button
              color="danger"
              onClick={this.fetchCustomers.bind(this)}>
              Fetch Customers
            </Button>
          </Col>
          <Col xs="6" sm="4">
            <Button
              color="danger"
              onClick={this.fetchOrders.bind(this)}>
              Fetch Orders
            </Button>
          </Col>
        </Row>
      </Container>
    );
  }
}
