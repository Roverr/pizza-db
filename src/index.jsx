import 'bootstrap/dist/css/bootstrap.css';
import React from 'react';
import ReactDOM from 'react-dom';
import App from './components/App.jsx';
import Menu from './components/Menu.jsx';
import Pizza from './components/Pizza.jsx';


ReactDOM.render(<App />, document.getElementById('root'));
ReactDOM.render(<Menu renderPizza={renderPizza} />, document.getElementById('menu'));

function renderPizza(){
  ReactDOM.render(<Pizza />, document.getElementById('root'));
}
