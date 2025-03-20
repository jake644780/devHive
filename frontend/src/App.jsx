import { BrowserRouter as Router, Routes, Route} from "react-router-dom";

import {Navbar} from './assets/components/Navbar';
import {Home} from './assets/components/Home';
import {AddIdea} from './assets/components/AddIdea';
import {Footer} from './assets/components/Footer';


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