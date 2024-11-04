import { Header } from "../../widgets/pageComponents/Header";
import { ListBooks } from "../../widgets/pageComponents/ListBooks";
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
        <div className={cl["username"]}>smallgigach</div>
        <div className={cl["container__description"]}>
          <NavLink
            to={"/myfavourite"}
            className={cl["container__description-link"]}
          >
            Любимые <br />и прочитанные
          </NavLink>
          <textarea
            className={cl["description__input"]}
            onChange={(e) => console.log(e.target.value)}
          />
          <NavLink
            to={"/myfriends"}
            className={cl["container__description-link"]}
          >
            Друзья
          </NavLink>
        </div>
        <div className={cl["container__Books"]}>
          <span className={cl['container__Books-title-recommendations']}>Рекомендует книги</span>
          {[""].map((title) => (
            <ListBooks key={title}
              coverPaths={[
                "https://s3-alpha-sig.figma.com/img/5c93/9246/39a6427823c336a71051c8aa79c76531?Expires=1731283200&Key-Pair-Id=APKAQ4GOSFWCVNEHN3O4&Signature=JIumyb4McZ8Gc5g4NnD6ltdwKnMLvd5iy~TR6KKeDUwRuO31GWapHBOydJpMl4tgYr99NyrtbDjRhTUq~rGgX-Yo-ILeyXL5HqujGv1ya4dXReawTkb4P468OZSA0dreCGU7Aj5HxJmWnT77mzIJY20L7sPWmUS~IyDrd2u5qThXNNgEecd0w8kUK8J8HX5IQGV6Z~H7nxqXEVD8kzPgNu-AXBofnHHpHwkJtAB-FxdLlrGEsFl~JhAnsMTXwYXgfNbQBNlBXrb7ykUyAIP2nxxot7df1wJEEBQ5ClzEGxJ5N1TD9oSCi5lNcU9SjCit~Qu6XVSQzL2ZrycUCK~cRg__",
                "https://s3-alpha-sig.figma.com/img/5c93/9246/39a6427823c336a71051c8aa79c76531?Expires=1731283200&Key-Pair-Id=APKAQ4GOSFWCVNEHN3O4&Signature=JIumyb4McZ8Gc5g4NnD6ltdwKnMLvd5iy~TR6KKeDUwRuO31GWapHBOydJpMl4tgYr99NyrtbDjRhTUq~rGgX-Yo-ILeyXL5HqujGv1ya4dXReawTkb4P468OZSA0dreCGU7Aj5HxJmWnT77mzIJY20L7sPWmUS~IyDrd2u5qThXNNgEecd0w8kUK8J8HX5IQGV6Z~H7nxqXEVD8kzPgNu-AXBofnHHpHwkJtAB-FxdLlrGEsFl~JhAnsMTXwYXgfNbQBNlBXrb7ykUyAIP2nxxot7df1wJEEBQ5ClzEGxJ5N1TD9oSCi5lNcU9SjCit~Qu6XVSQzL2ZrycUCK~cRg__",
                "https://s3-alpha-sig.figma.com/img/5c93/9246/39a6427823c336a71051c8aa79c76531?Expires=1731283200&Key-Pair-Id=APKAQ4GOSFWCVNEHN3O4&Signature=JIumyb4McZ8Gc5g4NnD6ltdwKnMLvd5iy~TR6KKeDUwRuO31GWapHBOydJpMl4tgYr99NyrtbDjRhTUq~rGgX-Yo-ILeyXL5HqujGv1ya4dXReawTkb4P468OZSA0dreCGU7Aj5HxJmWnT77mzIJY20L7sPWmUS~IyDrd2u5qThXNNgEecd0w8kUK8J8HX5IQGV6Z~H7nxqXEVD8kzPgNu-AXBofnHHpHwkJtAB-FxdLlrGEsFl~JhAnsMTXwYXgfNbQBNlBXrb7ykUyAIP2nxxot7df1wJEEBQ5ClzEGxJ5N1TD9oSCi5lNcU9SjCit~Qu6XVSQzL2ZrycUCK~cRg__",
                "https://s3-alpha-sig.figma.com/img/5c93/9246/39a6427823c336a71051c8aa79c76531?Expires=1731283200&Key-Pair-Id=APKAQ4GOSFWCVNEHN3O4&Signature=JIumyb4McZ8Gc5g4NnD6ltdwKnMLvd5iy~TR6KKeDUwRuO31GWapHBOydJpMl4tgYr99NyrtbDjRhTUq~rGgX-Yo-ILeyXL5HqujGv1ya4dXReawTkb4P468OZSA0dreCGU7Aj5HxJmWnT77mzIJY20L7sPWmUS~IyDrd2u5qThXNNgEecd0w8kUK8J8HX5IQGV6Z~H7nxqXEVD8kzPgNu-AXBofnHHpHwkJtAB-FxdLlrGEsFl~JhAnsMTXwYXgfNbQBNlBXrb7ykUyAIP2nxxot7df1wJEEBQ5ClzEGxJ5N1TD9oSCi5lNcU9SjCit~Qu6XVSQzL2ZrycUCK~cRg__",
                "https://s3-alpha-sig.figma.com/img/5c93/9246/39a6427823c336a71051c8aa79c76531?Expires=1731283200&Key-Pair-Id=APKAQ4GOSFWCVNEHN3O4&Signature=JIumyb4McZ8Gc5g4NnD6ltdwKnMLvd5iy~TR6KKeDUwRuO31GWapHBOydJpMl4tgYr99NyrtbDjRhTUq~rGgX-Yo-ILeyXL5HqujGv1ya4dXReawTkb4P468OZSA0dreCGU7Aj5HxJmWnT77mzIJY20L7sPWmUS~IyDrd2u5qThXNNgEecd0w8kUK8J8HX5IQGV6Z~H7nxqXEVD8kzPgNu-AXBofnHHpHwkJtAB-FxdLlrGEsFl~JhAnsMTXwYXgfNbQBNlBXrb7ykUyAIP2nxxot7df1wJEEBQ5ClzEGxJ5N1TD9oSCi5lNcU9SjCit~Qu6XVSQzL2ZrycUCK~cRg__",
              ]}
              title={title}
            />
          ))}
        </div>
      </div>
    </>
  );
}
