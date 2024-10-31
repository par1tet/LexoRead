import styles from "../signInUp/SignInButtons.module.css";
interface signInProps {
  children?: React.ReactNode;
}
export default function SignInButtons({ children }: signInProps) {
  return (
    <>
      <div className={styles.container}>
        <form>
          <button className={styles.btn}>{children}</button>
          
        </form>
      </div>
    </>
  );
}
