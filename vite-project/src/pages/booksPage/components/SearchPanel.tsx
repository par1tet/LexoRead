import cl from './../BooksPage.module.css'

export function SearchPanel(){
    return (
        <div className={cl['searchPanel']}>
            <input placeholder='Search books' />
        </div>
    )
}
