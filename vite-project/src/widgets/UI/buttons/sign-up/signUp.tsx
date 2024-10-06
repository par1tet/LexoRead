import  styles from "../../buttons/sign-up/SignUpButtons.module.css";
interface signUpProps {
  children?: React.ReactNode;
  onClick?: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
		
	};
};
export default function SignUpButton({ children, onClick }: signUpProps) {
  return (
    <div>
      <button className={styles.btn} onClick={onClick}>{children}</button>
    </div>
  );
}
