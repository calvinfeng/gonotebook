import React from 'react';
import axios from 'axios';
import RaisedButton from 'material-ui/RaisedButton';


class Home extends React.Component {
  handleLogout = () => {
    axios.delete('api/logout').then((res) => {
      console.log(`User ${res.data.name} has logged out.`);
      this.props.handleClearCurrentUser();
    }).catch((err) => {
      console.log(err.response.data.error);
    })
  };

  render() {
    return (
      <section>
        <h1>Hi {this.props.currentUser.name}</h1>
        <RaisedButton label={'Logout'} onClick={this.handleLogout} />
      </section>
    );
  }
}

export default Home;
