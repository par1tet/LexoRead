import SignUpButton from "../../widgets/UI/buttons/signInUp/SignUpButtons.tsx";
import SignUpInputs from "../../widgets/UI/inputs/signInUp/SignUpInputs.tsx";
import styles from  "../signInUpPage/SignUp.module.css";
import OAuth from "../../widgets/pagesComponents/oAuth/oAuth.tsx";
export default function SignUp() {

  return (
    <>
      <h1 className={styles.h1}>LexoRead</h1>
      <h2 className={styles.h2}>Зарегистрируйтесь в <br /> LexoRead</h2>
      <SignUpInputs></SignUpInputs>
      <SignUpButton>Зарегистрироваться</SignUpButton>
      <OAuth></OAuth>
    </>
  );
}
