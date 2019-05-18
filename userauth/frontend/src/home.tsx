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
        e.preventDefault()
        axios.delete('api/users/logout').then((res: AxiosResponse) => {
            console.log(`user ${res.data.name} has logged out`)
            this.props.handleClearUser();
        }).catch((err: any) => {
            console.log(err.response.data.error)
        })
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