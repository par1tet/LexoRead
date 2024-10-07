import  styles from "../../inputs/signInUp/SignUpInputs.module.css";

export default function SignUpInputs() {
	return (
		<>
		<input type="text" placeholder="Ваше Имя" className={styles.input}/>
		<input type="email" placeholder="Введите почту" className={styles.input}/>
		<input type="password" placeholder="Пароль" className={styles.input}/>
		<input type="password" placeholder="Повторите пароль" className={styles.input}/>
		</>
	)
}
