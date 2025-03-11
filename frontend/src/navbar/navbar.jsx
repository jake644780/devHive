import styles from "./style.module.css";

export default function Navbar(){
    return (
        <nav className={styles.navbar}>
            <div className={styles.home}>
                Home
            </div>
            <ul>
                <li>idk1</li>
                <li>idk2</li>
                <li>idk3</li>
            </ul>
        </nav>
    );
}