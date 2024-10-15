import styles from "../pagesComponents/MyFavourite.module.css";
import avatar from "../../../assets/img/avatar.png";
import { ListBooks } from "../../../widgets/pageComponents/ListBooks";
import { Header } from "../../../widgets/pageComponents/Header";
export default function MyFavourite() {
  return (
    <>
    <Header></Header>
      <div className={styles["container"]}>
        <div className={styles["header-logo"]}>
          <img src={avatar} alt="avatarUrl" />
        </div>
				<div className={styles["username"]}>Smallgigach</div>
				<span className={styles["sub-title"]}>Мои Книги</span>
				{(['Мои любимые']).map(title => <ListBooks coverPaths={[
                'https://i.pinimg.com/564x/f5/6a/23/f56a2337c49ee22c298b740c77cc8d17.jpg',
                'https://i.pinimg.com/736x/af/44/80/af4480b2d95058a9c6790b45d0e05d13.jpg',
                'https://i.pinimg.com/564x/7f/22/d5/7f22d598d7993db72141ce8bd3826b5b.jpg',
                'https://i.pinimg.com/564x/9f/8b/37/9f8b377f90c0919cffb31a83d0eb8f36.jpg',
                'https://i.pinimg.com/564x/0f/22/0f/0f220f78958c8f5e81270d18b692ba9c.jpg',
            ]} title={title}/>)}
      </div>
    </>
  );
}
