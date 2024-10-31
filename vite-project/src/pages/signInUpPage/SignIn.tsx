import OAuth from "../../widgets/pagesComponents/oAuth/oAuth";
import SignInButtons from "../../widgets/UI/buttons/signInUp/SignInButtons";
import SignInInputs from "../../widgets/UI/inputs/signInUp/SignInInputs";
import styles from "../signInUpPage/SignIn.module.css";
import SignUpButton from "../../widgets/UI/buttons/signInUp/SignUpButtons";
export function SignIn() {
  return (
    <>
      <div className={styles["background-image"]}>
        <div className={styles.containerSignIn}>
          <h1 className={styles.h1}>LexoRead</h1>
          <h2 className={styles.h2}>Войдите в LexoRead</h2>
          <SignInInputs></SignInInputs>
          <SignInButtons>Войти</SignInButtons>
          <div className={styles["or"]}>или</div>

          <SignUpButton>Зарегистрироваться</SignUpButton>

          <OAuth></OAuth>
        </div>
      </div>
    </>
  );
}
