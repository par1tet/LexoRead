import { Header } from "../../widgets/pageComponents/Header";
import styles from "../ProfilePage/Profile.module.css";
import avatar from "./../../assets/img/avatar.png";
import { NavLink } from "react-router-dom";
export default function Profile() {
  return (
    <>
    <Header></Header>
      <div className={styles["container"]}>
        <div className={styles["header-logo"]}>
          <img src={avatar} alt="avatarUrl" />
        </div>
        <div className={styles["username"]}>Smallgigach</div>
				<span className={styles["sub-title"]}>О себе</span>
        <div className={styles["description"]}>
          <NavLink to={"/myfriends"}>Друзья</NavLink>
          <textarea className={styles["description__input"]} />
          <NavLink to={"/myfavourite"} >Мои любимые</NavLink>
        </div>
      </div>
    </>
  );
}
