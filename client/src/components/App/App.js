import React, { Component } from 'react';
import Header from '../Header/Header';
import Intro from '../Intro/Intro';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Header />

        <main className="App-container">
          <Intro />
        </main>
      </div>
    );
  }
}

export default App;
