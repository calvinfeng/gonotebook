import * as React from 'react'
import * as ReactDOM from 'react-dom'
import './index.scss'

type IndexProps = {}

type IndexState = {
  username: string,
  message: string,
  payloads: Payload[],
  connected: boolean,
  error: string
}

type Payload = {
  username: string,
  message: string
}

class Index extends React.Component<IndexProps, IndexState> {
  private websocket: WebSocket

  constructor(props) {
    super(props)
    this.state = {
      payloads: [],
      username: "",
      message: "",
      connected: false,
      error: ""
    }
  }

  handleWebsocketOpen = (ev: Event) => {
    this.setState({ connected: true })
  }
  
  handleWebsocketClose = (cev: CloseEvent) => {
    this.setState({ connected: false })
  }
  
  handleWebsocketMessage = (mev: MessageEvent) => {
    const newPayload = JSON.parse(mev.data);
    this.setState({ payloads: this.state.payloads.concat([newPayload]) })
  }
  
  handleWebsocketError = (ev: Event) => {
    this.setState({ error: "encountered websocket error " + ev })
  }

  handleUsernameChange = (fev: React.FormEvent<HTMLInputElement>) => {
    this.setState({ username: fev.currentTarget.value });
  };

  handleMessageChange = (fev: React.FormEvent<HTMLInputElement>) => {
    this.setState({ message: fev.currentTarget.value });
  };

  handleFormSubmit = (fev: React.FormEvent<HTMLFormElement>) => {
    fev.preventDefault();
    
    const p: Payload = {
      username: this.state.username,
      message: this.state.message
    };

    if (this.websocket && this.websocket.readyState === WebSocket.OPEN) {
      this.websocket.send(JSON.stringify(p));
    }

    // Clear the message after submission.
    this.setState({ message: '' });
  };

  componentDidMount() {
    this.websocket = new WebSocket("ws://localhost:8000/streams/messages")
    this.websocket.onopen = this.handleWebsocketOpen
    this.websocket.onclose = this.handleWebsocketClose
    this.websocket.onmessage = this.handleWebsocketMessage
    this.websocket.onerror = this.handleWebsocketError
  }

  get connectionState() {
    if (this.state.connected) {
      return <p>Socket is connected</p>;
    }

    return <p>Attempting to connect</p>;
  }

  get messageRecord() {
    const msgs = this.state.payloads.map((payload) => {
      return <p key={Math.random()}>{payload.username}: {payload.message}</p>;
    });

    return (
      <article className="messsages">
        <h1>Messages</h1>
        {msgs}
      </article>
    )
  }

  get messageInput() {
    return (
      <form onSubmit={this.handleFormSubmit} className="input-form">
        <label>
          Username:<input type="text" name="username" onChange={this.handleUsernameChange} value={this.state.username} />
        </label>
        <label>
          Message:<input type="text" name="message" onChange={this.handleMessageChange} value={this.state.message} />
        </label>
        <input type="submit" value="Submit" />
      </form>
    )
  }
  
  render() {
    return (
      <section className="chatroom">
        {this.connectionState}
        {this.messageRecord}
        {this.messageInput}
      </section>
    )
  }
}

document.addEventListener("DOMContentLoaded", () => {
    ReactDOM.render(<Index />, document.getElementById("root"))
  })
  