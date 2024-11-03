import { Header } from "../../widgets/pageComponents/Header";
import cl from "../ProfilePage/Profile.module.css";
import avatar from "./../../assets/img/avatar.png";
import { NavLink } from "react-router-dom";
export default function Profile() {
  return (
    <>
      <Header></Header>
      <div className={cl["container"]}>
        <div className={cl["header-logo"]}>
          <img src={avatar} alt="avatarUrl" />
        </div>
        <div className={cl["username"]}>Irayem</div>
        <div className={cl["container__description"]}>
          <NavLink
            to={"/myfavourite"}
            className={cl["container__description-link"]}
          >
            Любимые <br />и прочитанные
          </NavLink>
          <textarea className={cl["description__input"]} />
          <NavLink
            to={"/myfriends"}
            className={cl["container__description-link"]}
          >
            Друзья
          </NavLink>
        </div>
      </div>
    </>
  );
}
