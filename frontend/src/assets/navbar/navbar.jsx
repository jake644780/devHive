import styles from "./style.module.css";
import { Link } from "react-router-dom";

export default function Navbar(){
    return(
        <div>
            <ul>
                <li>
                    <Link to="/">Home</Link>
                </li>
                <li>
                    <Link to="/addIdea">addIdea</Link>
                </li>
            </ul>
        </div>
    );
}
