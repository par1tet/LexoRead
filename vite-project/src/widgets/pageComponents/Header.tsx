import cl from "./Header.module.css";
import avatar from "./../../assets/img/avatar.png";
import { Link } from "react-router-dom";
import search from "./../../assets/img/search.svg";

export function Header() {
  return (
      <div className={cl["header"]}>
        <div className={cl["header__links"]}>
          <div className={cl["header__links-lexoRead"]}>
            <Link to="/">LexoRead</Link>
          </div>
          <div className={cl["header__links-main"]}>
            <Link to="/">Главное</Link>
          </div>
          <div className={cl["header__links-books"]}>
            <Link to="/">Книги</Link>
          </div>
          <div className={cl["header__links-aboutus"]}>
            <Link to="/aboutUs">О нас</Link>
          </div>
          <div className={cl["header__links-search"]}>
            <img src={search} />
          </div>
        </div>
        <div className={cl["header__logo"]}>
          <img src={avatar} />
        </div>
      </div>
  );
}
