import 'bootstrap/dist/css/bootstrap.css';
import React from 'react';
import ReactDOM from 'react-dom';
import Menu from './components/Menu.jsx';
import Pizza from './components/Pizza.jsx';
import Customer from './components/Customer.jsx';
import Order from './components/Order.jsx';
import { Jumbotron } from 'reactstrap';


class ViewHandler {
  constructor(url, dom) {
    this.base = url;
    this.dom = dom;
    this.pizzas();
  }
  pizzas() {
    this.dom.render(<h1> Pizzas</h1>, document.getElementById('title'));
    this.dom.render(<Pizza url={this.base} />, document.getElementById('root'));
  }

  customers() {
    this.dom.render(<h1> Customers</h1>, document.getElementById('title'));
    this.dom.render(<Customer url={this.base} />, document.getElementById('root'));
  }

  orders() {
    this.dom.render(<h1> Orders</h1>, document.getElementById('title'));
    this.dom.render(<Order url={this.base} />, document.getElementById('root'));
  }
}

const vh = new ViewHandler('http://192.168.99.100:8080', ReactDOM);
ReactDOM.render(<Menu viewHandler={vh} />, document.getElementById('menu'));
