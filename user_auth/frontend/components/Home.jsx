import React from 'react';


class Home extends React.Component {
  render() {
    return (
      <section>
        <h1>Hi {this.props.currentUser.name}</h1>
      </section>
    );
  }
}

export default Home;
