import React, { Component } from 'react';
import Header from '../Header/Header';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Header />

        <main className="App-container">
          <section className="App-intro">
            <h2>A gym for programmers</h2>
            <div className="App-intro-image"></div>
            <div className="App-intro-btn-container">
            	<input type="button" value="Sign Up" className="App-intro-btn" />
          	  <input type="button" value="Apply" className="App-intro-btn" />
            </div>
            <hr className="App-intro-line" />
          </section>
        </main>
      </div>
    );
  }
}

export default App;
