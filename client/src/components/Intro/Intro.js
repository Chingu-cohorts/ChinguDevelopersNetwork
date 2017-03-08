import React from 'react';
import './Intro.css';

const Intro = () => (
  <section className="Intro">
    <h2>A gym for programmers</h2>
    <div className="Intro-image"></div>
    <div className="Intro-link-container">
      <a href="#" className="Intro-link">
        <span className="Intro-link-title">Sign Up</span>
      </a>
      <a href="#" className="Intro-link">
        <span className="Intro-link-title">Apply</span>
      </a>      
    </div>
    <hr className="Intro-line" />
  </section>
);

export default Intro;
