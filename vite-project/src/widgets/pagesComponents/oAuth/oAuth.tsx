import styles from "../oAuth/oAuth.module.css";
import GoogleImg from "../../../assets/img/oAuth/google.png";
import vk from "../../../assets/img/oAuth/vk.png";
import gitHub from "../../../assets/img/oAuth/gitHub.png";
export default function OAuth() {
  return (
    <>
      <p className={styles.subTitle}>Войти с помощью</p>
      <div className={styles["container"]}>
        <div className={styles['grids']}>
          <a href="#">
            <img
              src={GoogleImg}
              className={styles.googleImg}
              alt="googleAuth"
            />
          </a>

          <a href="#">
            <img src={vk} alt="vkAuth" className={styles.vkImg} />
          </a>
          <a href="#">
            <img src={gitHub} alt="gitHub" className={styles.gitHubImg} />
          </a>
        </div>
      </div>
    </>
  );
}
