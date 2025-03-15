import { BrowserRouter as Router, Routes, Route} from "react-router-dom";

import {Navbar} from './assets/Navbar';
import {Home} from './assets/Home';
import {AddIdea} from './assets/AddIdea';

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home/>}/>
        <Route path="/AddIdea" element={<AddIdea/>} />
      </Routes>
      <Footer />
    </Router>
  );
}

export default App;