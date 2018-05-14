import React from 'react';
import ReactDOM from 'react-dom';
import App from './components/app';

document.addEventListener('DOMContentLoaded', function() {
    console.log("asd")
    ReactDOM.render(
        React.createElement(App),
        document.getElementById('mount')
    );
});