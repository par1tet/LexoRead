import cl from "./Header.module.css";
import avatar from "./../../assets/img/avatar.png";
import {NavLink} from "react-router-dom";
import backgroundLigth from "./../../assets/img/changeBackground/icons8-day-and-night-50.png"
import backgroundDark from "./../../assets/img/changeBackground/icons8-день-и-ночь-50.png"
import clsx from 'clsx'
import {useEffect, useState} from "react";
import ModalWindow from '../UI/ModalWindow/ModalWindow.tsx';

export function Header() {
    let [backgroundFon, setBackgroundFon] = useState(() => {
        const saved: boolean = localStorage.getItem("background");
        const initialValue = JSON.parse(saved);
        return initialValue || false;
    })
  const [isModalVisible, setModalVisible] = useState(false)
    const changeBackground = () => {
        setBackgroundFon(!backgroundFon);
    }
    useEffect(() => {
        if(backgroundFon) {
            localStorage.setItem("background", JSON.stringify(backgroundFon));
            document.body.style.backgroundColor= 'black';
            document.body.style.color = 'white';

        } else {
            localStorage.setItem("background", JSON.stringify(backgroundFon));
            document.body.style.backgroundColor = 'white';
            document.body.style.color = '#6F4C3E';
        }
    }, [backgroundFon]);

    return (
        <header className={cl["header"]}>
            <div className={cl["header__links"]}>
                <div className={clsx(cl["header__links-aboutus"], cl["header__links-link"])}>
                    <NavLink
                        to="/aboutUs"
                        className={({isActive}) => (isActive ? cl["about_us"] : undefined)}
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
                        className={({isActive}) => (isActive ? cl["about_us"] : undefined)}
                    >
                        Поиск
                    </NavLink>
                </div>
            </div>
            <div className={cl['background-img']} onClick={changeBackground}>
                {
                    !backgroundFon ?
                       <>
                        <img src={backgroundLigth}/>

                       </>
                        : (
                        <img src={backgroundDark}/>
                    )
                }
            </div>
            <div className={cl["header__logo"]}>
                <img src={avatar} onMouseOver={() => setModalVisible(true)}  />
              {isModalVisible && <ModalWindow></ModalWindow>}
            </div>
        </header>
    );
}
