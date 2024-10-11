import { Header } from './../../widgets/pageComponents/Header.tsx'
import cl from './BooksPage.module.css'
import { SearchPanel } from './components/SearchPanel.tsx'
import { FilterPanel } from './components/FilterPanel.tsx'

export function BooksPage(){
    return (
        <>
            <Header />
            <SearchPanel />
            <FilterPanel />
        </>
    )
}
