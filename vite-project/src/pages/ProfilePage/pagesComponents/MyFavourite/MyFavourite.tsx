import cl from '../MyFavourite/MyFavourite.module.css';
import { ListBooks } from "../../../../widgets/pageComponents/ListBooks";
import { Header } from "../../../../widgets/pageComponents/Header";
export default function MyFavourite() {
  return (
    <>
    <Header></Header>
      <div className={cl["container__FavouritesAndRead"]}>
				{(['Любимые']).map(title => <ListBooks coverPaths={[
                'https://cdn1.ozone.ru/s3/multimedia-c/6189288048.jpg',
                'https://i.pinimg.com/736x/af/44/80/af4480b2d95058a9c6790b45d0e05d13.jpg',
                
            ]} title={title}/>)}
      </div>
    </>
  );
}
