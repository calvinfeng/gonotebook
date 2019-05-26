import * as React from 'react'
import axios, { AxiosResponse } from 'axios'
import RaisedButton from '@material-ui/core/Button'
import { User } from './index'

interface HomeProps {
    handleClearUser: () => void
    currentUser: User
}

class Home extends React.Component<HomeProps, any> {
    handleLogout = (e: React.MouseEvent<HTMLElement, MouseEvent>) => {
        // TODO: reset token on server side
        e.preventDefault()
        localStorage.clear()
        this.props.handleClearUser()
    }

    render() {
        return (
            <section>
                <h1>Hi {this.props.currentUser.name} </h1>
                <RaisedButton onClick={this.handleLogout}>Logout</RaisedButton>
            </section>
        )
    }
}

export default Home