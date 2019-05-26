import * as React from 'react'
import * as ReactDOM from 'react-dom'
import axios, { AxiosResponse } from 'axios'
import Button from '@material-ui/core/Button'
import Home from './home'
import Welcome from './welcome'
import './index.scss'

interface IndexState {
    user: User       
}

export type User = {
    name: string
    email: string
    jwt: string
}

class Index extends React.Component<any, IndexState> {
    constructor(props) {
        super(props)

        this.state = {
            user: null
        }
    }

    handleSetUser = (user: User) => {
        this.setState({ user })
    }

    handleClearUser = () => {
        this.setState({ user: null })
    }

    componentDidMount() {
        const token = localStorage.getItem("jwt_token")
        if (token) {
            axios.get("api/users/current", {
                headers: { "Token": token }
            }).then((res: AxiosResponse) => {
                this.setState({ user: res.data })
            }).catch((err) => {
                console.log("no current user")
            })
        }
    }

    get content() {
        if (this.state.user) {
            return <Home currentUser={this.state.user} handleClearUser={this.handleClearUser} />
        }

        return <Welcome handleSetUser={this.handleSetUser} />
    }

    render() {
        return (
            <article>
                <h1>User Authentication in Go</h1>
                {this.content}
            </article>
        )
    }
}

document.addEventListener("DOMContentLoaded", () => {
    ReactDOM.render(<Index />, document.getElementById('react-application'))
})