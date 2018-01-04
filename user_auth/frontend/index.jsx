import React from 'react';
import ReactDOM from 'react-dom';
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

  get content() {
    if (this.state.currentUser) {
      return <Home currentUser={this.state.currentUser} />;
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
