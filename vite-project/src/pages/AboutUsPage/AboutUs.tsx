import { Header } from "../../widgets/pageComponents/Header";
import styles from "../AboutUsPage/AboutUs.module.css";
export function AboutUs() {
  return (
    <>
      <div className={styles.container}>
        <Header></Header>
        <div className={styles["title_1"]}>
          <p className={styles.aboutUs}>О нас</p>
          <p className={styles.subtitle}>
            LexoRead это сайт с открытым
            <br /> исходным кодом где вы сможете
            <br /> прочитать книги на ваш вкус
          </p>
        </div>

        <hr className={styles["hr"]} />
        <div className={styles["title_2"]}>
          <p className={styles["designer"]}>Дизайнер</p>
          <p className={styles["designer-nickname"]}>Irayem</p>
        </div>
        <hr className={styles["hr"]} />
        <div className={styles["title_3"]}>
          <p className={styles["developers"]}>Разработчики</p>
          <p className={styles["developers-nickname"]}>
            par1tet, quthery, qruhich, Smallgigach, Ogonek
          </p>
        </div>
        <hr className={styles["hr_last"]} />
      </div>
    </>
  );
}
