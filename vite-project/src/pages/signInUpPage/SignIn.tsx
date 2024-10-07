import OAuth from '../../widgets/pagesComponents/oAuth/oAuth'
import SignInButtons from '../../widgets/UI/buttons/signInUp/SignInButtons'
import SignInInputs from '../../widgets/UI/inputs/signInUp/SignInInputs'
import styles from '../signInUpPage/SignIn.module.css'

export function SignIn() {
	
	return (
		<>
		<h1 className={styles.h1}>LexoRead</h1>
		<h2 className={styles.h2}>Войдите в <br />LexoRead</h2>
		<SignInInputs></SignInInputs>
		<SignInButtons>Войти</SignInButtons>
		<OAuth></OAuth>
		</>
	)
}
