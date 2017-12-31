import React from 'react';
import ReactDOM from 'react-dom';
import './index.scss';


class Application extends React.Component {
  render() {
    return (
      <section>
        <h1>User Authentication</h1>
      </section>
    );
  }
}

document.addEventListener("DOMContentLoaded", () => {
    ReactDOM.render(<Application />, document.getElementById('react-application'));
});