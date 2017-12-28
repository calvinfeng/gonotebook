// Node modules
import React from 'react';
import ReactDOM from 'react-dom';

// Stylesheet
import "./index.scss";


class Application extends React.Component {
  state = {
    messages: []
  };
  
  handleSocketOpen = () => {
    console.log('Opened');
  }
  
  componentDidMount() {
    this.websocket = new WebSocket('ws://localhost:8000/ws');
    this.websocket.onopen = this.handleSocketOpen
  }
  
  render() {
    return (
      <div id="react-application">
        <h1>Messages</h1>
      </div>
    );
  }
}

document.addEventListener('DOMContentLoaded', () => {
  ReactDOM.render(<Application />, document.getElementById('react-root'));
});