import SignButton from "../../widgets/UI/buttons/signInUp/SignButton.tsx";
import SignInput from "../../widgets/UI/inputs/signInUp/SignInput.tsx";
import {Link, useSearchParams} from "react-router-dom";
import cl from "./SignInUp.module.css";
import clsx from 'clsx'
import googleIcon from './../../assets/img/signInWith/google.svg'
import githubIcon from './../../assets/img/signInWith/github.svg'
import vkIcon from './../../assets/img/signInWith/vk.svg'
export function SignInPage() {

  return (
    <>
        <div className={cl['background']}></div>
        <div className={cl['containerSign']}>
          <h1 className={cl['containerSign__title']}>
            LexoRead
          </h1>
          <h2 className={cl['containerSign__subTitle']}>
            Зарегистрируйтесь в LexoRead
          </h2>
        <div className={cl['contanerSign__inputs']}>
          <SignInput type="text" placeholder="Логин или email"></SignInput>
          <SignInput type="text" placeholder="Пароль" ></SignInput>
        </div>
        <div className={clsx(cl['contanerSign__sign'], cl['contanerSign__sign-signin'])}>
          <SignButton onClick={() => console.log('hi')}>Войти</SignButton>
          <Link to={"/reg"} className={cl['containerSign__link']}>
            Или зарегистрироваться
          </Link>
          <div className={cl['signinWith']}>
            <img src={googleIcon} />
            <img src={vkIcon} />
            <img src={githubIcon} />
          </div>
          </div>
        </div>
    </>
  );
}
