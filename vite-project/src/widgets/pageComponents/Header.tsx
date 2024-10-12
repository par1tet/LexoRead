import cl from "./Header.module.css";
import avatar from "./../../assets/img/avatar.png";
import { NavLink } from "react-router-dom";
import search from "./../../assets/img/search.svg";

export function Header() {
  return (
    <div className={cl["header"]}>
      <div className={cl["header__links"]}>
        <div className={cl["header__links-lexoRead"]}>
          <NavLink to="/">LexoRead</NavLink>
        </div>
        <div className={cl["header__links-main"]}>
          <NavLink 
            to="/"
            className={({ isActive }) => (isActive ? cl["about_us"] : undefined)}
            >
            Главное</NavLink>
        </div>
        <div className={cl["header__links-aboutus"]}>
          <NavLink
            to="/aboutUs"
            className={({ isActive }) => (isActive ? cl["about_us"] : undefined)}
          >
            О нас
          </NavLink>
        </div>
      </div>
      <div className={cl["header__logo"]}>
        <img src={avatar} />
      </div>
    </div>
  );
}
