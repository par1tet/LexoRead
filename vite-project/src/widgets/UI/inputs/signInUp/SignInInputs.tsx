import styles from "../signInUp/SignInInputs.module.css";
export default function SignInInputs() {
  return (
    <>
      <div className={styles.container}>
        <input type="text" placeholder="Логин или email" className={styles.input} />
        <input type="password" placeholder="Пароль" className={styles.input} />
      </div>
    </>
  );
}
