import * as React from 'react'
import RaisedButton from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import axios, { AxiosResponse } from 'axios'
import { User } from './index'
import { Button } from '@material-ui/core';

enum Page {
    Login = "LOGIN",
    Register = "REGISTER"
}

interface WelcomeProps {
    handleSetUser: (user: User) => void
}

interface WelcomeState {
    page: Page
    name: string
    email: string
    password: string
    error: string
}

class Welcome extends React.Component<WelcomeProps, WelcomeState> {
    constructor(props) {
        super(props)
        this.state = {
            page: Page.Register,
            name: '',
            email: '',
            password: '',
            error: '' 
        }
    }

    createPageStateHandler = (page: Page) => {
        return () => {
            this.setState({ page })
        }
    }

    createTextFieldOnChangeHandler = (field: string) => {
        return (e: React.ChangeEvent<HTMLInputElement>) => {
            e.preventDefault()
            const newState: WelcomeState = Object.assign({}, this.state)
            newState[field] = e.target.value
            this.setState(newState)
        }
    }

    handleRegisterSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        axios.post('/api/users/register', {
            name: this.state.name,
            email: this.state.email,
            password: this.state.password
        }).then((res: AxiosResponse) => {
            this.props.handleSetUser(res.data)
        }).catch((err: any) => {
            this.setState({ error: err.response.data.error })
        })
    }

    handleLoginSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        axios.post('api/users/login', {
            email: this.state.email,
            password: this.state.password
        }).then((res: AxiosResponse) => {
            this.props.handleSetUser(res.data)
        }).catch((err: any) => {
            this.setState({ error: err.response.data.error })
        })
    }

    get pageStateToggle() {
        if (this.state.page === Page.Register) {
            return <RaisedButton onClick={this.createPageStateHandler(Page.Login)}>Login</RaisedButton>
        }

        return <RaisedButton onClick={this.createPageStateHandler(Page.Register)}>Register</RaisedButton>
    }

    get title() {
        if (this.state.page == Page.Register) {
            return <h2>Register</h2>
        }

        return <h2>Login</h2>
    }

    get form() {
        if (this.state.page === Page.Register) {
            return (
                <form onSubmit={this.handleRegisterSubmit}>
                    <TextField placeholder={"Your Name"} onChange={this.createTextFieldOnChangeHandler("name")} />
                    <br/><br/>
                    <TextField placeholder={"Your Email"} onChange={this.createTextFieldOnChangeHandler("email")} />
                    <br/><br/>
                    <TextField placeholder={"Your Password"} type={"password"} onChange={this.createTextFieldOnChangeHandler("password")} />
                    <br/><br/>
                    <RaisedButton type="submit">Submit</RaisedButton>
                </form>
            )
        }

        return (
            <form onSubmit={this.handleLoginSubmit}>
                <TextField placeholder={"Email"} onChange={this.createTextFieldOnChangeHandler("email")} />
                <br/><br/>
                <TextField placeholder={"Password"} type={"password"} onChange={this.createTextFieldOnChangeHandler("password")} />
                <br/><br/>
                <RaisedButton type="submit">Login</RaisedButton>
            </form>
        )
    }

    get error() {
        if (this.state.error.length > 0) {
          return <p>Error: {this.state.error}</p>;
        }
    }

    render() {
        return (
            <section>
                <h2>Please login or register</h2>
                {this.pageStateToggle}
            {this.title}
            {this.form}
            {this.error}
            </section>
        )
    }
}

export default Welcome