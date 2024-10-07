import styles from '../signInUp/SignInButtons.module.css'
import SignUpButton from './SignUpButtons';
interface signInProps {
  children?: React.ReactNode;
  signUpClick?: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
	};
};
export default function SignInButtons({children}: signInProps) {
	return (
		<>
		<form>
		<button className={styles.btn} >{children}</button>
		<SignUpButton>Зарегистрироваться</SignUpButton>
		</form>
		</>
	)
}
