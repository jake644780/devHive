import styles from "./style.module.scss";

export default function Home(){
    return(
        <div className={styles.main}>
            <div className="title">
                Welcome back
            </div>
        </div>
    );
}