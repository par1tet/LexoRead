import cl from './MainPage.module.css'
import { Header } from './../../widgets/pageComponents/Header.tsx'
import { ListBooks } from './../../widgets/pageComponents/ListBooks.tsx'
import { SearchPanel } from './components/SearchPanel.tsx'
import { FilterPanel } from './components/FilterPanel.tsx'

export function MainPage({}){
    return (
        <main className={cl['main']}>
            <Header />
            <FilterPanel />
            {(['Новое', 'Популярное', 'Программирование', 'Фантастика', 'Литература', 'Романтика', 'Бестселлеры', 'Биографии', 'Мистика', 'Ужасы']).map(title => <ListBooks coverPaths={[
                'https://i.pinimg.com/564x/f5/6a/23/f56a2337c49ee22c298b740c77cc8d17.jpg',
                'https://i.pinimg.com/736x/af/44/80/af4480b2d95058a9c6790b45d0e05d13.jpg',
                'https://i.pinimg.com/564x/7f/22/d5/7f22d598d7993db72141ce8bd3826b5b.jpg',
                'https://i.pinimg.com/564x/9f/8b/37/9f8b377f90c0919cffb31a83d0eb8f36.jpg',
                'https://i.pinimg.com/564x/0f/22/0f/0f220f78958c8f5e81270d18b692ba9c.jpg',
            ]} title={title}/>)}
        </main>
    )
}
