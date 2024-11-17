import AvatarAndNick from "../../widgets/pageComponents/AvatarAndNicknameForProfilePage/AvatarAndNick";
import {Header} from "../../widgets/pageComponents/Header";
import {ListBooks} from "../../widgets/pageComponents/ListBooks";
import cl from "../ProfilePage/Profile.module.css";
import {NavLink} from "react-router-dom";
import {useEffect, useState} from "react";

export default function Profile() {
    const [description, setDescription] = useState(() => {
        const saved: string | null = localStorage.getItem("description");
        const initialValue = JSON.parse(saved);
        return initialValue || "";
    });
    useEffect(() => {
        localStorage.setItem('description', JSON.stringify(description));
    }, [description]);

    return (
        <>
            <Header></Header>
            <div className={cl["container"]}>
                <AvatarAndNick nickname={"smallgigach"} friends=""></AvatarAndNick>
                <div className={cl["container__description"]}>
                    <NavLink
                        to={"/myfavourite"}
                        className={cl["container__description-link"]}
                    >
                        Любимые <br/>и прочитанные
                    </NavLink>
                    <textarea
                        className={cl["description__input"]}
                        onChange={(e) => setDescription(e.target.value)}
                        value={description}
                    />
                    <NavLink
                        to={"/myfriends"}
                        className={cl["container__description-link"]}
                    >
                        Друзья
                    </NavLink>
                </div>
                <div className={cl["container__Books"]}>
          <span className={cl["container__Books-title-recommendations"]}>
            Рекомендует книги
          </span>
                    {[""].map((title) => (
                        <ListBooks
                            key={title}
                            coverPaths={[
                                "https://cdn.litres.ru/pub/c/cover_415/69040270.webp",
                                "https://cdn.litres.ru/pub/c/cover_415/69040270.webp",
                                "https://cdn.litres.ru/pub/c/cover_415/69040270.webp",
                                "https://cdn.litres.ru/pub/c/cover_415/69040270.webp",
                                'https://cdn.litres.ru/pub/c/cover_415/69040270.webp'
                            ]}
                            title={title}
                        />
                    ))}
                </div>

            </div>
        </>
    );
}
