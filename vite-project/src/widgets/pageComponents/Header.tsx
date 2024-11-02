import cl from "./Header.module.css";
import avatar from "./../../assets/img/avatar.png";
import { NavLink } from "react-router-dom";
import search from "./../../assets/img/search.svg";
import clsx from 'clsx'

export function Header() {
  return (
    <header className={cl["header"]}>
      <div className={cl["header__links"]}>
        <div className={clsx(cl["header__links-aboutus"], cl["header__links-link"])}>
          <NavLink
            to="/aboutUs"
            className={({ isActive }) => (isActive ? cl["about_us"] : undefined)}
          >
            О нас
          </NavLink>
        </div> 
        <div className={clsx(cl["header__links-lexoRead"], cl["header__links-link"])}>
          <NavLink to="/">LexoRead</NavLink>
        </div>
        <div className={clsx(cl["header__links-aboutus"], cl["header__links-link"])}>
          <NavLink
            to="/aboutUs"
            className={({ isActive }) => (isActive ? cl["about_us"] : undefined)}
          >
            Поиск
          </NavLink>
        </div>
      </div>
      <div className={cl["header__logo"]}>
        <img src={avatar} />
      </div>
    </header>
  );
}
