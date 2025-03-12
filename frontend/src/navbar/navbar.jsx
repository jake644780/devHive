import styles from "./style.module.css";

export default function Navbar(){

    let navItems = [];
    let navNames = [
        'home',
        'add idea',
    ];
    for (let i = 0; i < navNames.length; i++) navItems.push(
        <li className={styles.navElement}>
            { navNames[i] }
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