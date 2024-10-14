import styles from "../signInUp/SignInButtons.module.css";
import SignUpButton from "./SignUpButtons";
interface signInProps {
  children?: React.ReactNode;
}
export default function SignInButtons({ children }: signInProps) {
  return (
    <>
      <div className={styles.container}>
        <form>
          <button className={styles.btn}>{children}</button>
          <SignUpButton>Зарегистрироваться</SignUpButton>
        </form>
      </div>
    </>
  );
}
