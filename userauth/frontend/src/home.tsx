import * as React from 'react'
import axios, { AxiosResponse } from 'axios'
import RaisedButton from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import { User } from './index'

export type Message = {
    body: string
    id: number
}

interface HomeProps {
    handleClearUser: () => void
    currentUser: User
}

interface HomeState {
    messages: Message[]
    textFieldInput: string
}

class Home extends React.Component<HomeProps, HomeState> {
    constructor(props) {
        super(props)

        this.state = {
            messages: [],
            textFieldInput: ""
        }
    }

    refresh() {
        const token = localStorage.getItem("jwt_token")
        if (token) {
            axios.get("api/messages/current", {
                headers: { "Token": token }
            }).then((res: AxiosResponse) => {
                this.setState({
                    messages: res.data
                })
            }).catch((err) => {
                console.log("failed to fetch messages for current user")
            })
        }
    }

    componentDidMount() {
        this.refresh()
    }

    handleLogout = (e: React.MouseEvent<HTMLElement, MouseEvent>) => {
        e.preventDefault()
        localStorage.clear()
        this.props.handleClearUser()
    }

    handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        const token = localStorage.getItem("jwt_token")
        if (this.state.textFieldInput.length > 0 && token) {
            axios.post("api/messages/", {
                body: this.state.textFieldInput 
            },{
                headers: { "Token": token }
            }).then((res: AxiosResponse) => {
                this.refresh()
            }).catch((err) => {
                console.error("failed to create message")
            })
            
            const newState: HomeState = Object.assign({}, this.state)
            newState.textFieldInput = ""
            this.setState(newState)
        }
    }

    messageOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault()
        const newState: HomeState = Object.assign({}, this.state)
        newState.textFieldInput = e.target.value
        this.setState(newState)
    }

    get form() {
        return (
            <form onSubmit={this.handleSubmit}>
                <TextField placeholder={"Type your message"} onChange={this.messageOnChange} value={this.state.textFieldInput} />
                <RaisedButton type="submit">Send</RaisedButton>
            </form>
        )   
    }

    get messages() {
        const items = this.state.messages.map((msg: Message) => {
            return <li key={msg.id}>{msg.body}</li>
        })

        return <ul>{items}</ul>
    }    

    render() {
        return (
            <section>
                <h1>Hi {this.props.currentUser.name} </h1>
                <h2>Your Messages</h2>
                {this.messages}
                {this.form}
                <RaisedButton onClick={this.handleLogout}>Logout</RaisedButton>
            </section>
        )
    }
}

export default Home