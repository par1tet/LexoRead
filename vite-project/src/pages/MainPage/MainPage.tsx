import cl from './MainPage.module.css'
import { Header } from './../../widgets/pageComponents/Header.tsx'
import { ListBooks } from './../../widgets/pageComponents/ListBooks.tsx'

export function MainPage({}){
    return (
        <main className={cl['main']}>
            <Header />
            <div className={cl['main__lexoRead']}>
                <span>Lexoread</span>
            </div>
            <ListBooks coverPaths={[
                'https://i.pinimg.com/564x/f5/6a/23/f56a2337c49ee22c298b740c77cc8d17.jpg',
                'https://i.pinimg.com/736x/af/44/80/af4480b2d95058a9c6790b45d0e05d13.jpg'
            ]} title='Популярное'/>
        </main>
    )
}
