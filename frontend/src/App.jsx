import { BrowserRouter as Router, Routes, Route} from "react-router-dom";

import {Navbar} from './assets/components/Navbar';
import {Home} from './assets/components/Home';
import {AddIdea} from './assets/components/AddIdea';
import {Footer} from './assets/components/Footer';
import {Ideas} from './assets/components/Ideas';


function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home/>}/>
        <Route path="/AddIdea" element={<AddIdea/>} />
        <Route path="/Ideas" element={<Ideas/>} />
      </Routes>
      <Footer />
    </Router>
  );
}

export default App;