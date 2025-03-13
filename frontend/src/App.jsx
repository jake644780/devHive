import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg';
import {Navbar} from './assets/navbar/index';
import { BrowserRouter as Router} from "react-router-dom";
import {Routess} from "./routes/index";

function App() {
  return (
    <Router>
      <Routess />
      <Navbar />
      <div className="container">
        
      </div>
    </Router>
  );
}

export default App;