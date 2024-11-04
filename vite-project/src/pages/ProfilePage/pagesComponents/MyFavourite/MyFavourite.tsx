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
            {(['Прочитанные']).map(title => <ListBooks coverPaths={[
                'https://content.img-gorod.ru/pim/products/images/0a/65/018f5e7b-de5a-78c0-8c88-fa0a304d0a65.jpg?width=304&height=438&fit=bounds',
                'https://content.img-gorod.ru/pim/products/images/0e/1a/01907874-04dd-7858-ab86-90b7bd840e1a.jpg?width=336&height=480&fit=bounds',
                'https://i.pinimg.com/564x/7f/22/d5/7f22d598d7993db72141ce8bd3826b5b.jpg',
                'https://i.pinimg.com/564x/9f/8b/37/9f8b377f90c0919cffb31a83d0eb8f36.jpg',
                'https://i.pinimg.com/564x/0f/22/0f/0f220f78958c8f5e81270d18b692ba9c.jpg',
            ]} title={title}/>)}
      </div>
    </>
  );
}
