import cl from "./SignButton.module.css";
import { Link } from "react-router-dom";

type signUpProps = {
    onClick: any,
  children?: React.ReactNode;
}

export default function SignButton({children, onClick}: signUpProps) {

  return (
    <button onClick={onClick} className={cl['button']}>
        <span>{children}</span>
    </button>
  );
}
