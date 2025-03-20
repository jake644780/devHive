import styles from "./style.module.scss";
import { Link } from "react-router-dom";

export default function Navbar(){
    return(
        <div className={styles.navbar}>
                <div className={styles.element}>
                    <Link to="/">Home</Link>
                </div>
                <div className={styles.element}>
                    <Link to="/AddIdea">addIdea</Link>
                </div>
                <div className={styles.element}>
                    <Link to="/Ideas">Ideas</Link>
                </div>
        </div>
    );
}
