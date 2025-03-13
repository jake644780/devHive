import styles from "./style.module.css";
import { Link } from "react-router-dom";



export default function Navbar(){

    let navItems = [];
    let navNames = [
        'home',
        'addIdea',
    ];
    for (let i = 0; i < navNames.length; i++) navItems.push(
        <li className={styles.navElement}>
            <Link to={"/" + navNames[i]}>
                { navNames[i] }
            </Link>
        </li>
    );
        


    return (
        <nav className={styles.navbar}>
            <div className={styles.home}>
                Home
            </div>
            <ul className={styles.navElements}>
                { navItems }
            </ul>
        </nav>
    );
}