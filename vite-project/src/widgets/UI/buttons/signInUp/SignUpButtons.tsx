import styles from "../../buttons/signInUp/SignUpButtons.module.css";
import { Link } from "react-router-dom";
import { Button } from "semantic-ui-react";
interface signUpProps {
  children?: React.ReactNode;
}

export default function SignUpButton({ children }: signUpProps) {
  return (
    <div>
      <form>
        <Link to={"/reg"}><Button className={styles.btn}>{children}</Button></Link>
      </form>
    </div>
  );
}
