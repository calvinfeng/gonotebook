import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import lightBaseTheme from 'material-ui/styles/baseThemes/lightBaseTheme';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import Home from './components/Home';
import Welcome from './components/Welcome';
import './index.scss';


class Application extends React.Component {
  state = {};

  handleSetCurrentUser = (user) => {
    this.setState({ currentUser: user });
  };

  handleClearCurrentUser = () => {
    this.setState({ currentUser: undefined });
  };

  componentDidMount() {
    axios.get('api/authenticate').then((res) => {
      this.setState({ currentUser: res.data });
    }).catch((err) => {
      console.log('No current user');
    });
  }

  get content() {
    if (this.state.currentUser) {
      return <Home currentUser={this.state.currentUser} handleClearCurrentUser={this.handleClearCurrentUser} />;
    }

    return <Welcome handleSetCurrentUser={this.handleSetCurrentUser} />;
  }

  render() {
    return (
      <MuiThemeProvider muiTheme={getMuiTheme(lightBaseTheme)}>
        <article>
          <h1>User Authentication in Go</h1>
          {this.content}
        </article>
      </MuiThemeProvider>
    );
  }
}

document.addEventListener("DOMContentLoaded", () => {
    ReactDOM.render(<Application />, document.getElementById('react-application'));
});
