import styles from "../signInUp/SignInInputs.module.css"
export default function SignInInputs() {
	return (
		<>
		<input type="text" placeholder="Почту" className={styles.input}/>
		<input type="password" placeholder="Пароль" className={styles.input}/>
		</>
	)
}
