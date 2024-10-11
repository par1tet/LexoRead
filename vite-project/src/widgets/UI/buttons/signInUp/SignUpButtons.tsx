import styles from "../../buttons/signInUp/SignUpButtons.module.css";
import { Link } from "react-router-dom";
interface signUpProps {
  children?: React.ReactNode;
}

export default function SignUpButton({ children }: signUpProps) {
  return (
    <div className={styles.container}>
      <form>
        <Link to={"/reg"} className={styles.buttonToReg}>{children}</Link>
      </form>
    </div>
  );
}
