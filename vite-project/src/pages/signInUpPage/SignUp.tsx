import SignUpButton from "../../widgets/UI/buttons/sign-up/SignUp.tsx";
import SignUpInputs from "../../widgets/UI/inputs/sign-up/SignUp.tsx";
import styles from  "../signInUpPage/SignUp.module.css";
export default function SignUp() {
  return (
    <>
      <h1 className={styles.h1}>LexoRead</h1>
      <h2 className={styles.h2}>Зарегистрируйтесь в <br /> LexoRead</h2>
      <SignUpInputs></SignUpInputs>
      <SignUpButton>Зарегистрироваться</SignUpButton>
      
    </>
  );
}
