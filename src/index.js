import React from 'react';
import ReactDOM from 'react-dom';
import App from './components/App.jsx';
import Button from './components/Buttons.jsx';
import Pizza from './components/Pizza.jsx';

ReactDOM.render(<App />, document.getElementById('root'));
ReactDOM.render(<Button renderPizza={renderPizza} />, document.getElementById('menu'));

function renderPizza(){
  ReactDOM.render(<Pizza />, document.getElementById('root'));
}
