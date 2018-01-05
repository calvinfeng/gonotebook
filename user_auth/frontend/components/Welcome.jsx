import React from 'react';
import RaisedButton from 'material-ui/RaisedButton';
import TextField from 'material-ui/TextField';
import axios from 'axios';

const LOGIN = 'LOGIN';
const REGISTER = 'REGISTER';

class Welcome extends React.Component {
  state = {
    form: REGISTER,
    name: '',
    email: '',
    password: '',
    error: ''
  };

  createFormStateHandler = (formState) => {
    return () => {
      this.setState({ form: formState });
    };
  };

  handleRegisterSubmit = (e) => {
    e.preventDefault();
    axios.post('/api/users/register', {
      name: this.state.name,
      email: this.state.email,
      password: this.state.password
    }).then((res) => {
      this.props.handleSetCurrentUser(res.data);
    }).catch((err) => {
      this.setState({ error : err.response.data.error });
    });
  };

  handleLoginSubmit = (e) => {
    e.preventDefault();
    axios.post('/api/users/login', {
      email: this.state.email,
      password: this.state.password
    }).then((res) => {
      this.props.handleSetCurrentUser(res.data);
    }).catch((err) => {
      this.setState({ error: err.response.data.error });
    });
  };

  createTextFieldChangeHandler = (fieldName) => {
    return (e, val) => {
      const newState = Object.assign({}, this.state);
      newState[fieldName] = val;
      this.setState(newState);
    };
  };

  get formStateToggle() {
    if (this.state.form === REGISTER) {
      return <RaisedButton label={'Login'} onClick={this.createFormStateHandler(LOGIN)} />;
    }

    return <RaisedButton label={'Register'} onClick={this.createFormStateHandler(REGISTER)} />;
  }

  get title() {
    if (this.state.form === REGISTER) {
        return <h2>Register</h2>;
    }

    return <h2>Login</h2>;
  }

  get form() {
    if (this.state.form === REGISTER) {
      return (
        <section>
          <form onSubmit={this.handleRegisterSubmit}>
            <TextField hintText="name" onChange={this.createTextFieldChangeHandler('name')} /><br />
            <TextField hintText="email" onChange={this.createTextFieldChangeHandler('email')} /><br />
            <TextField hintText="password" onChange={this.createTextFieldChangeHandler('password')} type="password" /><br />
            <input type="submit" label="Submit" />
          </form>
        </section>
      );
    }

    return (
      <section>
        <form onSubmit={this.handleLoginSubmit}>
          <TextField hintText="email" onChange={this.createTextFieldChangeHandler('email')} /><br />
          <TextField hintText="password" onChange={this.createTextFieldChangeHandler('password')} type="password" /><br />
          <input type="submit" label="Submit" />
        </form>
      </section>
    )
  }

  get error() {
    if (this.state.error.length > 0) {
      return <p>Error: {this.state.error}</p>;
    }

    return;
  }

  render() {
    return (
      <section>
        <h2>Welcome to Go Academy</h2>
        {this.formStateToggle}
        {this.title}
        {this.form}
        {this.error}
      </section>
    );
  }
}

export default Welcome;
