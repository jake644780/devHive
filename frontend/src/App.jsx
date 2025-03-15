import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg';
import { BrowserRouter as Router, Routes, Route} from "react-router-dom";
import {Navbar} from './assets/navbar';
import {Home} from './assets/home';
import {AddIdea} from './assets/addIdea';

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home/>}/>
        <Route path="/AddIdea" element={<AddIdea/>} />
      </Routes>
    </Router>
  );
}

export default App;