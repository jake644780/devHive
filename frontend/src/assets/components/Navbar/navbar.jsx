import styles from "./style.module.css";
import { Link } from "react-router-dom";
import logo from "../../assets/images/logo.png";

export default function Navbar(){
    return(
        <div className={styles.navbar}>
            <div className={styles.logo}><img src="" alt="" /></div>
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
