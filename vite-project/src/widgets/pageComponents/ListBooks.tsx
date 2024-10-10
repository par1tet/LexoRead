import IBook from './../../entities/interfaces/IBook.ts'
import { createBookList } from './../../features/createBookList.ts'
import cl from './ListBooks.module.css'

type listBooksProps = {
    coverPaths: string[]
    title: string
}

export function ListBooks({coverPaths, title}: listBooksProps){
    return (
        <div className={cl['listBooks']}> 
            <div className={cl['listBooks__title']}>
            
                <span>{title}</span>
            </div>
            <div className={cl['listBooks__books']}>
                {(createBookList(coverPaths)).map((book: IBook, index: number) =>
                    <img src={book.coverUrl} key={index}/>
                )}
            </div>
        </div>
    )
}
