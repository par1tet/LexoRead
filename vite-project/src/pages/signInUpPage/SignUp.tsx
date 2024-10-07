import SignUpButton from "../../widgets/UI/buttons/signInUp/SignUpButtons.tsx";
import SignUpInputs from "../../widgets/UI/inputs/signInUp/SignUpInputs.tsx";
import styles from  "./SignUp.module.css";
import OAuth from "../../widgets/pagesComponents/oAuth/oAuth.tsx";
import { Link } from "react-router-dom";

export function SignUp() {
  return (
    <>
      <h1 className={styles.h1}>LexoRead</h1>
      <h2 className={styles.h2}>Зарегистрируйтесь в <br /> LexoRead</h2>
      <SignUpInputs></SignUpInputs>
      <SignUpButton>Зарегистрироваться</SignUpButton>
      <Link to={"/auth"} className={styles.backToAuth}>Войти в аккаунт</Link>
      <OAuth></OAuth>
    </>
  );
}
