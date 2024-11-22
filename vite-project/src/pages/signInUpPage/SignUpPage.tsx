import SignButton from "../../widgets/UI/buttons/signInUp/SignButton.tsx";
import SignInput from "../../widgets/UI/inputs/signInUp/SignInput.tsx";
import { Link } from "react-router-dom";
import cl from "./SignInUp.module.css";
export function SignUpPage() {
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
          <SignInput type="text" placeholder="Логин"></SignInput>
          <SignInput type="text" placeholder="Почта"></SignInput>
          <SignInput type="password" placeholder="Пароль"></SignInput>
          <SignInput type="password" placeholder="Повторный пароль"></SignInput>
        </div>
        <div className={cl['contanerSign__sign']}>
          <SignButton onClick={() => console.log('hi')}>Зарегистрироваться</SignButton>
          <Link to={"/auth"} className={cl['containerSign__link']}>
            Или войти в аккаунт
          </Link>
          </div>
        </div>
    </>
  );
}
