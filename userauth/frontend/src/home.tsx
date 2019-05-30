import * as React from 'react'
import axios, { AxiosResponse } from 'axios'
import RaisedButton from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import { User, Message } from './shared'
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import InputLabel from '@material-ui/core/InputLabel';

interface HomeProps {
    handleClearUser: () => void
    currentUser: User
}

interface HomeState {
    received: Message[]
    sent: Message[]
    recipients: User[]
    inputFieldText: string
}

class Home extends React.Component<HomeProps, HomeState> {
    constructor(props) {
        super(props)

        this.state = {
            received: [],
            sent: [],
            recipients: [],
            inputFieldText: ""
        }
    }

    fetchUsers() {
        const token = localStorage.getItem("jwt_token")
        if (token) {
            axios.get("api/users/", {
                headers: { "Token": token }
            }).then((res: AxiosResponse) => {
                this.setState({
                    recipients: res.data
                })
            }).catch((err) => {
                console.log("failed to fetch messages for current user")
            })
        }
    }

    fetchSentMessages() {
        const token = localStorage.getItem("jwt_token")
        if (token) {
            axios.get("api/messages/sent/", {
                headers: { "Token": token }
            }).then((res: AxiosResponse) => {
                this.setState({
                    sent: res.data
                })
            }).catch((err) => {
                console.log("failed to fetch messages for current user")
            })
        }
    }

    fetchReceivedMessages() {
        const token = localStorage.getItem("jwt_token")
        if (token) {
            axios.get("api/messages/received/", {
                headers: { "Token": token }
            }).then((res: AxiosResponse) => {
                this.setState({
                    received: res.data
                })
            }).catch((err) => {
                console.log("failed to fetch messages for current user")
            })
        }
    }

    componentDidMount() {
        this.fetchSentMessages()
        this.fetchReceivedMessages()
        this.fetchUsers()
    }

    handleLogout = (e: React.MouseEvent<HTMLElement, MouseEvent>) => {
        e.preventDefault()
        localStorage.clear()
        this.props.handleClearUser()
    }

    handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        const token = localStorage.getItem("jwt_token")
        if (this.state.inputFieldText.length > 0 && token) {
            axios.post("api/messages/", {
                body: this.state.inputFieldText 
            },{
                headers: { "Token": token }
            }).then((res: AxiosResponse) => {
                this.fetchSentMessages()
            }).catch((err) => {
                console.error("failed to create message")
            })
            
            const newState: HomeState = Object.assign({}, this.state)
            newState.inputFieldText = ""
            this.setState(newState)
        }
    }

    messageOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault()
        const newState: HomeState = Object.assign({}, this.state)
        newState.inputFieldText = e.target.value
        this.setState(newState)
    }

    get form() {
        return (
            <form onSubmit={this.handleSubmit}>
                {this.selector}<br/>
                <TextField placeholder={"Type your message"} onChange={this.messageOnChange} value={this.state.inputFieldText} />
                <RaisedButton type="submit">Send</RaisedButton>
            </form>
        )   
    }

    get receivedMessages() {
        console.log(this.state.received)
        const items = this.state.received.map((msg: Message) => {
            return <li key={msg.id}>From {msg.sender.name}: {msg.body}</li>
        })

        return <ul>{items}</ul>
    }
    
    get sentMessages() {
        console.log(this.state.sent)
        const items = this.state.received.map((msg: Message) => {
            return <li key={msg.id}>To {msg.receiver.name}: {msg.body}</li>
        })

        return <ul>{items}</ul>
    }

    get selector() {
        const items = this.state.recipients.map((user: User) => {
            return <MenuItem value={user.id}>{user.name}</MenuItem>
        })

        return (
            <FormControl>
                <InputLabel>Send to</InputLabel>
                <Select value={10}>{items}</Select>
            </FormControl>
        )
    }

    render() {
        return (
            <section>
                <h1>Hi {this.props.currentUser.name} </h1>
                <h2>Received Messages</h2>
                {this.receivedMessages}
                <h2>Sent Messages</h2>
                {this.sentMessages}
                {this.form}
                <RaisedButton onClick={this.handleLogout}>Logout</RaisedButton>
            </section>
        )
    }
}

export default Home