import { Routes, Route } from "react-router-dom";
import { AddIdea } from "../assets/addIdea/index";

export default function Routess(){

    return(
    <Routes>
        <Route path="/addIdea" element={<AddIdea />} />
    </Routes>
    )
    
}