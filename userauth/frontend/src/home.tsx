import * as React from 'react'
import axios, { AxiosResponse } from 'axios'
import RaisedButton from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import { User, Message } from './shared'
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

interface HomeProps {
    handleClearUser: () => void
    currentUser: User
}

interface HomeState {
    received: Message[]
    sent: Message[]
    recipients: User[]
    receiverID: number
    inputFieldText: string
    errorText: string
}

class Home extends React.Component<HomeProps, HomeState> {
    constructor(props) {
        super(props)

        this.state = {
            received: [],
            sent: [],
            recipients: [],
            inputFieldText: "",
            errorText: "",
            receiverID: 0
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

    handleMessageOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault()
        const newState: HomeState = Object.assign({}, this.state)
        newState.inputFieldText = e.target.value
        this.setState(newState)
    }

    handleReceiverOnChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        e.preventDefault()
        const newState: HomeState = Object.assign({}, this.state)
        newState.receiverID = parseInt(e.target.value)
        this.setState(newState)
    }

    handleLogout = (e: React.MouseEvent<HTMLElement, MouseEvent>) => {
        e.preventDefault()
        localStorage.clear()
        this.props.handleClearUser()
    }

    handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault() 
        const token = localStorage.getItem("jwt_token")
        if (this.state.inputFieldText.length > 0 && token && this.state.receiverID !== 0) {
            axios.post("api/messages/", {
                body: this.state.inputFieldText,
                receiver_id: this.state.receiverID
            },{
                headers: { "Token": token }
            }).then((res: AxiosResponse) => {
                this.fetchSentMessages()
                this.fetchReceivedMessages()
            }).catch((err) => {
                console.error("failed to create message")
            })
            
            const newState: HomeState = Object.assign({}, this.state)
            newState.inputFieldText = ""
            newState.errorText = ""
            this.setState(newState)
        }

        let errorText = ""
        if (this.state.inputFieldText.length === 0) {
            errorText = "please write a message please you click send"
        }

        if (this.state.receiverID === 0) {
            errorText = "please select a recipient before you click send"
        }

        if (errorText.length > 0) {
            const newState: HomeState = Object.assign({}, this.state)
            newState.errorText = errorText
            this.setState(newState)
        }
    }

    get composeForm() {
        return (
            <form onSubmit={this.handleSubmit}>
                {this.selector}<br/>
                <TextField placeholder={"Type your message"} onChange={this.handleMessageOnChange} value={this.state.inputFieldText} />
                <RaisedButton type="submit">Send</RaisedButton>
            </form>
        )   
    }

    get receivedMessages() {
        const items = this.state.received.map((msg: Message) => {
            return (
                <ListItem key={msg.id}>
                    <ListItemText primary={msg.body} secondary={`from ${msg.sender.name}`} />
                </ListItem>
            )
        })

        return (
            <section>
                <h2>Received</h2>
                <List>{items}</List>
            </section>
        )
    }
    
    get sentMessages() {
        const items = this.state.sent.map((msg: Message) => {
            return (
                <ListItem key={msg.id}>
                    <ListItemText primary={msg.body} secondary={`to ${msg.receiver.name}`} />
                </ListItem>
            )
        })
        return (
            <section>
                <h2>Sent</h2>
                <List>{items}</List>
            </section>
        )
    }

    get selector() {
        const items = this.state.recipients.map((user: User) => {
            return <MenuItem key={user.id} value={user.id}>{user.name}</MenuItem>
        })

        return (
            <section>
                <h2>Compose</h2>
                <FormControl className="form-control">
                    <Select value={this.state.receiverID} onChange={this.handleReceiverOnChange}>{items}</Select>
                </FormControl>
            </section>
        )
    }

    get errorTextSection() {
        if (this.state.errorText === "") {
            return <section></section>
        }

        return (
            <section>
                <h2>Error</h2>
                <p>{this.state.errorText}</p>
            </section>
        )
    }

    render() {
        return (
            <section>
                <h2>Hi {this.props.currentUser.name}</h2>
                <RaisedButton onClick={this.handleLogout}>Logout</RaisedButton>
                {this.receivedMessages}
                {this.sentMessages}
                {this.composeForm}
                {this.errorTextSection}
            </section>
        )
    }
}

export default Home