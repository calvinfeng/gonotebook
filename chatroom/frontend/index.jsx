// Node modules
import React from 'react';
import ReactDOM from 'react-dom';
import uuid from 'uuid';

// Stylesheet
import "./index.scss";


class Application extends React.Component {
  state = {
    payloads: [],
    message: '',
    username: '',
    email: ''
  };

  handleSocketOpen = () => {
    this.setState({ connected: true });
  };

  handleSocketClose = () => {
    this.setState({ connected: false });
  };

  handleSocketMessage = (e) => {
    const newPayload = JSON.parse(e.data);
    this.setState({ payloads: this.state.payloads.concat([newPayload]) });
  };

  handleSocketError = (err) => {
    console.log(err);
  };

  handleUsernameChange = (e) => {
    this.setState({ username: e.target.value });
  };

  handleMessageChange = (e) => {
    this.setState({ message: e.target.value });
  };

  handleEmailChange = (e) => {
    this.setState({ email: e.target.value });
  }

  handleSubmit = (e) => {
    const payload = {
      username: this.state.username,
      email: this.state.email,
      message: this.state.message
    };

    if (this.websocket && this.websocket.readyState === WebSocket.OPEN) {
      this.websocket.send(JSON.stringify(payload));
    }

    this.setState({ message: '' });
    e.preventDefault();
  };

  componentDidMount() {
    this.websocket = new WebSocket('ws://localhost:8000/ws');
    this.websocket.onopen = this.handleSocketOpen;
    this.websocket.onmessage = this.handleSocketMessage;
    this.websocket.onerror = this.handleSocketError;
    this.websocket.onclose = this.handleSocketClose;
  }

  get messageRecordList() {
    return this.state.payloads.map((payload) => {
      return <p key={uuid()}>{payload.username}: {payload.message}</p>;
    });
  }

  get connectionState() {
    if (this.state.connected) {
      return <p>Socket is connected</p>;
    }

    return <p>Attempting to connect</p>;
  }

  render() {
    return (
      <section id="react-application">
        {this.connectionState}
        <h1>Messages</h1>
        <div className="message-record-list">
          {this.messageRecordList}
        </div>
        <h1>Input</h1>
        <form onSubmit={this.handleSubmit}>
          <label>
            Email:
            <input type="text" name="username" onChange={this.handleEmailChange} value={this.state.email} />
          </label>
          <label>
            Username:
            <input type="text" name="username" onChange={this.handleUsernameChange} value={this.state.username} />
          </label>
          <label>
            Message:
            <input type="text" name="message" onChange={this.handleMessageChange} value={this.state.message} />
          </label>
          <input type="submit" value="Submit" />
        </form>
      </section>
    );
  }
}

document.addEventListener('DOMContentLoaded', () => {
  ReactDOM.render(<Application />, document.getElementById('react-root'));
});
