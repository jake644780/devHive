import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg';
import {Navbar} from './navbar/index';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";



function App() {
  return (
    <Router>
      <Navbar />
      <div className="container">
        
      </div>
    </Router>
  );
}

export default App;