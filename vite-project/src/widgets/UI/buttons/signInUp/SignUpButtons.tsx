import styles from "../../buttons/signInUp/SignUpButtons.module.css";
import { Link } from "react-router-dom";
interface signUpProps {
  children?: React.ReactNode;
}

export default function SignUpButton({ children }: signUpProps) {
  return (
    <div>
      <form>
        <Link to={"/reg"} className={styles.buttonstn}>{children}</Link>
      </form>
    </div>
  );
}
